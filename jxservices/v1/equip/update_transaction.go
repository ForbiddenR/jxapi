package equip

import (
	"context"
	"encoding/json"
	"time"

	services "github.com/ForbiddenR/jxapi/jxservices"

	"github.com/makasim/amqpextra/publisher"
	amqp "github.com/rabbitmq/amqp091-go"
)

const updateTransactionQueue = services.QueuePrefix + "transaction"

var _ services.Request = &equipUpdateTransactionRequest{}

type equipUpdateTransactionRequest struct {
	services.Base
	Data *equipUpdateTransactionReqeustDetail `json:"data"`
}

func (equipUpdateTransactionRequest) GetName() services.Request2ServicesNameType {
	return services.UpdateTransaction
}

func (equipUpdateTransactionRequest) IsCallback() bool {
	return false
}

type equipUpdateTransactionReqeustDetail struct {
	Actime        int64       `json:"actime"`
	TransactionId string      `json:"transactionId"`
	EvseId        *string     `json:"evseSerial"`
	ConnectorId   string      `json:"connectorSerial"`
	Offline       bool        `json:"offline"`
	Timestamp     int64       `json:"timestamp"`
	MeterValue    *MeterValue `json:"MeterValue"`
	Tariff        *Tariff     `json:"tariff"`
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
			Actime:        time.Now().Unix(),
			TransactionId: transactionId,
			ConnectorId:   connectorId,
			Offline:       offline,
			Timestamp:     timestamp,
			MeterValue:    &MeterValue{},
			Tariff: &Tariff{
				Id: -1,
			},
			ChargingState: chargeState,
		},
	}

	return updateTransaction
}

func UpdateTransactionReqeust(req services.Request, p *publisher.Publisher) error {
	ctx := context.Background()
	bytes, err := json.Marshal(req)
	if err != nil {
		return err
	}
	message := publisher.Message{
		Context: ctx,
		Key:     updateTransactionQueue,
		Publishing: amqp.Publishing{
			ContentType: "application/json",
			Body:        bytes,
		},
	}
	err = p.Publish(message)
	if err != nil {
		return err
	}
	return nil
	// ctx := context.Background()
	// err := rabbitmq.Publish(ctx, updateTransactionQueue, nil, req)
	// if err != nil {
	// 	return err
	// }
	// return nil
}
