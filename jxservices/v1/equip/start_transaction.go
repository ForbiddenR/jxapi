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
	IdTokenType         IdTokenType `json:"idTokenType"`
	MeterStart          *int        `json:"meterStart"`
	EvseSerial          *string     `json:"evseSerial"`
	ConnectorSerial     string      `json:"connectorSerial"`
	ReservationId       *int64      `json:"reservationId"`
	TransactionId       *string     `json:"transactionId"`
	RemoteStartId       *int64      `json:"remoteStartId"`
	Offline             bool        `json:"offline"`
	Timestamp           int64       `json:"timestamp"`
	MeterValue          *MeterValue `json:"meterValue"`
	Tariff              *Tariff     `json:"tariff"`
	ChargingState       uint8       `json:"chargingState"`
	Vin                 *string     `json:"vin"`
	RemainingTime       *int        `json:"remainingTime"`
	ChargingProfileId   *int64      `json:"chargingProfileId"`
	ChargingProfileUnit *uint8      `json:"chargingProfileUnit"`
}

type StartTransactionRequestConfig struct {
	services.ReusedConfig
	IdToken       string
	ConnectorId   string
	Timestamp     int64
	Offline       bool
	ChargingState uint8
}

func NewEquipStartTransactionRequestWithConfig(config *StartTransactionRequestConfig) *equipStartTransactionRequest {
	req :=  &equipStartTransactionRequest{
		Base: services.Base{
			EquipmentSn: config.Sn,
			Protocol:    config.Protocol,
			Category:    services.StartTransaction.FirstUpper(),
			AccessPod:   config.Pod,
			MsgID:       config.MsgID,
		},
		Data: &equipStartTransactionRequestDetail{
			IdTokenType: IdTokenType{
				IdToken: config.IdToken,
			},
			ConnectorSerial: config.ConnectorId,
			Timestamp:       config.Timestamp,
			Offline:         config.Offline,
			ChargingState:   config.ChargingState,
			Tariff: &Tariff{
				Id: -1,
			},
		},
	}
	
	if !config.Protocol.Equal(services.OCPP16()) {
		req.Data.MeterValue = &MeterValue{}
	}
	return req
}

func (equipStartTransactionRequest) GetName() string {
	return services.StartTransaction.String()
}

func NewEquipStartTransactionRequest(sn, pod, msgID string, p *services.Protocol,
	idToken string, connectorId string,
	timestamp int64) *equipStartTransactionRequest {
	req := &equipStartTransactionRequest{
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
		},
	}
	req.Data.Tariff = &Tariff{
		Id: -1,
	}
	if !p.Equal(services.OCPP16()) {
		req.Data.MeterValue = &MeterValue{}
	}
	return req
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

func StartTransactionRequest(ctx context.Context, req services.Request) (*equipStartTransactionResponse, error) {
	header := services.GetSimpleHeaderValue(services.StartTransaction)

	url := services.GetSimpleURL(req)

	return services.RequestWithResponse(ctx, req, url, header, &equipStartTransactionResponse{})
}
