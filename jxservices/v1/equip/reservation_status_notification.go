package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type equipReservationStatusNotificationRequest struct {
	services.Base
	Data *equipReservationStatusNotificationRequestDetail `json:"data"`
}

type equipReservationStatusNotificationRequestDetail struct {
	ReservationId int `json:"reservationId"`
	Status        int `json:"status"`
}

func (r equipReservationStatusNotificationRequest) GetName() string {
	return services.ReservationStatusNotification.String()
}

func NewEquipReservationStatusNotification(sn, pod, msgID string, p *services.Protocol) *equipReservationStatusNotificationRequest {
	return &equipReservationStatusNotificationRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.ReservationStatusNotification.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
	}
}

type equipReservationStatusNotificationResponse struct {
	api.Response
	Data *equipAuthorizeTransactionRequestDetail `json:"data"`
}

type equipReservationStatusNotificationResponseDetail struct {
}

func (resp *equipReservationStatusNotificationResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipReservationStatusNotificationResponse) GetMsg() string {
	return resp.Msg
}

func ReservationStatusNotificationRequest(ctx context.Context, req *equipReservationStatusNotificationRequest) error {
	header := services.GetSimpleHeaderValue(services.ReservationStatusNotification)

	url := services.GetSimpleURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipReservationStatusNotificationResponse{})
}
