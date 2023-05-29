package jxservices

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
)

// These constants are usually used in services package many times.
const (
	Equip           = "ac"
	Equipment       = "equip"
	Callback        = "callback"
	QueuePrefix     = "mq_services_"
	TestConnectorId = "1"
	TestSN          = "JK000000006"
	TestAccessPod   = "jx-acos-0"
	CallbackSuffix  = "Callback"
	//Acos            = "Acos"
)

// Define some default status fields in callback.
const (
	CallbackError = -1
	Successful    = 0
	Failed        = 1
)

// These constants are associated with the topics of mqtt.
// todo Some customize feature related to DataTransfer and sending to charging point will be actually announced.
const (
	// Core
	DataTransferFeatureName           = "dataTransfer"
	ChangeAvailabilityFeatureName     = "changeAvailability"
	ChangeConfigurationFeatureName    = "setVariables"
	GetConfigurationFeatureName       = "getVariables"
	GetBaseReportFeatureName          = "getBaseReport"
	ClearCacheFeatureName             = "clearCache"
	RemoteStartTransactionFeatureName = "remoteStartTransaction"
	RemoteStopTransactionFeatureName  = "remoteStopTransaction"
	ResetFeatureName                  = "reset"
	UnlockConnectorFeatureName        = "unlockConnector"
	// Firmware
	GetDiagnosticsFeatureName = "getDiagnostics"
	UpdateFirmwareFeatureName = "pushFirmware"
	// LocalAuth
	GetLocalListVersionFeatureName = "getLocalListVersion"
	SendLocalListFeatureName       = "setLocalAuthorizeList"
	// RemoteTrigger
	TriggerMessageFeatureName         = "triggerMessage"
	CallStatusNotificationFeatureName = "callStatusNotification"
	// Reservation
	CancelReservationFeatureName = "cancelReservation"
	ReserveNowFeatureName        = "reserveNow"
	// SmartCharging
	ClearChargingProfileFeatureName = "clearChargingProfile"
	GetCompositeScheduleFeatureName = "detCompositeSchedule"
	SetChargingProfileFeatureName   = "setChargingProfile"
	// Customize Features
	SetChargingTimerFeatureName      = "setChargingTimer"
	SetLoadBalanceFeatureName        = "setLoadBalance"
	SetFactoryResetFeatureName       = "setFactoryReset"
	CloseFeatureName                 = "close"
	SetIntellectChargeFeatureName    = "setIntellectCharge"
	CancelIntellectChargeFeatureName = "cancelIntellectCharge"
	SetPriceSchemeFeatureName        = "setPriceScheme"
	// Customized Features about 104 protocol
	SendQRCodeFeatureName = "sendQRCode"
)

type Request2ServicesNameType string

const (
	Authorize                     Request2ServicesNameType = "authorize"
	BootNotification              Request2ServicesNameType = "bootNotification"
	ClearCache                    Request2ServicesNameType = "clear"
	DataTransfer                  Request2ServicesNameType = "dataTransfer"
	GetBaseReport                 Request2ServicesNameType = "getBaseReport"
	GetConfiguration              Request2ServicesNameType = "getVariables"
	MeterValues                   Request2ServicesNameType = "meterValues"
	UpdateTransaction             Request2ServicesNameType = "updateTransaction"
	Online                        Request2ServicesNameType = "equipOnline"
	Offline                       Request2ServicesNameType = "equipOffline"
	Register                      Request2ServicesNameType = "equipRegister"
	StatusNotification            Request2ServicesNameType = "statusNotification"
	StartTransaction              Request2ServicesNameType = "startTransaction"
	StopTransaction               Request2ServicesNameType = "stopTransaction"
	RemoteStartTransaction        Request2ServicesNameType = "remoteStartTransaction"
	RemoteStopTransaction         Request2ServicesNameType = "remoteStopTransaction"
	Reset                         Request2ServicesNameType = "reset"
	ReservationStatusNotification Request2ServicesNameType = "reservationStatusNotification"
	ChangeConfiguration           Request2ServicesNameType = "setVariables"
	SendLocalList                 Request2ServicesNameType = "setLocalAuthorizeList"
	SetChargingTimer              Request2ServicesNameType = "setChargingTimer"
	ChargingTimerNotification     Request2ServicesNameType = "timingStatusNotification"
	// ExpiredChargingTimer          Request2ServicesNameType = "expiredChargingTimer"
	UpdateFirmware                Request2ServicesNameType = "pushFirmware"
	FirmwareStatusNotification    Request2ServicesNameType = "firmwareStatusNotification"
	CallStatusNotification        Request2ServicesNameType = "callStatusNotification"
	GetDiagnostics                Request2ServicesNameType = "getDiagnostics"
	DiagnosticsStatusNotification Request2ServicesNameType = "diagnosticsStatusNotification"
	SetLoadBalance                Request2ServicesNameType = "setLoadBalance"
	SetFactoryReset               Request2ServicesNameType = "setFactoryReset"
	NotifyEvent                   Request2ServicesNameType = "notifyEvent"
	NotifyReport                  Request2ServicesNameType = "notifyReport"
	CancelReservation             Request2ServicesNameType = "cancelReservation"
	ReserveNow                    Request2ServicesNameType = "reserveNow"
	QRCode                        Request2ServicesNameType = "qrcode"
	SendQRCode                    Request2ServicesNameType = "sendQRCode"
	SetIntellectCharge            Request2ServicesNameType = "setIntellectCharge"
	CancelIntellectCharge         Request2ServicesNameType = "cancelIntellectCharge"
	SetPriceScheme                Request2ServicesNameType = "setPriceScheme"
	//TriggerMessage             Request2ServicesNameType = "callStatusNotification"
)

// FirstUpper is only for the interfaces having a regular category.
func (r Request2ServicesNameType) FirstUpper() string {
	s := r.String()
	return strings.ToUpper(s[:1]) + s[1:]
}

func (r Request2ServicesNameType) String() string {
	return string(r)
}

//type Request2ServicesPermsType string

// Split returns the value of "Perms".
func (r Request2ServicesNameType) Split() []string {
	switch r {
	case UpdateFirmware:
		return []string{"push", "firmware", "equipment"}
	case FirmwareStatusNotification:
		return []string{"push", "firmware", "notification"}
	case RemoteStartTransaction:
		return []string{"remote", "Start"}
	case RemoteStopTransaction:
		return []string{"remote", "stop"}
	case SendLocalList:
		return []string{"set", "local", "authorize"}
	case StartTransaction:
		return []string{"start", "transaction"}
	case StopTransaction:
		return []string{"stop", "transaction"}
	}
	for i := 0; i < len(r.String()); i++ {
		str := r.String()[i : i+1]
		if str == strings.ToUpper(str) {
			switch r.String()[:i] {
			case "equip":
				return []string{strings.ToLower(r.String()[i:])}
			case "authorize":
				return []string{r.String()[:i]}
			default:
				return r.SplitName()
			}
		}
	}
	return []string{r.String()}
}

// SplitName will be used by the function above to parser all the regular attributes.
func (r Request2ServicesNameType) SplitName() []string {
	var head, tail int
	var result []string
	for tail = 0; tail < len(r.String()); tail++ {
		str := r.String()[tail : tail+1]
		if str == strings.ToUpper(str) && tail != 0 {
			result = append(result, strings.ToLower(r.String()[head:tail]))
			head = tail
		}
	}
	if head < tail {
		result = append(result, strings.ToLower(r.String()[head:tail]))
	}
	return result
}

func (r Request2ServicesNameType) GetCallbackCategory() string {
	return r.FirstUpper() + CallbackSuffix
}

type RequestMeterValue struct {
	// Timestamp is the time when the Charging Station starts to get meter datas.
	Timestamp int64 `json:"timestamp"`
	// SampledValue uses the type, "MeterValueElemSampledValueElem", directly.
	SampledValue []MeterValueElemSampledValueElem `json:"sampledValue"`
}

type Base struct {
	EquipmentSn string    `json:"equipmentSn"`
	Protocol    *Protocol `json:"protocol"`
	Category    string    `json:"category"`
	AccessPod   string    `json:"accessPod"`
	MsgID       string    `json:"msgId"`
	// Callback    *CB       `json:"callback,omitempty"`
}

type Protocol struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func (p *Protocol) String() string {
	return p.Name + "" + p.Version
}

func (p *Protocol) Equal(p2 *Protocol) bool {
	return p.Name == p2.Name && p.Version == p2.Version
}

func (p *Protocol) UnmarshalJSON(data []byte) error {
	var v struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v.Name {
	case "OCPP":
		if v.Version != "1.6" && v.Version != "2.0.1" {
			return errors.New("invalid OCPP version: " + v.Version)
		}
	case "IEC104":
		if v.Version != "0.0.1" && v.Version != "0.0.2" {
			return errors.New("invalid IEC104 version: " + v.Version)
		}
	default:
		return errors.New("invalid protocol name: " + v.Name)
	}
	p.Name = v.Name
	p.Version = v.Version

	return nil
}

var ocpp16p = &Protocol{Name: "OCPP", Version: "1.6"}
var ocpp201p = &Protocol{Name: "OCPP", Version: "2.0.1"}
var iec001 = &Protocol{Name: "IEC104", Version: "0.0.1"}
var iec002 = &Protocol{Name: "IEC104", Version: "0.0.2"}

func OCPP16() *Protocol {
	return ocpp16p
}

func OCPP201() *Protocol {
	return ocpp201p
}

func IEC001() *Protocol {
	return iec001
}

func IEC002() *Protocol {
	return iec002
}

// CB includes callback information
type CB struct {
	Status int                              `json:"status"`
	Code   *apierrors.CallbackErrorCodeType `json:"code,omitempty"`
	Msg    *string                          `json:"msg,omitempty"`
}

func NewCB(status int) CB {
	return CB{Status: status}
}

func NewCBError(err *apierrors.CallbackError) CB {
	code := err.Code()
	msg := err.Error()
	return CB{Status: CallbackError, Code: &code, Msg: &msg}
}

// type IdTagInfoClass struct {
// 	ExpiryDate    *int64         `json:"expiryDate"`
// 	ParentIdToken *string        `json:"parentIdToken,omitempty"`
// 	Status        StatusTypeEnum `json:"status"`
// 	GroupIdToken  *GroupIdToken  `json:"groupIdToken,omitempty"`
// }

// type GroupIdToken struct {
// 	IdToken string `json:"idToken"`
// 	Type    string `json:"type"`
// }

// CallbackUniversalPortion is the common part of the callback request,
// including "status", "code", and "msg".
//type CallbackUniversalPortion struct {
//	Status int                           `json:"status"`
//	Code   *errors.CallbackErrorCodeType `json:"code"`
//	Msg    *string                       `json:"msg"`
//}

// Request interface needs to be implemented by all api.
type Request interface {
	GetName() string
}

// Response interface is implemented by all response struct
type Response interface {
	GetStatus() int
	GetMsg() string
}

// CallbackRequest interface needs to be implemented by all the callback api.
type CallbackRequest interface {
	Request
}

// boxing a callback request, the function can be invoked indirectly

// whichCallbackErr recognizes the type of the error, returning a corresponding callback error
func whichCallbackErr(clientId string, command string, err *apierrors.Error) *apierrors.CallbackError {
	switch err.Code {
	case apierrors.NotSupported:
		return apierrors.NewCallbackErrorSupported(clientId, command)
	case apierrors.NotImplemented:
		return apierrors.NewCallbackErrorNotImplemented(clientId, command)
	case apierrors.InternalError:
		return apierrors.NewCallbackErrorInternalError(clientId, command, err.Description)
	case apierrors.SecurityError:
		return apierrors.NewCallbackErrorSecurityError(clientId, command, err.Description)
	case apierrors.ProtocolError,
		apierrors.FormationViolation,
		apierrors.PropertyConstraintViolation,
		apierrors.TypeConstraintViolation:
		return apierrors.NewCallbackErrorPayloadError(clientId, command, err.Description)
	default:
		return apierrors.NewCallbackErrorGenericError(clientId, command, err.Description)
	}
}

// GetProperCallbackError turns an entering error into callback error
func GetProperCallbackError(clientId string, command string, err error) *apierrors.CallbackError {
	if cbErr, ok := err.(*apierrors.CallbackError); ok {
		return cbErr
	}

	if ocpError, ok := err.(*apierrors.Error); ok {
		cbErr := whichCallbackErr(clientId, command, ocpError)
		return cbErr
	}

	return apierrors.NewCallbackErrorOffline(clientId, command)
}

//func GetSimpleCategory(alias Request2ServicesNameType) string {
//	return alias.FirstUpper()
//}
//
//func GetCallbackCategory(alias Request2ServicesNameType) string {
//	return alias.FirstUpper() + CallbackSuffix
//}

func GetSimpleHeaderValue(alias Request2ServicesNameType) map[string]string {
	headerValue := make([]string, 0)
	headerValue = append(headerValue, api.Services, Equipment)
	headerValue = append(headerValue, alias.Split()...)
	header := map[string]string{api.Perms: strings.Join(headerValue, ":")}
	return header
}

func GetCallbackHeaderValue(alias Request2ServicesNameType) map[string]string {
	headerValue := make([]string, 0)
	headerValue = append(headerValue, api.Services)
	headerValue = append(headerValue, alias.Split()...)
	headerValue = append(headerValue, Callback)
	header := map[string]string{api.Perms: strings.Join(headerValue, ":")}
	return header
}

func GetSimpleURL(req Request) string {
	return api.ServicesUrl + Equip + "/" + req.GetName()
}

func GetCallbackURL(req Request) string {
	return api.ServicesUrl + Equip + "/" + Callback + "/" + req.GetName() + CallbackSuffix
}

func RequestWithoutResponse[T Response](ctx context.Context, req Request, url string, header map[string]string, t T) (err error) {
	message, err := api.SendRequest(ctx, url, req, header)
	if err != nil {
		return
	}

	err = json.Unmarshal(message, t)
	if err != nil {
		request, _ := json.Marshal(req)
		return apierrors.GetFailedResponseUnmarshalError(url, request, message, err)
	}

	// check whether it get an error from the services
	if t.GetStatus() == 1 {
		return errors.New(t.GetMsg())
	}
	return err
}

func RequestWithResponse[T Response](ctx context.Context, req Request, url string, header map[string]string, t T) (resp T, err error) {
	message, err := api.SendRequest(ctx, url, req, header)
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal(message, t)
	if err != nil {
		// The marshaling function has been verified before.
		request, _ := json.Marshal(req)
		return resp, apierrors.GetFailedResponseUnmarshalError(url, request, message, err)
	}
	resp = t
	// check whether it get an error from the services
	if resp.GetStatus() == 1 {
		return resp, errors.New(resp.GetMsg())
	}
	return resp, err
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

// fillErrorCallback boxes the fields of the structs implementing the concrete interface.
// func fillErrorCallback(req CallbackRequest, clientId, featureName string, err error) {
// 	callbackErr := GetProperCallbackError(clientId, featureName, err).(*errors.CallbackError)
// 	//code := string(callbackErr.Code())
// 	//msg := callbackErr.Error()
// 	//code, msg := processCallbackError(clientId, featureName, err)
// 	req.SetError(callbackErr)
// }

// transferFeatureName paseres the inputted string into standard form.
//func transferFeatureName(featureName string) string {
//	nameSlice := make([]string, 0)
//	var head, tail int
//	for tail = 0; tail < len(featureName); tail++ {
//		word := featureName[tail : tail+1]
//		if word != strings.ToUpper(word) || tail == 0 {
//			continue
//		}
//		nameSlice = append(nameSlice, strings.ToLower(featureName[head:tail]))
//		head = tail
//	}
//	if head < tail {
//		nameSlice = append(nameSlice, strings.ToLower(featureName[head:tail]))
//	}
//	return strings.Join(nameSlice, " ")
//}

// SendErrorCallbackRequest function is an ordinary method to send the assigned struct to the services.
// func SendErrorCallbackRequest(req CallbackRequest, clientId, featureName string, err error,
// 	logger *log.Logger, requestFunc func(CallbackRequest) error) {
// 	featureName = transferFeatureName(featureName)
// 	logger.Error(fmt.Sprintf("%s falied", featureName), zap.String("id", clientId), zap.Error(err))
// 	if req == nil || requestFunc == nil {
// 		logger.Error("some null input occurred in SendErrorCallbackRequest", zap.String("id", clientId))
// 		return
// 	}

// 	fillErrorCallback(req, clientId, featureName, err)

// 	err = requestFunc(req)
// 	if err != nil {
// 		logger.Error(fmt.Sprintf("%s fails to receive correct response of err callback request", featureName),
// 			zap.String("id", clientId), zap.Error(err))
// 	}
// }
