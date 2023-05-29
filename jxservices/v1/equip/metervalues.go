package equip

import (
	"context"

	// "gitee.com/csms/jxeu-ocpp/internal/rabbitmq"
	services "github.com/ForbiddenR/jxapi/jxservices"
	"github.com/Kotodian/gokit/datasource/rabbitmq"
)

// This feature is transferred by rabbitmq.

const meterQueue = services.QueuePrefix + "metervalues"

type equipMeterValuesRequest struct {
	services.Base
	Data *equipMeterValuesRequestDetail `json:"data"`
}

type equipMeterValuesRequestDetail struct {
	EvseId        *string     `json:"evseSerial,omitempty"`
	TransactionId *int64      `json:"transactionId,omitempty"`
	ConnectorId   *string     `json:"connectorSerial,omitempty"`
	Timestamp     *int64      `json:"timestamp,omitempty"`
	MeterValue    *MeterValue `json:"meterValue"`
}

func (equipMeterValuesRequest) GetName() string {
	return services.MeterValues.String()
}

func NewEquipMeterValuesOCPP16Request(sn, pod, msgID string, connectorId string) *equipMeterValuesRequest {
	meterValue := &equipMeterValuesRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    services.OCPP16(),
			Category:    services.MeterValues.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipMeterValuesRequestDetail{
			ConnectorId: &connectorId,
			MeterValue:  &MeterValue{},
		},
	}
	meterValue.Data.MeterValue.SampleValue = make([]MeterValueElemSampledValueElem, 0)
	return meterValue
}

func NewEquipMeterValuesRequest(sn, pod, msgID string, p *services.Protocol, evseId string, timestamp int64) *equipMeterValuesRequest {
	meterValue := &equipMeterValuesRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.MeterValues.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipMeterValuesRequestDetail{
			EvseId: &evseId,
			MeterValue: &MeterValue{
				Timestamp: timestamp,
			},
		},
	}
	meterValue.Data.MeterValue.SampleValue = make([]MeterValueElemSampledValueElem, 0)
	return meterValue
}

func MeterValuesRequest(req *equipMeterValuesRequest) error {
	ctx := context.Background()
	err := rabbitmq.Publish(ctx, meterQueue, nil, req)
	if err != nil {
		return err
	}
	return nil
}
