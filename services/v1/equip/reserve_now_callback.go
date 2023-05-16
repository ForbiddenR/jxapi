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

type equipReserveNowCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (r *equipReserveNowCallbackRequest) GetName() string {
	return services.ReserveNow.String()
}

func NewEquipReserveNowCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipReserveNowCallbackRequest {
	req := &equipReserveNowCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.ReserveNow.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipReserveNowCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *callback.CallbackError) *equipReserveNowCallbackRequest {
	req := &equipReserveNowCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.ReserveNow.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
	}
	return req
}

type equipReserveNowCallbackResponse struct {
	api.Response
}

func (resp *equipReserveNowCallbackResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipReserveNowCallbackResponse) GetMsg() string {
	return resp.Msg
}

func ReserveNowCallbackRequest(ctx context.Context, req services.CallbackRequest) error {
	headerValue := make([]string, 0)
	//headerValue = append(headerValue, api.Services, "remote", "start", services.Callback)
	headerValue = append(headerValue, api.Services)
	headerValue = append(headerValue, services.ReserveNow.Split()...)
	headerValue = append(headerValue, services.Callback)

	header := map[string]string{api.Perms: strings.Join(headerValue, ":")}
	url := config.App.ServicesUrl + services.Equip + "/" + services.Callback + "/" + req.GetName() + "Callback"

	message, err := api.SendRequest(ctx, url, req, header)
	if err != nil {
		return err
	}

	resp := &equipReserveNowCallbackResponse{}
	err = json.Unmarshal(message, resp)
	if err != nil {
		return err
	}

	if resp.Status == 1 {
		return errors.New(resp.Msg)
	}

	return nil
}

func ReserveNowCallbackRequestWithGeneric(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.ReserveNow)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipReserveNowCallbackResponse{})
}
