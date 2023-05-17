package equip

import (
	"context"

	api "github.com/ForbiddenR/jx-api"
	"github.com/ForbiddenR/jx-api/services"
)

// This feature only supports for the version "OCPP1.6".

type StopReasonTypeMenu int

const (
	StopReasonTypeDeAuthorized   StopReasonTypeMenu = 1
	StopReasonTypeEVDisconnected                    = StopReasonTypeDeAuthorized + 1
	StopReasonTypeEmergencyStop                     = StopReasonTypeEVDisconnected + 1
	StopReasonTypeHardReset                         = StopReasonTypeEmergencyStop + 1
	StopReasonTypeLocal                             = StopReasonTypeHardReset + 1
	StopReasonTypeOther                             = StopReasonTypeLocal + 1
	StopReasonTypePowerLoss                         = StopReasonTypeOther + 1
	StopReasonTypeReboot                            = StopReasonTypePowerLoss + 1
	StopReasonTypeRemote                            = StopReasonTypeReboot + 1
	StopReasonTypeSoftReset                         = StopReasonTypeRemote + 1
	StopReasonTypeUnlockCommand                     = StopReasonTypeSoftReset + 1
)

// // OCPP16StopReason is a transfer function turning "protocol.StopTransaction" into "StopReasonTypeMenu".
// func OCPP16StopReason(r protocol.StopTransactionJsonReason) StopReasonTypeMenu {
// 	switch r {
// 	case protocol.StopTransactionJsonReasonDeAuthorized:
// 		return StopReasonTypeDeAuthorized
// 	case protocol.StopTransactionJsonReasonEVDisconnected:
// 		return StopReasonTypeEVDisconnected
// 	case protocol.StopTransactionJsonReasonEmergencyStop:
// 		return StopReasonTypeEmergencyStop
// 	case protocol.StopTransactionJsonReasonHardReset:
// 		return StopReasonTypeHardReset
// 	case protocol.StopTransactionJsonReasonLocal:
// 		return StopReasonTypeLocal
// 	case protocol.StopTransactionJsonReasonOther:
// 		return StopReasonTypeOther
// 	case protocol.StopTransactionJsonReasonPowerLoss:
// 		return StopReasonTypePowerLoss
// 	case protocol.StopTransactionJsonReasonReboot:
// 		return StopReasonTypeReboot
// 	case protocol.StopTransactionJsonReasonRemote:
// 		return StopReasonTypeRemote
// 	case protocol.StopTransactionJsonReasonSoftReset:
// 		return StopReasonTypeSoftReset
// 	case protocol.StopTransactionJsonReasonUnlockCommand:
// 		return StopReasonTypeUnlockCommand
// 	default:
// 		return StopReasonTypeOther
// 	}
// }

type equipStopTransactionOCPP16Request struct {
	services.Base
	Data *equipStopTransactionRequestDetail `json:"data"`
}

type equipStopTransactionRequestDetail struct {
	// IdToken can be null if the stop events are caused by Charging Station itself.
	IdToken       *string                            `json:"idToken,omitempty"`
	MeterStop     int                                `json:"meterStop"`
	Reason        StopReasonTypeMenu                 `json:"reason"`
	TransactionId string                             `json:"transactionId"`
	IsOffline     bool                               `json:"offline"`
	Timestamp     int64                              `json:"timestamp"`
	MeterValue    *equipMeterValuesRequestMeterValue `json:"meterValue"`
}

//type equipStopTransactionRequestMeterValue struct {
//	// Timestamp is the time when the Charging Station gets meter datas.
//	Timestamp int64 `json:"timestamp"`
//	// SampledValue uses the type, "MeterValueElemSampledValueElem", indirectly.
//	SampledValue []protocol.MeterValueElemSampledValueElem `json:"sampledValue"`
//}

//type equipStopTransactionRequestSampledValue struct {
//	Value     string                                           `json:"value"`
//	Context   protocol.MeterValueElemSampledValueElemContext   `json:"context"`
//	//Format  protocol.
//	Measurand protocol.MeterValueElemSampledValueElemMeasurand `json:"measurand"`
//	Phase     *protocol.MeterValueElemSampledValueElemPhase    `json:"phase"`
//	Location  protocol.MeterValueElemSampledValueElemLocation  `json:"location"`
//	Unit      protocol.MeterValueElemSampledValueElemUnit      `json:"unit"`
//}

//func NewEquipStopTransactionRequestSampledValue(value string,
//	context protocol.MeterValueElemSampledValueElemContext,
//	measurand protocol.MeterValueElemSampledValueElemMeasurand,
//	location protocol.MeterValueElemSampledValueElemLocation,
//	unit protocol.MeterValueElemSampledValueElemUnit) *equipStopTransactionRequestSampledValue {
//	return &equipStopTransactionRequestSampledValue{
//		Value:     value,
//		Context:   context,
//		Measurand: measurand,
//		Location:  location,
//		Unit:      unit,
//	}
//}

func (r equipStopTransactionOCPP16Request) GetName() string {
	return services.StopTransaction.String()
}

func NewEquipStopTransactionOCPP16Request(sn, pod string, msgID string,
	meterStop int, reason StopReasonTypeMenu, transactionId string, isOffline bool, timestamp int64) *equipStopTransactionOCPP16Request {
	request := &equipStopTransactionOCPP16Request{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    services.OCPP16(),
			Category:    services.StopTransaction.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipStopTransactionRequestDetail{
			MeterStop:     meterStop,
			Reason:        reason,
			TransactionId: transactionId,
			Timestamp:     timestamp,
			IsOffline:     isOffline,
			MeterValue: &equipMeterValuesRequestMeterValue{
				SampledValue: nil,
			},
		},
	}
	sampledValue := make([]MeterValueElemSampledValueElem, 0)

	request.Data.MeterValue.SampledValue = sampledValue
	return request
}

var _ services.Response = &equipStopTransactionOCPP16Response{}

type equipStopTransactionOCPP16Response struct {
	api.Response
	Data *equipStopTransactionOCPP16ResponseDetail `json:"data"`
}

func (resp *equipStopTransactionOCPP16Response) GetStatus() int {
	return resp.Status
}

func (resp *equipStopTransactionOCPP16Response) GetMsg() string {
	return resp.Msg
}

type equipStopTransactionOCPP16ResponseDetail struct {
	IdTokenInfo
}

func StopTransactionRequestWithGeneric(ctx context.Context, req *equipStopTransactionOCPP16Request) (*equipStopTransactionOCPP16Response, error) {
	header := services.GetSimpleHeaderValue(services.StopTransaction)

	url := services.GetSimpleURL(req)

	return services.RequestWithResponse(ctx, req, url, header, &equipStopTransactionOCPP16Response{})
}
