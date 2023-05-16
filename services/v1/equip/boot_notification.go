package equip

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"gitee.com/csms/jxeu-ocpp/internal/config"
	"gitee.com/csms/jxeu-ocpp/pkg/api"
	"gitee.com/csms/jxeu-ocpp/pkg/api/services"
)

type equipBootNotificationRequest struct {
	services.Base
	Data *equipBootNotificationRequestDetail `json:"data"`
}

type equipBootNotificationRequestDetail struct {
	ModelCode        string  `json:"modelCode"`
	ManufacturerCode string  `json:"manufacturerCode"`
	FirmwareVersion  *string `json:"firmwareVersion"`
	Iccid            *string `json:"iccid"`
	Imsi             *string `json:"imsi"`
}

func (r *equipBootNotificationRequest) GetName() string {
	return services.BootNotification.String()
}

func NewEquipBootNotificationRequest(sn, pod, msgID string, p *services.Protocol) *equipBootNotificationRequest {
	request := &equipBootNotificationRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.BootNotification.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
	}
	request.Data = &equipBootNotificationRequestDetail{}
	return request
}

var _ services.Response = &equipBootNotificationResponse{}

type equipBootNotificationResponse struct {
	api.Response
}

func (resp *equipBootNotificationResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipBootNotificationResponse) GetMsg() string {
	return resp.Msg
}

func BootNotificationRequest(ctx context.Context, req *equipBootNotificationRequest) error {
	headerValue := make([]string, 0)
	headerValue = append(headerValue, api.Services, services.Equip)
	headerValue = append(headerValue, services.Register.Split()...)
	header := map[string]string{api.Perms: strings.Join(headerValue, ":")}

	url := config.App.ServicesUrl + services.Equip + "/" + req.GetName()

	message, err := api.SendRequest(ctx, url, req, header)
	if err != nil {
		return err
	}

	resp := &equipBootNotificationResponse{}
	err = json.Unmarshal(message, resp)
	if err != nil {
		return err
	}

	if resp.Status == 1 {
		return errors.New(resp.Msg)
	}

	return nil
}

func BootNotificationRequestWithGeneric(ctx context.Context, req *equipBootNotificationRequest) error {
	header := services.GetSimpleHeaderValue(services.Register)

	url := services.GetSimpleURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipBootNotificationResponse{})
}
