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

type equipRemoteStartTransactionCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (r *equipRemoteStartTransactionCallbackRequest) GetName() string {
	return services.RemoteStartTransaction.String()
}

func NewEquipRemoteStartTransactionCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipRemoteStartTransactionCallbackRequest {
	req := &equipRemoteStartTransactionCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.RemoteStartTransaction.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipRemoteStartTransactionCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *callback.CallbackError) *equipRemoteStartTransactionCallbackRequest {
	req := &equipRemoteStartTransactionCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.RemoteStartTransaction.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
	}
	return req
}

var _ services.Response = &equipRemoteStartTransactionCallbackResponse{}

type equipRemoteStartTransactionCallbackResponse struct {
	api.Response
}

func (resp *equipRemoteStartTransactionCallbackResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipRemoteStartTransactionCallbackResponse) GetMsg() string {
	return resp.Msg
}

func RemoteStartTransactionCallbackRequest(ctx context.Context, req services.CallbackRequest) error {
	headerValue := make([]string, 0)
	//headerValue = append(headerValue, api.Services, "remote", "start", services.Callback)
	headerValue = append(headerValue, api.Services)
	headerValue = append(headerValue, services.RemoteStartTransaction.Split()...)
	headerValue = append(headerValue, services.Callback)

	header := map[string]string{api.Perms: strings.Join(headerValue, ":")}
	url := config.App.ServicesUrl + services.Equip + "/" + services.Callback + "/" + req.GetName() + "Callback"

	message, err := api.SendRequest(ctx, url, req, header)
	if err != nil {
		return err
	}

	resp := &equipRemoteStartTransactionCallbackResponse{}
	err = json.Unmarshal(message, resp)
	if err != nil {
		return err
	}

	if resp.Status == 1 {
		return errors.New(resp.Msg)
	}

	return nil
}

func RemoteStartTransactionCallbackRequestWithGeneric(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.RemoteStartTransaction)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipRemoteStartTransactionCallbackResponse{})
}
