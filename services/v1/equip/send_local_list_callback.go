package equip

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"gitee.com/csms/jxeu-ocpp/internal/config"
	callback "gitee.com/csms/jxeu-ocpp/internal/errors"
	"gitee.com/csms/jxeu-ocpp/pkg/api"
	"gitee.com/csms/jxeu-ocpp/pkg/api/services"
	"gitee.com/csms/jxeu-ocpp/pkg/ocpp1.6/protocol"
)

const (
	SendLocalListAccept          = 0
	SendLocalListFailed          = 1
	SendLocalListVersionMismatch = 2
	SendLocalListNotSupported    = 3
)

func OCPP16SendLocalListStatus(status protocol.SendLocalListResponseJsonStatus) int {
	switch status {
	case protocol.SendLocalListResponseJsonStatusAccepted:
		return SendLocalListAccept
	case protocol.SendLocalListResponseJsonStatusFailed:
		return SendLocalListFailed
	case protocol.SendLocalListResponseJsonStatusVersionMismatch:
		return SendLocalListVersionMismatch
	default:
		return SendLocalListNotSupported
	}
}

type equipSendLocalListCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (s *equipSendLocalListCallbackRequest) GetName() string {
	return services.SendLocalList.String()
}

func NewEquipSendLocalListCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipSendLocalListCallbackRequest {
	req := &equipSendLocalListCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.SendLocalList.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipSendLocalListCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *callback.CallbackError) *equipSendLocalListCallbackRequest {
	req := &equipSendLocalListCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.SendLocalList.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
	}
	return req
}

var _ services.Response = &equipSendLocalListResponse{}

type equipSendLocalListResponse struct {
	api.Response
}

func (resp *equipSendLocalListResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipSendLocalListResponse) GetMsg() string {
	return resp.Msg
}

func SendLocalListCallbackRequest(ctx context.Context, req services.CallbackRequest) error {
	headerValue := make([]string, 0)
	headerValue = append(headerValue, api.Services)
	headerValue = append(headerValue, "set", "local", "authorize", services.Callback)

	header := map[string]string{api.Perms: strings.Join(headerValue, ":")}

	url := config.App.ServicesUrl + services.Equip + "/" + services.Callback + "/" + req.GetName() + "Callback"

	message, err := api.SendRequest(ctx, url, req, header)

	if err != nil {
		return err
	}

	resp := &equipSendLocalListResponse{}
	err = json.Unmarshal(message, resp)

	if err != nil {
		return err
	}

	if resp.Status == 1 {
		return errors.New(resp.Msg)
	}

	return nil
}

func SendLocalListCallbackRequestWithGeneric(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.SendLocalList)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipSendLocalListResponse{})
}
