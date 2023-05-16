package equip

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"gitee.com/csms/jxeu-ocpp/internal/config"
	"gitee.com/csms/jxeu-ocpp/pkg/api"
	"gitee.com/csms/jxeu-ocpp/pkg/api/services"
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

//type StartTransactionResponse interface {
//	startTransaction()
//}

func StartTransactionOCPP16Request(ctx context.Context, req *equipStartTransactionOCPP16Request) (*equipStartTransactionResponse, error) {
	headerValue := make([]string, 0)
	headerValue = append(headerValue, api.Services)
	headerValue = append(headerValue, services.StartTransaction.Split()...)

	header := map[string]string{api.Perms: strings.Join(headerValue, ":")}

	url := config.App.ServicesUrl + services.Equip + "/" + req.GetName()

	message, err := api.SendRequest(ctx, url, req, header)

	if err != nil {
		return nil, err
	}

	resp := &equipStartTransactionResponse{}

	err = json.Unmarshal(message, resp)

	if err != nil {
		return nil, err
	}

	if resp.Status == 1 {
		return nil, errors.New(resp.Msg)
	}

	return resp, nil
}

func StartTransactionOCPP16RequestWithGeneric(ctx context.Context, req services.Request) (*equipStartTransactionResponse, error) {
	header := services.GetSimpleHeaderValue(services.StartTransaction)

	url := services.GetSimpleURL(req)

	return services.RequestWithResponse(ctx, req, url, header, &equipStartTransactionResponse{})
}
