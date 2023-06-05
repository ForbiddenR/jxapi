package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type equipStartTransactionRequest struct {
	services.Base
	Data *equipStartTransactionRequestDetail `json:"data"`
}

type equipStartTransactionRequestDetail struct {
	IdTokenType IdTokenType `json:"idTokenType"`
	// IdToken         string `json:"idToken"`
	MeterStart          *int        `json:"meterStart"`
	EvseSerial          *string     `json:"evseSerial"`
	ConnectorSerial     string      `json:"connectorSerial"`
	ReservationId       *int64      `json:"reservationId"`
	TransactionId       *string     `json:"transactionId"`
	RemoteStartId       *int64      `json:"remoteStartId"`
	Offline             *bool       `json:"offline"`
	Timestamp           int64       `json:"timestamp"`
	MeterValue          *MeterValue `json:"meterValue"`
	Tariff              *Tariff     `json:"tariff"`
	ChargingState       *uint8      `json:"chargingState"`
	Vin                 *string     `json:"vin"`
	RemainingTime       *int        `json:"remainingTime"`
	ChargingProfileId   *int64      `json:"chargingProfileId"`
	ChargingProfileUnit *uint8      `json:"chargingProfileUnit"`
}

func (equipStartTransactionRequest) GetName() string {
	return services.StartTransaction.String()
}

func NewEquipStartTransactionRequest(sn, pod, msgID string, p *services.Protocol,
	idToken string, connectorId string,
	timestamp int64) *equipStartTransactionRequest {
	return &equipStartTransactionRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.StartTransaction.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipStartTransactionRequestDetail{
			IdTokenType: IdTokenType{
				IdToken: idToken,
			},
			ConnectorSerial: connectorId,
			Timestamp:       timestamp,
			MeterValue:      &MeterValue{},
			Tariff:          &Tariff{},
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

func StartTransactionRequestWithGeneric(ctx context.Context, req services.Request) (*equipStartTransactionResponse, error) {
	header := services.GetSimpleHeaderValue(services.StartTransaction)

	url := services.GetSimpleURL(req)

	return services.RequestWithResponse(ctx, req, url, header, &equipStartTransactionResponse{})
}
