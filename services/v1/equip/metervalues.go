package equip

import (
	"context"

	// "gitee.com/csms/jxeu-ocpp/internal/rabbitmq"
	"github.com/Kotodian/gokit/datasource/rabbitmq"
	"github.com/ForbiddenR/jx-api/services"
)

// func OCPP16SampledValueToEquipSampledValue(s ocpp16.MeterValueElemSampledValueElem) MeterValueElemSampledValueElem {
// 	e := MeterValueElemSampledValueElem{}
// 	if s.Context != nil {
// 		ctx := MeterValueElemSampledValueElemContext(*s.Context)
// 		e.Context = &ctx
// 	}
// 	if s.Format != nil {
// 		f := MeterValueElemSampledValueElemFormat(*s.Format)
// 		e.Format = &f
// 	}
// 	if s.Location != nil {
// 		l := MeterValueElemSampledValueElemLocation(*s.Location)
// 		e.Location = &l
// 	}
// 	if s.Measurand != nil {
// 		m := MeterValueElemSampledValueElemMeasurand(*s.Measurand)
// 		e.Measurand = &m
// 	}
// 	if s.Phase != nil {
// 		p := MeterValueElemSampledValueElemPhase(*s.Phase)
// 		e.Phase = &p
// 	}
// 	if s.Unit != nil {
// 		u := MeterValueElemSampledValueElemUnit(*s.Unit)
// 		e.Unit = &u
// 	}
// 	e.Value = s.Value
// 	return e

// }

type MeterValueElemSampledValueElem struct {
	// Context corresponds to the JSON schema field "context".
	Context *MeterValueElemSampledValueElemContext `json:"context,omitempty" yaml:"context,omitempty"`

	// Format corresponds to the JSON schema field "format".
	Format *MeterValueElemSampledValueElemFormat `json:"format,omitempty" yaml:"format,omitempty"`

	// Location corresponds to the JSON schema field "location".
	Location *MeterValueElemSampledValueElemLocation `json:"location,omitempty" yaml:"location,omitempty"`

	// Measurand corresponds to the JSON schema field "measurand".
	Measurand *MeterValueElemSampledValueElemMeasurand `json:"measurand,omitempty" yaml:"measurand,omitempty"`

	// Phase corresponds to the JSON schema field "phase".
	Phase *MeterValueElemSampledValueElemPhase `json:"phase,omitempty" yaml:"phase,omitempty"`

	// Unit corresponds to the JSON schema field "unit".
	Unit *MeterValueElemSampledValueElemUnit `json:"unit,omitempty" yaml:"unit,omitempty"`

	// Value corresponds to the JSON schema field "value".
	Value string `json:"value" yaml:"value"`
}

type MeterValueElemSampledValueElemContext string

const MeterValueElemSampledValueElemContextInterruptionBegin MeterValueElemSampledValueElemContext = "Interruption.Begin"
const MeterValueElemSampledValueElemContextInterruptionEnd MeterValueElemSampledValueElemContext = "Interruption.End"
const MeterValueElemSampledValueElemContextOther MeterValueElemSampledValueElemContext = "Other"
const MeterValueElemSampledValueElemContextSampleClock MeterValueElemSampledValueElemContext = "Sample.Clock"
const MeterValueElemSampledValueElemContextSamplePeriodic MeterValueElemSampledValueElemContext = "Sample.Periodic"
const MeterValueElemSampledValueElemContextTransactionBegin MeterValueElemSampledValueElemContext = "Transaction.Begin"
const MeterValueElemSampledValueElemContextTransactionEnd MeterValueElemSampledValueElemContext = "Transaction.End"
const MeterValueElemSampledValueElemContextTrigger MeterValueElemSampledValueElemContext = "Trigger"

type MeterValueElemSampledValueElemFormat string

const MeterValueElemSampledValueElemFormatRaw MeterValueElemSampledValueElemFormat = "Raw"
const MeterValueElemSampledValueElemFormatSignedData MeterValueElemSampledValueElemFormat = "SignedData"

type MeterValueElemSampledValueElemLocation string

const MeterValueElemSampledValueElemLocationBody MeterValueElemSampledValueElemLocation = "Body"
const MeterValueElemSampledValueElemLocationCable MeterValueElemSampledValueElemLocation = "Cable"
const MeterValueElemSampledValueElemLocationEV MeterValueElemSampledValueElemLocation = "EV"
const MeterValueElemSampledValueElemLocationInlet MeterValueElemSampledValueElemLocation = "Inlet"
const MeterValueElemSampledValueElemLocationOutlet MeterValueElemSampledValueElemLocation = "Outlet"

type MeterValueElemSampledValueElemMeasurand string

const MeterValueElemSampledValueElemMeasurandCurrentExport MeterValueElemSampledValueElemMeasurand = "Current.Export"
const MeterValueElemSampledValueElemMeasurandCurrentImport MeterValueElemSampledValueElemMeasurand = "Current.Import"
const MeterValueElemSampledValueElemMeasurandCurrentOffered MeterValueElemSampledValueElemMeasurand = "Current.Offered"
const MeterValueElemSampledValueElemMeasurandEnergyActiveExportInterval MeterValueElemSampledValueElemMeasurand = "Energy.Active.Export.Interval"
const MeterValueElemSampledValueElemMeasurandEnergyActiveExportRegister MeterValueElemSampledValueElemMeasurand = "Energy.Active.Export.Register"
const MeterValueElemSampledValueElemMeasurandEnergyActiveImportRegister MeterValueElemSampledValueElemMeasurand = "Energy.Active.Import.Register"
const MeterValueElemSampledValueElemMeasurandEnergyReactiveExportInterval MeterValueElemSampledValueElemMeasurand = "Energy.Reactive.Export.Interval"
const MeterValueElemSampledValueElemMeasurandEnergyActiveImportInterval MeterValueElemSampledValueElemMeasurand = "Energy.Active.Import.Interval"
const MeterValueElemSampledValueElemMeasurandEnergyReactiveExportRegister MeterValueElemSampledValueElemMeasurand = "Energy.Reactive.Export.Register"
const MeterValueElemSampledValueElemMeasurandEnergyReactiveImportInterval MeterValueElemSampledValueElemMeasurand = "Energy.Reactive.Import.Interval"
const MeterValueElemSampledValueElemMeasurandEnergyReactiveImportRegister MeterValueElemSampledValueElemMeasurand = "Energy.Reactive.Import.Register"
const MeterValueElemSampledValueElemMeasurandFrequency MeterValueElemSampledValueElemMeasurand = "Frequency"
const MeterValueElemSampledValueElemMeasurandPowerActiveExport MeterValueElemSampledValueElemMeasurand = "Power.Active.Export"
const MeterValueElemSampledValueElemMeasurandPowerActiveImport MeterValueElemSampledValueElemMeasurand = "Power.Active.Import"
const MeterValueElemSampledValueElemMeasurandPowerFactor MeterValueElemSampledValueElemMeasurand = "Power.Factor"
const MeterValueElemSampledValueElemMeasurandPowerOffered MeterValueElemSampledValueElemMeasurand = "Power.Offered"
const MeterValueElemSampledValueElemMeasurandPowerReactiveExport MeterValueElemSampledValueElemMeasurand = "Power.Reactive.Export"
const MeterValueElemSampledValueElemMeasurandPowerReactiveImport MeterValueElemSampledValueElemMeasurand = "Power.Reactive.Import"
const MeterValueElemSampledValueElemMeasurandRPM MeterValueElemSampledValueElemMeasurand = "RPM"
const MeterValueElemSampledValueElemMeasurandSoC MeterValueElemSampledValueElemMeasurand = "SoC"
const MeterValueElemSampledValueElemMeasurandTemperature MeterValueElemSampledValueElemMeasurand = "Temperature"
const MeterValueElemSampledValueElemMeasurandVoltage MeterValueElemSampledValueElemMeasurand = "Voltage"

type MeterValueElemSampledValueElemPhase string

const MeterValuesJsonMeterValueElemSampledValueElemPhaseL1 MeterValueElemSampledValueElemPhase = "L1"
const MeterValuesJsonMeterValueElemSampledValueElemPhaseL1L2 MeterValueElemSampledValueElemPhase = "L1-L2"
const MeterValuesJsonMeterValueElemSampledValueElemPhaseL1N MeterValueElemSampledValueElemPhase = "L1-N"
const MeterValuesJsonMeterValueElemSampledValueElemPhaseL2 MeterValueElemSampledValueElemPhase = "L2"
const MeterValuesJsonMeterValueElemSampledValueElemPhaseL2L3 MeterValueElemSampledValueElemPhase = "L2-L3"
const MeterValuesJsonMeterValueElemSampledValueElemPhaseL2N MeterValueElemSampledValueElemPhase = "L2-N"
const MeterValuesJsonMeterValueElemSampledValueElemPhaseL3 MeterValueElemSampledValueElemPhase = "L3"
const MeterValuesJsonMeterValueElemSampledValueElemPhaseL3L1 MeterValueElemSampledValueElemPhase = "L3-L1"
const MeterValuesJsonMeterValueElemSampledValueElemPhaseL3N MeterValueElemSampledValueElemPhase = "L3-N"
const MeterValuesJsonMeterValueElemSampledValueElemPhaseN MeterValueElemSampledValueElemPhase = "N"

type MeterValueElemSampledValueElemUnit string

const MeterValueElemSampledValueElemUnitA MeterValueElemSampledValueElemUnit = "A"
const MeterValueElemSampledValueElemUnitCelcius MeterValueElemSampledValueElemUnit = "Celcius"
const MeterValueElemSampledValueElemUnitCelsius MeterValueElemSampledValueElemUnit = "Celsius"
const MeterValueElemSampledValueElemUnitFahrenheit MeterValueElemSampledValueElemUnit = "Fahrenheit"
const MeterValueElemSampledValueElemUnitK MeterValueElemSampledValueElemUnit = "K"
const MeterValueElemSampledValueElemUnitKVA MeterValueElemSampledValueElemUnit = "kVA"
const MeterValueElemSampledValueElemUnitKW MeterValueElemSampledValueElemUnit = "kW"
const MeterValueElemSampledValueElemUnitKWh MeterValueElemSampledValueElemUnit = "kWh"
const MeterValueElemSampledValueElemUnitKvar MeterValueElemSampledValueElemUnit = "kvar"
const MeterValueElemSampledValueElemUnitKvarh MeterValueElemSampledValueElemUnit = "kvarh"
const MeterValueElemSampledValueElemUnitPercent MeterValueElemSampledValueElemUnit = "Percent"
const MeterValueElemSampledValueElemUnitV MeterValueElemSampledValueElemUnit = "V"
const MeterValueElemSampledValueElemUnitVA MeterValueElemSampledValueElemUnit = "VA"
const MeterValueElemSampledValueElemUnitVar MeterValueElemSampledValueElemUnit = "var"
const MeterValueElemSampledValueElemUnitVarh MeterValueElemSampledValueElemUnit = "varh"
const MeterValueElemSampledValueElemUnitW MeterValueElemSampledValueElemUnit = "W"
const MeterValueElemSampledValueElemUnitWh MeterValueElemSampledValueElemUnit = "Wh"

// This feature is transferred by rabbitmq.

const meterQueue = services.QueuePrefix + "metervalues"

type equipMeterValuesRequest struct {
	services.Base
	Data *equipMeterValuesRequestDetail `json:"data"`
}

type equipMeterValuesRequestDetail struct {
	EvseId        *string                            `json:"evseSerial,omitempty"`
	TransactionId *int64                             `json:"transactionId,omitempty"`
	ConnectorId   *string                            `json:"connectorSerial,omitempty"`
	MeterValue    *equipMeterValuesRequestMeterValue `json:"meterValue"`
}

type equipMeterValuesRequestMeterValue struct {
	Timestamp    int64                            `json:"timestamp"`
	SampledValue []MeterValueElemSampledValueElem `json:"sampledValue"`
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
			MeterValue:  &equipMeterValuesRequestMeterValue{},
		},
	}
	meterValue.Data.MeterValue.SampledValue = make([]MeterValueElemSampledValueElem, 0)
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
			MeterValue: &equipMeterValuesRequestMeterValue{
				Timestamp: timestamp,
			},
		},
	}
	meterValue.Data.MeterValue.SampledValue = make([]MeterValueElemSampledValueElem, 0)
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