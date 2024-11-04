package equip

import (
	"context"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type ReservationStatusNotificationRequestStautsType int

const (
	ReservationStatusNotificationRequestStautsTypeExpired ReservationStatusNotificationRequestStautsType = 1
	ReservationStatusNotificationRequestStautsTypeRemoved ReservationStatusNotificationRequestStautsType = 2
)

var _ services.Request = &equipReservationStatusNotificationRequest{}

type equipReservationStatusNotificationRequest struct {
	services.Base
	Data *equipReservationStatusNotificationRequestDetail `json:"data"`
}

type equipReservationStatusNotificationRequestDetail struct {
	ReservationId uint64                                         `json:"reservationId"`
	Status        ReservationStatusNotificationRequestStautsType `json:"status"`
}

func (equipReservationStatusNotificationRequest) GetName() services.Request2ServicesNameType {
	return services.ReservationStatusNotification
}

func (e *equipReservationStatusNotificationRequest) TraceId() string {
	return e.MsgID
}

func (equipReservationStatusNotificationRequest) IsCallback() bool {
	return false
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

// type equipReservationStatusNotificationResponse struct {
// 	api.Response
// 	Data *equipReservationStatusNotificationResponseDetail `json:"data"`
// }

// type equipReservationStatusNotificationResponseDetail struct {
// }

// func (resp *equipReservationStatusNotificationResponse) GetStatus() int {
// 	return resp.Status
// }

// func (resp *equipReservationStatusNotificationResponse) GetMsg() string {
// 	return resp.Msg
// }

func ReservationStatusNotificationRequest(ctx context.Context, req *equipReservationStatusNotificationRequest) error {
	return services.Transport(ctx, req)
}
