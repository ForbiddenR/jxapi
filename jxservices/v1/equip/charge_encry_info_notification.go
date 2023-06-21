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
	EvseId           *string `json:"evseSerial"`
	ConnectorId      string  `json:"connectorSerial"`
	TransactionId    string  `json:"transactionId"`
	EncryptedData    string  `json:"encryptedData"`
	MeterNum         string  `json:"meterNum"`
	ElecmeterVersion int64   `json:"elecmeterVersion"`
	EncryType        uint64  `json:"encryType"`
}

func (equipChargeEncryInfoNotificationRequest) GetName() string {
	return services.ChargeEncryInfoNotification.String()
}

func NewEquipChargeEncryInfoNotificationRequest(sn string, p *services.Protocol, pod, msgID string,
	connectorId, transactionId, encryptedData, meterNum string, elecmeterVersion int64, encryType uint64) *equipChargeEncryInfoNotificationRequest {
	req := &equipChargeEncryInfoNotificationRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.ChargeEncryInfoNotification.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipChargeEncryInfoNotificationRequestDetail{
			ConnectorId:      connectorId,
			TransactionId:    transactionId,
			EncryptedData:    encryptedData,
			MeterNum:         meterNum,
			ElecmeterVersion: elecmeterVersion,
			EncryType:        encryType,
		},
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
