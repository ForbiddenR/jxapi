package equip

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"gitee.com/csms/jxeu-ocpp/internal/config"
	api "github.com/ForbiddenR/jx-api"
	"github.com/ForbiddenR/jx-api/apierrors"
	"github.com/ForbiddenR/jx-api/services"
)

type equipResetCallbackRequest struct {
	services.Base
	Callback services.CB                      `json:"callback"`
	Data     *equipResetCallbackRequestDetail `json:"data"`
}

type equipResetCallbackRequestDetail struct {
}

func (r *equipResetCallbackRequest) GetName() string {
	return services.Reset.String()
}

func NewEquipResetCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipResetCallbackRequest {
	req := &equipResetCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.Reset.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
	}

	return req
}

func NewEquipResetCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipResetCallbackRequest {
	req := &equipResetCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.Reset.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
	}
	return req
}

var _ services.Response = &equipResetResponse{}

type equipResetResponse struct {
	api.Response
}

func (resp *equipResetResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipResetResponse) GetMsg() string {
	return resp.Msg
}

type equipResetResponseDetail struct {
}

func ResetCallbackRequest(ctx context.Context, req services.CallbackRequest) error {
	headerValue := make([]string, 0)
	//headerValue = append(headerValue, "services", "equip", "reset", services.Callback)
	headerValue = append(headerValue, api.Services, "equip")
	headerValue = append(headerValue, services.Reset.Split()...)
	headerValue = append(headerValue, services.Callback)

	header := map[string]string{api.Perms: strings.Join(headerValue, ":")}

	url := config.App.ServicesUrl + services.Equip + "/" + services.Callback + "/" + req.GetName() + "Callback"

	message, err := api.SendRequest(ctx, url, req, header)

	if err != nil {
		return err
	}

	resp := &equipResetResponse{}
	err = json.Unmarshal(message, resp)
	if err != nil {
		return err
	}

	if resp.Status == 1 {
		return errors.New(resp.Msg)
	}

	return nil
}

func ResetCallbackRequestWithGeneric(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.Reset)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipResetResponse{})
}
