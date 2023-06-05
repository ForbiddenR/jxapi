package equip

import (
	"context"

	services "github.com/ForbiddenR/jxapi/jxservices"
	"github.com/Kotodian/gokit/datasource/rabbitmq"
)

const updateTransactionQueue = services.QueuePrefix + "transaction"

type equipUpdateTransactionRequest struct {
	services.Base
	Data *equipUpdateTransactionReqeustDetail `json:"data"`
}

func (equipUpdateTransactionRequest) GetName() string {
	return services.UpdateTransaction.String()
}

type equipUpdateTransactionReqeustDetail struct {
	TransactionId string      `json:"transactionId"`
	EvseId        *string     `json:"evseSerial"`
	ConnectorId   string      `json:"connectorSerial"`
	Offline       bool        `json:"offline"`
	Timestamp     int64       `json:"timestamp"`
	MeterValue    *MeterValue `json:"MeterValue"`
	Tariff        Tariff      `json:"tariff"`
	ChargingState uint8       `json:"chargingState"`
	RemainingTime *int        `json:"remainingTime"`
	VIN           *string     `json:"vin"`
}

func NewUpdateTransactionRequest(sn, pod, msgID string, p *services.Protocol, transactionId, connectorId string, offline bool, timestamp int64, chargeState uint8) *equipUpdateTransactionRequest {
	updateTransaction := &equipUpdateTransactionRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.UpdateTransaction.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipUpdateTransactionReqeustDetail{
			TransactionId: transactionId,
			ConnectorId:   connectorId,
			Offline:       offline,
			Timestamp:     timestamp,
			MeterValue:    make([]MeterValue, 0),
			Tariff:        Tariff{},
			ChargingState: chargeState,
		},
	}

	return updateTransaction
}

func UpdateTransactionReqeust(req services.Request) error {
	ctx := context.Background()
	err := rabbitmq.Publish(ctx, updateTransactionQueue, nil, req)
	if err != nil {
		return err
	}
	return nil
}
