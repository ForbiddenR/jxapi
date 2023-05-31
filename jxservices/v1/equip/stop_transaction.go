package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

// This feature only supports for the version "OCPP1.6".

type StoppedReasonType byte

const (
	StoppedReasonTypeNormal                                       StoppedReasonType = 0  // 正常停止充电
	StoppedReasonTypeIDCardStop                                   StoppedReasonType = 1  // 刷卡停止充电
	StoppedReasonTypeRemote                                       StoppedReasonType = 2  // 远程停止
	StoppedReasonTypeEmergencyStop                                StoppedReasonType = 3  // 急停
	StoppedReasonTypePowerLoss                                    StoppedReasonType = 4  // 掉电
	StoppedReasonTypeEVConnectTimeout                             StoppedReasonType = 5  // 与车端连接超时
	StoppedReasonTypeOther                                        StoppedReasonType = 6  // 其他
	StoppedReasonTypeEVDisconnected                               StoppedReasonType = 7  // 与车端断开
	StoppedReasonTypeSocFull                                      StoppedReasonType = 8  // soc充满停止
	StoppedReasonTypeOverMeter                                    StoppedReasonType = 9  // 达到智慧充电设置电量
	StoppedReasonTypeOverSOC                                      StoppedReasonType = 10 // 达到智慧充电设置SoC
	StoppedReasonTypeOverTime                                     StoppedReasonType = 11 // 达到智慧充电设置时间
	StoppedReasonTypeLightningProtection                          StoppedReasonType = 12 // 避雷
	StoppedReasonTypeAmbientOvertemperature                       StoppedReasonType = 13 // 环境过温
	StoppedReasonTypeMeterCommunicationFault                      StoppedReasonType = 14 // 电表通讯故障
	StoppedReasonTypeAccessControlFault                           StoppedReasonType = 15 // 门禁故障
	StoppedReasonTypeAbnormalCommunicationOfCardReader            StoppedReasonType = 16 // 刷卡器通讯异常
	StoppedReasonTypeCC1StatusError                               StoppedReasonType = 17 // CC1状态错误
	StoppedReasonTypeOpenSolenoidLockError                        StoppedReasonType = 18 // 开启电磁锁错误
	StoppedReasonTypeAbnormalElectronicLock                       StoppedReasonType = 19 // 电子锁异常
	StoppedReasonTypeAdaptiveVoltageError                         StoppedReasonType = 20 // 适配电压错误
	StoppedReasonTypeChargingRelayOpeningError                    StoppedReasonType = 21 // 充电器开启错误
	StoppedReasonTypeChargingRelayClosingError                    StoppedReasonType = 22 // 充电器关闭错误
	StoppedReasonTypeInsulationDetectionError                     StoppedReasonType = 23 // 绝缘监测错误
	StoppedReasonTypePreChargeVoltageAndMessageFivePercentError   StoppedReasonType = 24 // 预充电池电压和报文差5%错误
	StoppedReasonTypePreChargeVoltageOverMaxVoltageError          StoppedReasonType = 25 // 预充电压大于允许最大电压
	StoppedReasonTypePreChargeVoltageUnderMinVoltageError         StoppedReasonType = 26 // 预充电压小于允许最小电压
	StoppedReasonTypePreChargeAdapterVoltageError                 StoppedReasonType = 27 // 预充适配电压错误
	StoppedReasonTypeStopCurrentOverFiveError                     StoppedReasonType = 28 // 停止充电电流大于5A
	StoppedReasonTypeConnectionRelayOpenError                     StoppedReasonType = 29 // 充电继电器开启错误
	StoppedReasonTypeConnectionRelayClosingError                  StoppedReasonType = 30 // 充电继电器关闭错误
	StoppedReasonTypeOutputOverVoltageError                       StoppedReasonType = 31 // 输出过压错误
	StoppedReasonTypeOutputOverCurrentError                       StoppedReasonType = 32 // 输出过压错误
	StoppedReasonTypeConnectorOvertemperature                     StoppedReasonType = 33 // 枪过温
	StoppedReasonTypeChargingStationOtherFault                    StoppedReasonType = 34 // 充电桩其他故障
	StoppedReasonTypeBMSNotReady                                  StoppedReasonType = 35 // BMS未就绪
	StoppedReasonTypeBHMTimeoutError                              StoppedReasonType = 36 // BHM超时错误
	StoppedReasonTypeBRMTimeoutError                              StoppedReasonType = 37 // BRM超时错误
	StoppedReasonTypeBCPTimeoutError                              StoppedReasonType = 38 // BCP超时错误
	StoppedReasonTypeBROTimeoutError                              StoppedReasonType = 39 // BRO超时错误
	StoppedReasonTypeBROAATimeoutError                            StoppedReasonType = 40 // BROAA超时错误
	StoppedReasonTypeBROSeriousFault                              StoppedReasonType = 41 // BROAA重大错误停止充电
	StoppedReasonTypeBCLTimeoutError                              StoppedReasonType = 42 // BCL超时错误
	StoppedReasonTypeBCSTimeoutError                              StoppedReasonType = 43 // BCS超时错误
	StoppedReasonTypeBSMTimeoutError                              StoppedReasonType = 44 // BSM超时错误
	StoppedReasonTypeBSTTimeoutError                              StoppedReasonType = 45 // BST超时错误
	StoppedReasonTypeBSDTimeoutError                              StoppedReasonType = 46 // BSD超时错误
	StoppedReasonTypeBEMTimeoutError                              StoppedReasonType = 47 // BEM充电错误报文超时
	StoppedReasonTypeBSTInsulationFault                           StoppedReasonType = 48 // BST绝缘故障
	StoppedReasonTypeBSTComponentOvertemperature                  StoppedReasonType = 49 // BST元件过温
	StoppedReasonTypeBSTConnectorFault                            StoppedReasonType = 50 // BST连接器故障
	StoppedReasonTypeBSTBatteryGroupOvertemperature               StoppedReasonType = 51 // BST电池组过温
	StoppedReasonTypeBSTHighVoltageRelayFault                     StoppedReasonType = 52 // BST高压继电器故障
	StoppedReasonTypeBSTCheckpointTwoVoltageFault                 StoppedReasonType = 53 // BST检测点2电压故障
	StoppedReasonTypeBSTOvercurrent                               StoppedReasonType = 54 // BST电流过大
	StoppedReasonTypeBSTOvervoltage                               StoppedReasonType = 55 // BST电压过大
	StoppedReasonTypeSinglePowerBatteryOvervoltage                StoppedReasonType = 56 // 单体动力蓄电池电压过高
	StoppedReasonTypeSinglePowerBatteryUndervoltage               StoppedReasonType = 57 // 单体动力蓄电池电压过低
	StoppedReasonTypeEntireVehiclePowerBatteryChargeStateOverSoC  StoppedReasonType = 58 // 整车动力蓄电池荷电状态SoC过高
	StoppedReasonTypeEntireVehiclePowerBatteryChargeStateUnderSoC StoppedReasonType = 59 // 整车动力蓄电池喝点状态SoC过低
	StoppedReasonTypeMaximumBatteryPackOvertemperature            StoppedReasonType = 60 // 最高电池组过温
	StoppedReasonTypePowerBatteryChargingOvercurrent              StoppedReasonType = 61 // 动力蓄电池充电过流
	StoppedReasonTypePowerBatteryChargingOvertemperature          StoppedReasonType = 62 // 动力蓄电池温度过高
	StoppedReasonTypePowerBatteryIsulationFault                   StoppedReasonType = 63 // 动力蓄电池绝缘故障
	StoppedReasonTypePowerBatteryConnectorFault                   StoppedReasonType = 64 // 动力蓄电池连接器故障
	StoppedReasonTypeChargeNotAllowedStatus                       StoppedReasonType = 65 // 充电不允许的状态
	StoppedReasonTypeBMSOtherFault                                StoppedReasonType = 66 // BMS其他故障
	StoppedReasonTypePriceSchemeExepction                         StoppedReasonType = 67 // 计费模板异常
)

type equipStopTransactionRequest struct {
	services.Base
	Data *equipStopTransactionRequestDetail `json:"data"`
}

type equipStopTransactionRequestDetail struct {
	IdTokenType     *IdTokenType      `json:"idTokenType,omitempty"`
	MeterStop       *int              `json:"meterStop"`
	EvseSerial      *string           `json:"evseSerial,omitempty"`
	ConnectorSerial *string           `json:"connectorSerial,omitempty"`
	ReservationId   *int64            `json:"reservationId,omitempty"`
	TransactionId   string            `json:"transactionId"`
	RemoteStartId   *int64           `json:"remoteStartId,omitempty"`
	Offline         bool              `json:"offline"`
	Timestamp       int64             `json:"timestamp"`
	MeterValue      []MeterValue      `json:"meterValue,omitempty"`
	Tariff          *Tariff           `json:"tariff,omitempty"`
	ChargingState   uint8             `json:"chargingState"`
	Vin             *string           `json:"vin,omitempty"`
	StopReason      StoppedReasonType `json:"stopReason"`
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

func (*equipStopTransactionRequest) GetName() string {
	return services.StopTransaction.String()
}

func NewEquipStopTransactionRequest(sn, pod, msgID string, p *services.Protocol,
	reason StoppedReasonType, transactionId string, isOffline bool, timestamp int64) *equipStopTransactionRequest {
	request := &equipStopTransactionRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    services.OCPP16(),
			Category:    services.StopTransaction.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipStopTransactionRequestDetail{
			StopReason:    reason,
			TransactionId: transactionId,
			Timestamp:     timestamp,
			Offline:       isOffline,
			MeterValue:    make([]MeterValue, 0),
			Tariff:        &Tariff{},
		},
	}
	return request
}

var _ services.Response = &equipStopTransactionResponse{}

type equipStopTransactionResponse struct {
	api.Response
	Data *equipStopTransactionResponseDetail `json:"data"`
}

func (resp *equipStopTransactionResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipStopTransactionResponse) GetMsg() string {
	return resp.Msg
}

type equipStopTransactionResponseDetail struct {
	IdTokenInfo
}

func StopTransactionRequestWithGeneric(ctx context.Context, req *equipStopTransactionRequest) (*equipStopTransactionResponse, error) {
	header := services.GetSimpleHeaderValue(services.StopTransaction)

	url := services.GetSimpleURL(req)

	return services.RequestWithResponse(ctx, req, url, header, &equipStopTransactionResponse{})
}
