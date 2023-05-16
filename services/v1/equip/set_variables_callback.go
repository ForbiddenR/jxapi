package equip

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"gitee.com/csms/jxeu-ocpp/internal/config"
	internalError "gitee.com/csms/jxeu-ocpp/internal/errors"
	"gitee.com/csms/jxeu-ocpp/pkg/api"
	"gitee.com/csms/jxeu-ocpp/pkg/api/services"
	"gitee.com/csms/jxeu-ocpp/pkg/ocpp1.6/protocol"
)

const (
	SetVariablesAccept         = 0
	SetVariablesRejected       = 1
	SetVariablesRebootRequired = 2
	SetVariablesNotSupported   = 3
)

func OCPP16SetVariablesStatus(status protocol.ChangeConfigurationResponseJsonStatus) int {
	switch status {
	case protocol.ChangeConfigurationResponseJsonStatusAccepted:
		return SetVariablesAccept
	case protocol.ChangeConfigurationResponseJsonStatusRejected:
		return SetVariablesRejected
	case protocol.ChangeConfigurationResponseJsonStatusRebootRequired:
		return SetVariablesRebootRequired
	default:
		return SetVariablesNotSupported
	}
}

type equipSetVariablesCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (s *equipSetVariablesCallbackRequest) GetName() string {
	return services.ChangeConfiguration.String()
}

func NewEquipSetVariablesCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipSetVariablesCallbackRequest {
	req := &equipSetVariablesCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.ChangeConfiguration.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipSetVariablesRequestError(sn, pod, msgID string, p *services.Protocol, err *internalError.CallbackError) *equipSetVariablesCallbackRequest {
	req := &equipSetVariablesCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.ChangeConfiguration.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
	}
	return req
}

var _ services.Response = &equipSetVariablesCallbackResponse{}

type equipSetVariablesCallbackResponse struct {
	api.Response
}

func (resp *equipSetVariablesCallbackResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipSetVariablesCallbackResponse) GetMsg() string {
	return resp.Msg
}

func SetVariablesCallbackRequest(ctx context.Context, req services.CallbackRequest) error {
	headerValue := make([]string, 0)
	//headerValue = append(headerValue, api.Services, "set", "variables", services.Callback)
	headerValue = append(headerValue, api.Services)
	headerValue = append(headerValue, services.ChangeConfiguration.Split()...)
	headerValue = append(headerValue, services.Callback)

	header := map[string]string{api.Perms: strings.Join(headerValue, ":")}

	url := config.App.ServicesUrl + services.Equip + "/" + services.Callback + "/" + req.GetName() + "Callback"

	message, err := api.SendRequest(ctx, url, req, header)
	if err != nil {
		return err
	}

	resp := &equipSetVariablesCallbackResponse{}

	err = json.Unmarshal(message, resp)

	if err != nil {
		return err
	}

	if resp.Status == 1 {
		return errors.New(resp.Msg)
	}

	return nil
}

func SetVariablesRequestWithGeneric(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.ChangeConfiguration)

	url := services.GetCallbackURL(req)
	return services.RequestWithoutResponse(ctx, req, url, header, &equipSetVariablesCallbackResponse{})
}
