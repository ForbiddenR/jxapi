package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type ReservationStatusNotificationRequestStautsType int

const (
	ReservationStatusNotificationRequestStautsTypeExpired ReservationStatusNotificationRequestStautsType = 1
	ReservationStatusNotificationRequestStautsTypeRemoved ReservationStatusNotificationRequestStautsType = 2
)

type equipReservationStatusNotificationRequest struct {
	services.Base
	Data *equipReservationStatusNotificationRequestDetail `json:"data"`
}

type equipReservationStatusNotificationRequestDetail struct {
	ReservationId uint64                                         `json:"reservationId"`
	Status        ReservationStatusNotificationRequestStautsType `json:"status"`
}

func (r equipReservationStatusNotificationRequest) GetName() string {
	return services.ReservationStatusNotification.String()
}

func NewEquipReservationStatusNotification(sn, pod, msgID string, p *services.Protocol, reservationId uint64, status ReservationStatusNotificationRequestStautsType) *equipReservationStatusNotificationRequest {
	return &equipReservationStatusNotificationRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.ReservationStatusNotification.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipReservationStatusNotificationRequestDetail{
			ReservationId: reservationId,
			Status:        status,
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
