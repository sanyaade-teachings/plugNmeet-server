package bbbmodel

import (
	"fmt"
	"github.com/mynaparrot/plugnmeet-protocol/bbbapiwrapper"
	"github.com/mynaparrot/plugnmeet-server/pkg/config"
	"github.com/mynaparrot/plugnmeet-server/pkg/models/recordingmodel"
	"github.com/mynaparrot/plugnmeet-server/pkg/services/dbservice"
	log "github.com/sirupsen/logrus"
	"strings"
)

type BBBApiWrapperModel struct {
	ds *dbservice.DatabaseService
}

func NewBBBApiWrapperModel() *BBBApiWrapperModel {
	ds := dbservice.NewDBService(config.AppCnf.ORM)
	return &BBBApiWrapperModel{
		ds: ds,
	}
}

func (m *BBBApiWrapperModel) GetRecordings(host string, r *bbbapiwrapper.GetRecordingsReq) ([]*bbbapiwrapper.RecordingInfo, *bbbapiwrapper.Pagination, error) {
	oriIds := make(map[string]string)
	if r.Limit == 0 {
		// let's make it 50 for BBB as not all plugin still support pagination
		r.Limit = 50
	}
	var rIds []string

	if r.RecordID != "" {
		rIds = strings.Split(r.RecordID, ",")
	} else if r.MeetingID != "" {
		rIds = strings.Split(r.MeetingID, ",")
		for _, rd := range rIds {
			fId := bbbapiwrapper.CheckMeetingIdToMatchFormat(rd)
			oriIds[fId] = rd
		}
	}

	data, total, err := m.ds.GetRecordingsForBBB(rIds, rIds, r.Offset, r.Limit)
	if err != nil {
		return nil, nil, err
	}

	var recordings []*bbbapiwrapper.RecordingInfo
	for _, v := range data {
		//	err = rows.Scan(&recording.RecordID, &meetingId, &rSid, &path, &size, &recording.Published, &recording.Name, &participants, &created, &ended)
		recording := bbbapiwrapper.RecordingInfo{
			RecordID:          v.RecorderID,
			InternalMeetingID: v.RoomSid.String,
		}

		if oriIds[v.RoomID] != "" {
			recording.MeetingID = oriIds[v.RoomID]
		} else {
			recording.MeetingID = v.RoomID
		}

		// for path, let's create a download link directly
		url, err := m.createPlayBackURL(host, v.FilePath)
		if err != nil {
			log.Errorln(err)
			continue
		}
		recording.Playback.PlayBackFormat = []bbbapiwrapper.PlayBackFormat{
			{
				Type: "presentation",
				URL:  url,
			},
		}

		if mInfo, err := m.ds.GetRoomInfoBySid(v.RoomSid.String, nil); err != nil {
			recording.StartTime = mInfo.Created.UnixMilli()
			recording.EndTime = mInfo.Ended.UnixMilli()
			recording.Participants = uint64(mInfo.JoinedParticipants)
		}

		if recording.Published {
			recording.State = "published"
		}
		if v.Size > 0 {
			recording.RawSize = int64(v.Size * 1000000)
			recording.Size = recording.RawSize
		}
		recordings = append(recordings, &recording)
	}

	pagination := &bbbapiwrapper.Pagination{
		Pageable: &bbbapiwrapper.PaginationPageable{
			Offset: r.Offset,
			Limit:  r.Limit,
		},
		TotalElements: uint64(total),
	}
	if total == 0 {
		pagination.Empty = true
	}

	return recordings, pagination, nil
}

func (m *BBBApiWrapperModel) createPlayBackURL(host, path string) (string, error) {
	auth := recordingmodel.NewRecordingAuth()
	token, err := auth.CreateTokenForDownload(path)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/download/recording/%s", host, token)
	return url, nil
}
