package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type equipChargingTimerNotificationRequest struct {
	services.Base
	Data *equipChargingTimerNotificationRequestDetail `json:"data"`
}

type ChargingTimerStatus int

const (
	ChargingTimerStatusCharging ChargingTimerStatus = 1
	ChargingTimerStatusExpired  ChargingTimerStatus = 2
)

type equipChargingTimerNotificationRequestDetail struct {
	Status   ChargingTimerStatus `json:"status"`
	TimerId  int64               `json:"timerId"`
	Charging *Charging           `json:"charging"`
}

type Charging struct {
	EvseSerial      *string `json:"evseSerial"`
	ConnectorSerial string  `json:"connectorSerial"`
	TransactionId   *string `json:"transactionId"`
	TriggerTime     *int64  `json:"timestamp"`
	Version         *int64  `json:"version"`
}

func (*equipChargingTimerNotificationRequest) GetName() string {
	return services.ChargingTimerNotification.String()
}

func NewEquipChargingTimerNotificationRequest(sn, pod, msgID string, connectorId string, timerId int64, status ChargingTimerStatus) *equipChargingTimerNotificationRequest {
	return &equipChargingTimerNotificationRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    services.OCPP16(),
			Category:    services.ChargingTimerNotification.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipChargingTimerNotificationRequestDetail{
			Charging: &Charging{
				ConnectorSerial: connectorId,
			},
			TimerId:         timerId,
			Status:          status,
		},
	}
}

var _ services.Response = &equipChargingTimerNotificationResponse{}

type equipChargingTimerNotificationResponse struct {
	api.Response
}

func (resp *equipChargingTimerNotificationResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipChargingTimerNotificationResponse) GetMsg() string {
	return resp.Msg
}

func ChargingTimerNotificationRequestWithGeneric(ctx context.Context, req services.Request) error {
	header := services.GetSimpleHeaderValue(services.ChargingTimerNotification)

	url := services.GetSimpleURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipChargingTimerNotificationResponse{})
}
