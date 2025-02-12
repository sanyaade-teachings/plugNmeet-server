package controllers

import (
	"github.com/bufbuild/protovalidate-go"
	"github.com/gofiber/fiber/v2"
	"github.com/mynaparrot/plugnmeet-protocol/plugnmeet"
	"github.com/mynaparrot/plugnmeet-protocol/utils"
	"github.com/mynaparrot/plugnmeet-server/pkg/models"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var op = protojson.UnmarshalOptions{
	DiscardUnknown: true,
}

func validateProtoRequest(msg proto.Message) error {
	v, err := protovalidate.New()
	if err != nil {
		return err
	}

	if err = v.Validate(msg); err != nil {
		return err
	}
	return nil
}

func HandleFetchAnalytics(c *fiber.Ctx) error {
	req := new(plugnmeet.FetchAnalyticsReq)
	err := op.Unmarshal(c.Body(), req)
	if err != nil {
		return utils.SendCommonProtoJsonResponse(c, false, err.Error())
	}

	if err = validateProtoRequest(req); err != nil {
		return utils.SendCommonProtoJsonResponse(c, false, err.Error())
	}

	m := models.NewAnalyticsModel(nil, nil, nil)
	result, err := m.FetchAnalytics(req)

	if err != nil {
		return utils.SendCommonProtoJsonResponse(c, false, err.Error())
	}
	if result.GetTotalAnalytics() == 0 {
		return utils.SendCommonProtoJsonResponse(c, false, "no analytics found")
	}

	r := &plugnmeet.FetchAnalyticsRes{
		Status: true,
		Msg:    "success",
		Result: result,
	}
	return utils.SendProtoJsonResponse(c, r)
}

func HandleDeleteAnalytics(c *fiber.Ctx) error {
	req := new(plugnmeet.DeleteAnalyticsReq)
	err := op.Unmarshal(c.Body(), req)
	if err != nil {
		return utils.SendCommonProtoJsonResponse(c, false, err.Error())
	}

	if err = validateProtoRequest(req); err != nil {
		return utils.SendCommonProtoJsonResponse(c, false, err.Error())
	}

	m := models.NewAnalyticsModel(nil, nil, nil)
	err = m.DeleteAnalytics(req)
	if err != nil {
		return utils.SendCommonProtoJsonResponse(c, false, err.Error())
	}

	return utils.SendCommonProtoJsonResponse(c, true, "success")
}

func HandleGetAnalyticsDownloadToken(c *fiber.Ctx) error {
	req := new(plugnmeet.GetAnalyticsDownloadTokenReq)
	err := op.Unmarshal(c.Body(), req)
	if err != nil {
		return utils.SendCommonProtoJsonResponse(c, false, err.Error())
	}

	if err = validateProtoRequest(req); err != nil {
		return utils.SendCommonProtoJsonResponse(c, false, err.Error())
	}

	m := models.NewAnalyticsModel(nil, nil, nil)
	token, err := m.GetAnalyticsDownloadToken(req)

	if err != nil {
		return utils.SendCommonProtoJsonResponse(c, false, err.Error())
	}

	r := &plugnmeet.GetAnalyticsDownloadTokenRes{
		Status: true,
		Msg:    "success",
		Token:  &token,
	}
	return utils.SendProtoJsonResponse(c, r)
}

func HandleDownloadAnalytics(c *fiber.Ctx) error {
	token := c.Params("token")

	if len(token) == 0 {
		return c.Status(fiber.StatusUnauthorized).SendString("token require or invalid url")
	}

	m := models.NewAnalyticsModel(nil, nil, nil)
	file, status, err := m.VerifyAnalyticsToken(token)

	if err != nil {
		return c.Status(status).SendString(err.Error())
	}

	c.Attachment(file)
	return c.SendFile(file, false)
}
