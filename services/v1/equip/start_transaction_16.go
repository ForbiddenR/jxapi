package equip

import (
	"context"

	api "github.com/ForbiddenR/jx-api"
	"github.com/ForbiddenR/jx-api/services"
)

type equipStartTransactionOCPP16Request struct {
	services.Base
	Data *equipStartTransactionOCPP16RequestDetail `json:"data"`
}

type equipStartTransactionOCPP16RequestDetail struct {
	IdToken         string `json:"idToken"`
	MeterStart      int    `json:"meterStart"`
	ConnectorSerial string `json:"connectorSerial"`
	ReservationId   *int64 `json:"reservationId"`
	Timestamp       int64  `json:"timestamp"`
}

func (equipStartTransactionOCPP16Request) GetName() string {
	return services.StartTransaction.String()
}

func NewEquipStartTransactionRequestOCPP16(sn, pod string, msgID,
	idToken string, meterStart int, connectorId string,
	timestamp int64) *equipStartTransactionOCPP16Request {
	return &equipStartTransactionOCPP16Request{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    services.OCPP16(),
			Category:    services.StartTransaction.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipStartTransactionOCPP16RequestDetail{
			IdToken:         idToken,
			MeterStart:      meterStart,
			ConnectorSerial: connectorId,
			Timestamp:       timestamp,
		},
	}
}

var _ services.Response = &equipStartTransactionResponse{}

type equipStartTransactionResponse struct {
	api.Response
	Data *equipStartTransactionResponseDetail `json:"data"`
}

func (resp *equipStartTransactionResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipStartTransactionResponse) GetMsg() string {
	return resp.Msg
}

type equipStartTransactionResponseDetail struct {
	TransactionId string      `json:"transactionId"`
	IdTokenInfo   IdTokenInfo `json:"idTokenInfo"`
}

func StartTransactionOCPP16RequestWithGeneric(ctx context.Context, req services.Request) (*equipStartTransactionResponse, error) {
	header := services.GetSimpleHeaderValue(services.StartTransaction)

	url := services.GetSimpleURL(req)

	return services.RequestWithResponse(ctx, req, url, header, &equipStartTransactionResponse{})
}
