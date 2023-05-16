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
)

type equipSetChargingTimerCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

//func (e *equipSetChargingTimerRequest) SetError(err *callback.CallbackError) {
//	code := err.Code()
//	msg := err.Error()
//	e.Data.Code = &code
//	e.Data.Msg = &msg
//}

func (*equipSetChargingTimerCallbackRequest) GetName() string {
	return services.SetChargingTimer.String()
}

func NewEquipSetChargingTimerCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipSetChargingTimerCallbackRequest {
	req := &equipSetChargingTimerCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.SetChargingTimer.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipSetChargingTimerCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *callback.CallbackError) *equipSetChargingTimerCallbackRequest {
	req := &equipSetChargingTimerCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.SetChargingTimer.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
	}
	return req
}

var _ services.Response = &equipSetChargingTimerCallbackResponse{}

type equipSetChargingTimerCallbackResponse struct {
	api.Response
}

func (resp *equipSetChargingTimerCallbackResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipSetChargingTimerCallbackResponse) GetMsg() string {
	return resp.Msg
}

func SetChargingTimerCallbackRequest(ctx context.Context, req services.CallbackRequest) error {
	headerValue := make([]string, 0)
	headerValue = append(headerValue, api.Services)
	headerValue = append(headerValue, services.SetChargingTimer.Split()...)
	headerValue = append(headerValue, services.Callback)

	header := map[string]string{api.Perms: strings.Join(headerValue, ":")}

	url := config.App.ServicesUrl + services.Equip + "/" + services.Callback + "/" + req.GetName() + "Callback"

	message, err := api.SendRequest(ctx, url, req, header)
	if err != nil {
		return err
	}

	resp := &equipSetChargingTimerCallbackResponse{}
	err = json.Unmarshal(message, resp)
	if err != nil {
		return err
	}

	if resp.Status == 1 {
		return errors.New(resp.Msg)
	}

	return nil
}

func SetChargingTimerCallbackRequestWithGeneric(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.SetChargingTimer)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipSetChargingTimerCallbackResponse{})
}
