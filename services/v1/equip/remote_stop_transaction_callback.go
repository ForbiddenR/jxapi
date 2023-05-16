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

type equipRemoteStopTransactionCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (s *equipRemoteStopTransactionCallbackRequest) GetName() string {
	return services.RemoteStopTransaction.String()
}

func NewEquipRemoteStopTransactionCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipRemoteStopTransactionCallbackRequest {
	req := &equipRemoteStopTransactionCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.RemoteStopTransaction.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipRemoteStopTransactionCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *callback.CallbackError) *equipRemoteStopTransactionCallbackRequest {
	req := &equipRemoteStopTransactionCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.RemoteStopTransaction.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
	}
	return req
}

var _ services.Response = &equipRemoteStopTransactionCallbackResponse{}

type equipRemoteStopTransactionCallbackResponse struct {
	api.Response
}

func (resp *equipRemoteStopTransactionCallbackResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipRemoteStopTransactionCallbackResponse) GetMsg() string {
	return resp.Msg
}

func RemoteStopTransactionRequest(ctx context.Context, req services.CallbackRequest) error {
	headerValue := make([]string, 0)
	//headerValue = append(headerValue, api.Services, "remote", "stop", "transaction", services.Callback)
	headerValue = append(headerValue, api.Services)
	headerValue = append(headerValue, services.RemoteStopTransaction.Split()...)
	headerValue = append(headerValue, services.Callback)

	header := map[string]string{api.Perms: strings.Join(headerValue, ":")}

	url := config.App.ServicesUrl + services.Equip + "/" + services.Callback + "/" + req.GetName() + "Callback"

	message, err := api.SendRequest(ctx, url, req, header)

	if err != nil {
		return err
	}

	resp := &equipRemoteStopTransactionCallbackResponse{}
	err = json.Unmarshal(message, resp)

	if err != nil {
		return err
	}

	if resp.Status == 1 {
		return errors.New(resp.Msg)
	}

	return nil
}

func RemoteStopTransactionRequestWithGeneric(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.RemoteStopTransaction)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipRemoteStopTransactionCallbackResponse{})
}
