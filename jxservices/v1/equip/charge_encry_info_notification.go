package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type equipChargeEncryInfoNotificationRequest struct {
	services.Base
	Data *equipChargeEncryInfoNotificationRequestDetail `json:"data"`
}

type equipChargeEncryInfoNotificationRequestDetail struct {
}

func (equipChargeEncryInfoNotificationRequest) GetName() string {
	return services.ChargeEncryInfoNotification.String()
}

func NewEquipChargeEncryInfoNotificationRequest(sn, pod, msgID string, p *services.Protocol) *equipChargeEncryInfoNotificationRequest {
	req := &equipChargeEncryInfoNotificationRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.ChargeEncryInfoNotification.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipChargeEncryInfoNotificationRequestDetail{},
	}
	return req
}

var _ services.Response = &equipChargeEncryInfoNotificationResponse{}

type equipChargeEncryInfoNotificationResponse struct {
	api.Response
	Data *equipChargeEncryInfoNotificationResponseDetail `json:"data"`
}

func (resp *equipChargeEncryInfoNotificationResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipChargeEncryInfoNotificationResponse) GetMsg() string {
	return resp.Msg
}

type equipChargeEncryInfoNotificationResponseDetail struct {
}

func ChargeEncryInfoNotificationReqeust(ctx context.Context, req services.Request) error {
	header := services.GetSimpleHeaderValue(services.ChargeEncryInfoNotification)

	url := services.GetSimpleURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipChargeEncryInfoNotificationResponse{})
}