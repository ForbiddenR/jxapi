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

type equipCancelReservationCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (r *equipCancelReservationCallbackRequest) GetName() string {
	return services.CancelReservation.String()
}

func NewEquipCancelReseravtionCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipCancelReservationCallbackRequest {
	req := &equipCancelReservationCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.CancelReservation.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipCancelReservationCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *callback.CallbackError) *equipCancelReservationCallbackRequest {
	req := &equipCancelReservationCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.CancelReservation.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
	}
	return req
}

var _ services.Response = &equipCancelReservationCallbackResponse{}

type equipCancelReservationCallbackResponse struct {
	api.Response
}

func (resp *equipCancelReservationCallbackResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipCancelReservationCallbackResponse) GetMsg() string {
	return resp.Msg
}

func CancelReservationCallbackRequest(ctx context.Context, req services.CallbackRequest) error {
	headerValue := make([]string, 0)
	//headerValue = append(headerValue, api.Services, "remote", "start", services.Callback)
	headerValue = append(headerValue, api.Services)
	headerValue = append(headerValue, services.CancelReservation.Split()...)
	headerValue = append(headerValue, services.Callback)

	header := map[string]string{api.Perms: strings.Join(headerValue, ":")}
	url := config.App.ServicesUrl + services.Equip + "/" + services.Callback + "/" + req.GetName() + "Callback"

	message, err := api.SendRequest(ctx, url, req, header)
	if err != nil {
		return err
	}

	resp := &equipCancelReservationCallbackResponse{}
	err = json.Unmarshal(message, resp)
	if err != nil {
		return err
	}

	if resp.Status == 1 {
		return errors.New(resp.Msg)
	}

	return nil
}

func CancelReservationCallbackRequestWithGeneric(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.CancelReservation)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipCancelReservationCallbackResponse{})
}
