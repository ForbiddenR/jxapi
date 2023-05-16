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

type equipUpdateFirmwareCallbackRequest struct {
	services.Base
	Callback services.CB                               `json:"callback"`
	Data     *equipUpdateFirmwareCallbackRequestDetail `json:"data"`
}

type equipUpdateFirmwareCallbackRequestDetail struct {
}

func (*equipUpdateFirmwareCallbackRequest) GetName() string {
	return services.UpdateFirmware.String()
}

func NewEquipUpdateFirmwareCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipUpdateFirmwareCallbackRequest {
	return &equipUpdateFirmwareCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.UpdateFirmware.FirstUpper(),
			AccessPod:   pod,

			MsgID: msgID,
		},
		Callback: services.NewCB(status),
	}
}

func NewEquipUpdateFirmwareCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *callback.CallbackError) *equipUpdateFirmwareCallbackRequest {
	return &equipUpdateFirmwareCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.UpdateFirmware.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
		Data:     nil,
	}
}

var _ services.Response = &equipUpdateFirmwareCallbackResponse{}

type equipUpdateFirmwareCallbackResponse struct {
	api.Response
}

func (resp *equipUpdateFirmwareCallbackResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipUpdateFirmwareCallbackResponse) GetMsg() string {
	return resp.Msg
}

func UpdateFirmwareCallbackRequest(ctx context.Context, req services.CallbackRequest) error {
	headerValue := make([]string, 0)
	headerValue = append(headerValue, api.Services)
	headerValue = append(headerValue, services.UpdateFirmware.Split()...)
	headerValue = append(headerValue, services.Callback)

	header := map[string]string{api.Perms: strings.Join(headerValue, ":")}

	url := config.App.ServicesUrl + services.Equip + "/" + services.Callback + "/" + req.GetName() + "Callback"

	message, err := api.SendRequest(ctx, url, req, header)
	if err != nil {
		return err
	}

	response := &equipUpdateFirmwareCallbackResponse{}

	if err = json.Unmarshal(message, response); err != nil {
		return err
	}

	if response.Status == 1 {
		return errors.New(response.Msg)
	}

	return nil
}

func UpdateFirmwareCallbackRequestWithGeneric(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.UpdateFirmware)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipUpdateFirmwareCallbackResponse{})
}
