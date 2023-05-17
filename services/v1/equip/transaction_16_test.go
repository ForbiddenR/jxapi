package equip

// import (
// 	"context"
// 	"testing"
// 	"time"

// 	"github.com/Kotodian/gokit/id"
// 	"github.com/stretchr/testify/assert"

// 	"gitee.com/csms/jxeu-ocpp/internal/config"
// 	"gitee.com/csms/jxeu-ocpp/internal/log"
// 	"gitee.com/csms/jxeu-ocpp/internal/rabbitmq"
// 	"gitee.com/csms/jxeu-ocpp/pkg/api"
// 	"gitee.com/csms/jxeu-ocpp/pkg/api/services"
// )

// func TestTransactionRequestWithGeneric(t *testing.T) {
// 	config.TestConfig()
// 	api.Init()
// 	rabbitmq.Init(&rabbitmq.Option{
// 		Url:      "10.43.0.14:5672",
// 		Username: "admin",
// 		Password: "3hCjyu5RXOX1xW6W",
// 	})
// 	log.InitNopLogger()
// 	ctx := context.TODO()

// 	idToken := "fbf8e1faa42c18ba98330b77ca97d29a"
// 	cid := "1"
// 	p := services.OCPP16()

// 	reqOnline := newTestEquipOnlineRequest(services.TestSN, p, services.TestAccessPod, id.Next().String())
// 	equipID, err := OnlineRequestWithGeneric(ctx, reqOnline)

// 	assert.Nil(t, err)
// 	assert.NotEmpty(t, equipID)

// 	reqStatusNotification := newTestEquipStatusNotificationRequest(p, cid, ConnectorStatusPreparing)
// 	err = StatusNotificationRequestWithGeneric(ctx, reqStatusNotification)

// 	assert.Nil(t, err)

// 	reqAuthorization := newTestEquipAuthorizeTransactionRequest(p, idToken)
// 	respAuthorization, err := AuthorizeTransactionRequestWithGeneric(ctx, reqAuthorization)

// 	assert.Nil(t, err)
// 	assert.NotNil(t, respAuthorization)
// 	assert.Equal(t, Accepted, respAuthorization.Data.IdTokenInfo.Status)

// 	startReq := newTestEquipStartTransactionOCPP16Request(0, idToken, cid)
// 	startResp, err := StartTransactionOCPP16RequestWithGeneric(ctx, startReq)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, startResp)
// 	assert.NotEmpty(t, startResp.Data.TransactionId)

// 	// phase
// 	pl1 := MeterValuesJsonMeterValueElemSampledValueElemPhaseL1
// 	pl2 := MeterValuesJsonMeterValueElemSampledValueElemPhaseL2
// 	pl3 := MeterValuesJsonMeterValueElemSampledValueElemPhaseL3
// 	// measurand
// 	curmea := MeterValueElemSampledValueElemMeasurandCurrentImport
// 	irmea := MeterValueElemSampledValueElemMeasurandEnergyActiveImportRegister
// 	vomea := MeterValueElemSampledValueElemMeasurandVoltage
// 	pwmea := MeterValueElemSampledValueElemMeasurandPowerActiveImport
// 	socmea := MeterValueElemSampledValueElemMeasurandSoC
// 	tempmea := MeterValueElemSampledValueElemMeasurandTemperature
// 	// context
// 	ctxBegin := MeterValueElemSampledValueElemContextTransactionBegin
// 	ctxSamplePeriodic := MeterValueElemSampledValueElemContextSamplePeriodic
// 	ctxEnd := MeterValueElemSampledValueElemContextTransactionEnd
// 	// location
// 	loout := MeterValueElemSampledValueElemLocationOutlet
// 	loev := MeterValueElemSampledValueElemLocationEV
// 	// unit
// 	unitA := MeterValueElemSampledValueElemUnitA
// 	unitWh := MeterValueElemSampledValueElemUnitWh
// 	unitV := MeterValueElemSampledValueElemUnitV
// 	unitW := MeterValueElemSampledValueElemUnitW
// 	unitPercent := MeterValueElemSampledValueElemUnitPercent
// 	unitCelcius := MeterValueElemSampledValueElemUnitCelcius
// 	meterValue := &equipMeterValuesRequestMeterValue{
// 		SampledValue: []MeterValueElemSampledValueElem{
// 			{
// 				Context:   &ctxBegin,
// 				Phase:     &pl1,
// 				Measurand: &curmea,
// 				Value:     "10.00",
// 				Location:  &loout,
// 				Unit:      &unitA,
// 			},
// 			{
// 				Context:   &ctxBegin,
// 				Phase:     &pl2,
// 				Measurand: &curmea,
// 				Value:     "10.00",
// 				Location:  &loout,
// 				Unit:      &unitA,
// 			},
// 			{
// 				Context:   &ctxBegin,
// 				Phase:     &pl3,
// 				Measurand: &curmea,
// 				Value:     "10.00",
// 				Location:  &loout,
// 				Unit:      &unitA,
// 			},
// 			{
// 				Context:   &ctxBegin,
// 				Measurand: &irmea,
// 				Value:     "10.00",
// 				Location:  &loout,
// 				Unit:      &unitWh,
// 			},
// 			{
// 				Context:   &ctxBegin,
// 				Measurand: &vomea,
// 				Value:     "10.00",
// 				Location:  &loout,
// 				Unit:      &unitV,
// 				Phase:     &pl1,
// 			},
// 			{
// 				Context:   &ctxBegin,
// 				Measurand: &vomea,
// 				Value:     "10.00",
// 				Location:  &loout,
// 				Unit:      &unitV,
// 				Phase:     &pl2,
// 			},
// 			{
// 				Context:   &ctxBegin,
// 				Measurand: &vomea,
// 				Value:     "10.00",
// 				Location:  &loout,
// 				Unit:      &unitV,
// 				Phase:     &pl3,
// 			},
// 			{
// 				Context:   &ctxBegin,
// 				Measurand: &pwmea,
// 				Value:     "10.00",
// 				Location:  &loout,
// 				Unit:      &unitW,
// 			},
// 			{
// 				Context:   &ctxBegin,
// 				Measurand: &socmea,
// 				Value:     "10",
// 				Location:  &loev,
// 				Unit:      &unitPercent,
// 			},
// 			{
// 				Context:   &ctxBegin,
// 				Measurand: &tempmea,
// 				Value:     "10",
// 				Location:  &loev,
// 				Unit:      &unitCelcius,
// 			},
// 		},
// 	}
// 	meterReq := newTestEquipMeterValuesOCPP16Request(cid, meterValue)
// 	err = MeterValuesRequest(meterReq)
// 	assert.Nil(t, err)

// 	time.Sleep(2 * time.Second)
// 	meterValue = &equipMeterValuesRequestMeterValue{
// 		SampledValue: []MeterValueElemSampledValueElem{
// 			{
// 				Context:   &ctxSamplePeriodic,
// 				Phase:     &pl1,
// 				Measurand: &curmea,
// 				Value:     "11.00",
// 				Location:  &loout,
// 				Unit:      &unitA,
// 			},
// 			{
// 				Context:   &ctxSamplePeriodic,
// 				Phase:     &pl2,
// 				Measurand: &curmea,
// 				Value:     "11.00",
// 				Location:  &loout,
// 				Unit:      &unitA,
// 			},
// 			{
// 				Context:   &ctxSamplePeriodic,
// 				Phase:     &pl3,
// 				Measurand: &curmea,
// 				Value:     "11.00",
// 				Location:  &loout,
// 				Unit:      &unitA,
// 			},
// 			{
// 				Context:   &ctxSamplePeriodic,
// 				Measurand: &irmea,
// 				Value:     "11.00",
// 				Location:  &loout,
// 				Unit:      &unitWh,
// 			},
// 			{
// 				Context:   &ctxSamplePeriodic,
// 				Measurand: &vomea,
// 				Value:     "11.00",
// 				Location:  &loout,
// 				Unit:      &unitV,
// 				Phase:     &pl1,
// 			},
// 			{
// 				Context:   &ctxSamplePeriodic,
// 				Measurand: &vomea,
// 				Value:     "11.00",
// 				Location:  &loout,
// 				Unit:      &unitV,
// 				Phase:     &pl2,
// 			},
// 			{
// 				Context:   &ctxSamplePeriodic,
// 				Measurand: &vomea,
// 				Value:     "11.00",
// 				Location:  &loout,
// 				Unit:      &unitV,
// 				Phase:     &pl3,
// 			},
// 			{
// 				Context:   &ctxSamplePeriodic,
// 				Measurand: &pwmea,
// 				Value:     "11.00",
// 				Location:  &loout,
// 				Unit:      &unitW,
// 			},
// 			{
// 				Context:   &ctxSamplePeriodic,
// 				Measurand: &socmea,
// 				Value:     "11",
// 				Location:  &loev,
// 				Unit:      &unitPercent,
// 			},
// 			{
// 				Context:   &ctxSamplePeriodic,
// 				Measurand: &tempmea,
// 				Value:     "11",
// 				Location:  &loev,
// 				Unit:      &unitCelcius,
// 			},
// 		},
// 	}
// 	meterReq = newTestEquipMeterValuesOCPP16Request(cid, meterValue)
// 	err = MeterValuesRequest(meterReq)
// 	assert.Nil(t, err)

// 	time.Sleep(2 * time.Second)

// 	meterValue = &equipMeterValuesRequestMeterValue{
// 		SampledValue: []MeterValueElemSampledValueElem{
// 			{
// 				Context:   &ctxSamplePeriodic,
// 				Phase:     &pl1,
// 				Measurand: &curmea,
// 				Value:     "11.00",
// 				Location:  &loout,
// 				Unit:      &unitA,
// 			},
// 			{
// 				Context:   &ctxSamplePeriodic,
// 				Phase:     &pl2,
// 				Measurand: &curmea,
// 				Value:     "11.00",
// 				Location:  &loout,
// 				Unit:      &unitA,
// 			},
// 			{
// 				Context:   &ctxSamplePeriodic,
// 				Phase:     &pl3,
// 				Measurand: &curmea,
// 				Value:     "11.00",
// 				Location:  &loout,
// 				Unit:      &unitA,
// 			},
// 			{
// 				Context:   &ctxSamplePeriodic,
// 				Measurand: &irmea,
// 				Value:     "11.00",
// 				Location:  &loout,
// 				Unit:      &unitWh,
// 			},
// 			{
// 				Context:   &ctxSamplePeriodic,
// 				Measurand: &vomea,
// 				Value:     "11.00",
// 				Location:  &loout,
// 				Unit:      &unitV,
// 				Phase:     &pl1,
// 			},
// 			{
// 				Context:   &ctxSamplePeriodic,
// 				Measurand: &vomea,
// 				Value:     "11.00",
// 				Location:  &loout,
// 				Unit:      &unitV,
// 				Phase:     &pl2,
// 			},
// 			{
// 				Context:   &ctxSamplePeriodic,
// 				Measurand: &vomea,
// 				Value:     "11.00",
// 				Location:  &loout,
// 				Unit:      &unitV,
// 				Phase:     &pl3,
// 			},
// 			{
// 				Context:   &ctxSamplePeriodic,
// 				Measurand: &pwmea,
// 				Value:     "11.00",
// 				Location:  &loout,
// 				Unit:      &unitW,
// 			},
// 			{
// 				Context:   &ctxSamplePeriodic,
// 				Measurand: &socmea,
// 				Value:     "11",
// 				Location:  &loev,
// 				Unit:      &unitPercent,
// 			},
// 			{
// 				Context:   &ctxSamplePeriodic,
// 				Measurand: &tempmea,
// 				Value:     "11",
// 				Location:  &loev,
// 				Unit:      &unitCelcius,
// 			},
// 		},
// 	}
// 	meterReq = newTestEquipMeterValuesOCPP16Request(cid, meterValue)
// 	err = MeterValuesRequest(meterReq)
// 	assert.Nil(t, err)

// 	time.Sleep(2 * time.Second)

// 	meterValue = &equipMeterValuesRequestMeterValue{
// 		SampledValue: []MeterValueElemSampledValueElem{
// 			{
// 				Context:   &ctxEnd,
// 				Phase:     &pl1,
// 				Measurand: &curmea,
// 				Value:     "11.00",
// 				Location:  &loout,
// 				Unit:      &unitA,
// 			},
// 			{
// 				Context:   &ctxEnd,
// 				Phase:     &pl2,
// 				Measurand: &curmea,
// 				Value:     "11.00",
// 				Location:  &loout,
// 				Unit:      &unitA,
// 			},
// 			{
// 				Context:   &ctxEnd,
// 				Phase:     &pl3,
// 				Measurand: &curmea,
// 				Value:     "11.00",
// 				Location:  &loout,
// 				Unit:      &unitA,
// 			},
// 			{
// 				Context:   &ctxEnd,
// 				Measurand: &irmea,
// 				Value:     "11.00",
// 				Location:  &loout,
// 				Unit:      &unitWh,
// 			},
// 			{
// 				Context:   &ctxEnd,
// 				Measurand: &vomea,
// 				Value:     "11.00",
// 				Location:  &loout,
// 				Unit:      &unitV,
// 				Phase:     &pl1,
// 			},
// 			{
// 				Context:   &ctxEnd,
// 				Measurand: &vomea,
// 				Value:     "11.00",
// 				Location:  &loout,
// 				Unit:      &unitV,
// 				Phase:     &pl2,
// 			},
// 			{
// 				Context:   &ctxEnd,
// 				Measurand: &vomea,
// 				Value:     "11.00",
// 				Location:  &loout,
// 				Unit:      &unitV,
// 				Phase:     &pl3,
// 			},
// 			{
// 				Context:   &ctxEnd,
// 				Measurand: &pwmea,
// 				Value:     "11.00",
// 				Location:  &loout,
// 				Unit:      &unitW,
// 			},
// 			{
// 				Context:   &ctxEnd,
// 				Measurand: &socmea,
// 				Value:     "11",
// 				Location:  &loev,
// 				Unit:      &unitPercent,
// 			},
// 			{
// 				Context:   &ctxEnd,
// 				Measurand: &tempmea,
// 				Value:     "11",
// 				Location:  &loev,
// 				Unit:      &unitCelcius,
// 			},
// 		},
// 	}

// 	stopReq := newTestEquipStopTransactionOCPP16Request(1000, cid, startResp.Data.TransactionId, StopReasonTypeOther, meterValue)
// 	stopResp, err := StopTransactionRequestWithGeneric(ctx, stopReq)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, stopResp)

// 	reqOffline := newTestEquipOfflineRequest(services.TestSN, p, services.TestAccessPod, id.Next().String(), EOF)
// 	err = OfflineRequestWithGeneric(ctx, reqOffline)
// 	assert.Nil(t, err)
// }

// func newTestEquipStartTransactionOCPP16Request(meterStart int, idToken, cid string) *equipStartTransactionOCPP16Request {
// 	return NewEquipStartTransactionRequestOCPP16(services.TestSN, services.TestAccessPod, id.Next().String(), idToken, meterStart, cid, time.Now().Unix())
// }

// func newTestEquipMeterValuesOCPP16Request(cid string, meterValue *equipMeterValuesRequestMeterValue) *equipMeterValuesRequest {
// 	r := NewEquipMeterValuesOCPP16Request(services.TestSN, services.TestAccessPod, id.Next().String(), cid)
// 	r.Data.MeterValue = meterValue
// 	return r
// }

// func newTestEquipStopTransactionOCPP16Request(meterStop int, cid, transId string, reason StopReasonTypeMenu, meterValue *equipMeterValuesRequestMeterValue) *equipStopTransactionOCPP16Request {
// 	r := NewEquipStopTransactionOCPP16Request(services.TestSN, services.TestAccessPod, id.Next().String(), meterStop, reason, transId, false, time.Now().Unix())
// 	r.Data.MeterValue = meterValue
// 	return r
// }
