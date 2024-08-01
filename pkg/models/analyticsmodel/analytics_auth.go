package analyticsmodel

import (
	"errors"
	"fmt"
	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/mynaparrot/plugnmeet-protocol/plugnmeet"
	"github.com/mynaparrot/plugnmeet-server/pkg/config"
	"github.com/mynaparrot/plugnmeet-server/pkg/dbmodels"
	"github.com/mynaparrot/plugnmeet-server/pkg/services/dbservice"
	"os"
	"strings"
	"time"
)

type AnalyticsAuthModel struct {
	app *config.AppConfig
	ds  *dbservice.DatabaseService
}

func NewAnalyticsAuthModel(app *config.AppConfig, ds *dbservice.DatabaseService) *AnalyticsAuthModel {
	if app == nil {
		app = config.GetConfig()
	}
	if ds == nil {
		ds = dbservice.NewDBService(app.ORM)
	}

	return &AnalyticsAuthModel{
		app: config.GetConfig(),
		ds:  ds,
	}
}

func (m *AnalyticsAuthModel) AddAnalyticsFileToDB(roomTableId uint64, roomCreationTime int64, roomId, fileId string, stat os.FileInfo) (int64, error) {
	fSize := float64(stat.Size())
	// we'll convert bytes to KB
	if fSize > 1000 {
		fSize = fSize / 1000.0
	} else {
		fSize = 1
	}

	info := &dbmodels.Analytics{
		RoomTableID:      roomTableId,
		RoomID:           roomId,
		FileID:           fileId,
		FileName:         fileId + ".json",
		FileSize:         fSize,
		RoomCreationTime: roomCreationTime,
	}

	return m.ds.InsertAnalyticsData(info)
}

func (m *AnalyticsAuthModel) FetchAnalytics(r *plugnmeet.FetchAnalyticsReq) (*plugnmeet.FetchAnalyticsResult, error) {
	data, total, err := m.ds.GetAnalytics(r.RoomIds, uint64(r.From), uint64(r.Limit), &r.OrderBy)
	if err != nil {
		return nil, err
	}

	var analytics []*plugnmeet.AnalyticsInfo
	for _, v := range data {
		analytic := &plugnmeet.AnalyticsInfo{
			RoomId:           v.RoomID,
			FileId:           v.FileID,
			FileSize:         v.FileSize,
			CreationTime:     v.CreationTime,
			RoomCreationTime: v.RoomCreationTime,
		}
		analytics = append(analytics, analytic)
	}

	result := &plugnmeet.FetchAnalyticsResult{
		TotalAnalytics: total,
		From:           r.From,
		Limit:          r.Limit,
		OrderBy:        r.OrderBy,
		AnalyticsList:  analytics,
	}

	return result, nil
}

func (m *AnalyticsAuthModel) fetchAnalytic(fileId string) (*plugnmeet.AnalyticsInfo, error) {
	v, err := m.ds.GetAnalyticByFileId(fileId)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, errors.New("no info found")
	}
	analytic := &plugnmeet.AnalyticsInfo{
		RoomId:           v.RoomID,
		FileId:           v.FileID,
		FileSize:         v.FileSize,
		CreationTime:     v.CreationTime,
		RoomCreationTime: v.RoomCreationTime,
	}

	return analytic, nil
}

func (m *AnalyticsAuthModel) getAnalyticByRoomTableId(roomTableId uint64) (*plugnmeet.AnalyticsInfo, error) {
	v, err := m.ds.GetAnalyticByRoomTableId(roomTableId)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, errors.New("no info found")
	}
	analytic := &plugnmeet.AnalyticsInfo{
		RoomId:           v.RoomID,
		FileId:           v.FileID,
		FileSize:         v.FileSize,
		CreationTime:     v.CreationTime,
		RoomCreationTime: v.RoomCreationTime,
	}

	return analytic, nil
}

func (m *AnalyticsAuthModel) DeleteAnalytics(r *plugnmeet.DeleteAnalyticsReq) error {
	analytic, err := m.fetchAnalytic(r.FileId)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/%s", *config.GetConfig().AnalyticsSettings.FilesStorePath, analytic.FileName)

	// delete the main file
	err = os.Remove(path)
	if err != nil {
		// if file not exists then we can delete it from record without showing any error
		if !os.IsNotExist(err) {
			ms := strings.SplitN(err.Error(), "/", -1)
			return errors.New(ms[len(ms)-1])
		}
	}

	// delete compressed, if any
	_ = os.Remove(path + ".fiber.gz")

	// no error, so we'll delete record from DB
	_, err = m.ds.DeleteAnalyticByFileId(analytic.FileId)
	if err != nil {
		return err
	}
	return nil
}

// GetAnalyticsDownloadToken will use the same JWT token generator as plugNmeet is using
func (m *AnalyticsAuthModel) GetAnalyticsDownloadToken(r *plugnmeet.GetAnalyticsDownloadTokenReq) (string, error) {
	analytic, err := m.fetchAnalytic(r.FileId)
	if err != nil {
		return "", err
	}

	return m.generateToken(analytic.FileName)
}

func (m *AnalyticsAuthModel) generateToken(fileName string) (string, error) {
	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: []byte(m.app.Client.Secret)}, (&jose.SignerOptions{}).WithType("JWT"))

	if err != nil {
		return "", err
	}

	cl := jwt.Claims{
		Issuer:    m.app.Client.ApiKey,
		NotBefore: jwt.NewNumericDate(time.Now().UTC()),
		Expiry:    jwt.NewNumericDate(time.Now().UTC().Add(*m.app.AnalyticsSettings.TokenValidity)),
		Subject:   fileName,
	}

	return jwt.Signed(sig).Claims(cl).Serialize()
}

// VerifyAnalyticsToken verify token & provide file path
func (m *AnalyticsAuthModel) VerifyAnalyticsToken(token string) (string, int, error) {
	tok, err := jwt.ParseSigned(token, []jose.SignatureAlgorithm{jose.HS256})
	if err != nil {
		return "", fiber.StatusUnauthorized, err
	}

	out := jwt.Claims{}
	if err = tok.Claims([]byte(m.app.Client.Secret), &out); err != nil {
		return "", fiber.StatusUnauthorized, err
	}

	if err = out.Validate(jwt.Expected{Issuer: config.GetConfig().Client.ApiKey, Time: time.Now().UTC()}); err != nil {
		return "", fiber.StatusUnauthorized, err
	}

	file := fmt.Sprintf("%s/%s", *m.app.AnalyticsSettings.FilesStorePath, out.Subject)
	_, err = os.Lstat(file)

	if err != nil {
		ms := strings.SplitN(err.Error(), "/", -1)
		return "", fiber.StatusNotFound, errors.New(ms[len(ms)-1])
	}

	return file, fiber.StatusOK, nil
}
