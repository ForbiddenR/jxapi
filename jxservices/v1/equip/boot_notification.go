package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

var _ services.Request = &equipBootNotificationRequest{}

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

func (equipBootNotificationRequest) GetName() services.Request2ServicesNameType {
	return services.BootNotification
}

func (equipBootNotificationRequest) IsCallback() bool {
	return false
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
	header := services.GetSimpleHeaderValue(services.Register)
	url := services.GetSimpleURL(req)
	return services.RequestWithoutResponse(ctx, req, url, header, &equipBootNotificationResponse{})
}
