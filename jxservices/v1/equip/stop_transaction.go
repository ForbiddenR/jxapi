package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type StoppingReasonType int

const (
	StoppingReasonTypeNormal                                       StoppingReasonType = 0   // 按照服务类型（按电量充电，按时长充电，按金额充电,
	StoppingReasonTypeLocalCard                                    StoppingReasonType = 1   // 用户本地刷卡停止充电
	StoppingReasonTypeLocalPIN                                     StoppingReasonType = 2   // 用户本地输入校验码停止充电
	StoppingReasonTypeRemoteUser                                   StoppingReasonType = 3   // 用户远程结束
	StoppingReasonTypeRemoteAdmin                                  StoppingReasonType = 4   // 管理员远程结束
	StoppingReasonTypeEmergencyStop                                StoppingReasonType = 5   // 急停按下，停止充电
	StoppingReasonTypeEVDisconnected                               StoppingReasonType = 6   // 检测到枪头断开，停止充电
	StoppingReasonTypeReboot                                       StoppingReasonType = 7   // 系统重启停止充电
	StoppingReasonTypeOffLine                                      StoppingReasonType = 8   // 充电桩断线，停止充电（对于不允许离线充电的充电桩）
	StoppingReasonTypePowerLoss                                    StoppingReasonType = 9   // 充电桩掉电，停止充电
	StoppingReasonTypeSystemFault                                  StoppingReasonType = 10  // 充电桩故障，停止充电
	StoppingReasonTypeBMSFault                                     StoppingReasonType = 11  // 电动车故障，停止充电
	StoppingReasonTypeOther                                        StoppingReasonType = 12  // 其它原因，停止充电
	StoppingReasonTypeNotEnoughMoney                               StoppingReasonType = 17  // 余额不足
	StoppingReasonTypeOverLoad                                     StoppingReasonType = 18  // 过载停止
	StoppingReasonTypeOverVoltage                                  StoppingReasonType = 19  // 过压停止
	StoppingReasonTypeUnderVoltage                                 StoppingReasonType = 20  // 欠压停止
	StoppingReasonTypeNetTimeout                                   StoppingReasonType = 21  // 网络同步超时停止充电
	StoppingReasonTypeCPError                                      StoppingReasonType = 25  // CP错误
	StoppingReasonTypeLeakage                                      StoppingReasonType = 26  // 漏电故障
	StoppingReasonTypeSocFull                                      StoppingReasonType = 16  // soc充满停止
	StoppingReasonTypeOverTime                                     StoppingReasonType = 22  // 超过设置时间
	StoppingReasonTypeOverMeter                                    StoppingReasonType = 23  // 超过设置电量
	StoppingReasonTypeOverSOC                                      StoppingReasonType = 24  // 超过设置SOC
	StoppingReasonTypeLightningProtection                          StoppingReasonType = 27  // 避雷
	StoppingReasonTypeAmbientOvertemperature                       StoppingReasonType = 28  // 环境过温
	StoppingReasonTypeMeterCommunicationFault                      StoppingReasonType = 13  // 电表通讯故障
	StoppingReasonTypeAccessControlFault                           StoppingReasonType = 29  // 门禁故障
	StoppingReasonTypeAbnormalCommunicationOfCardReader            StoppingReasonType = 30  // 刷卡器通讯异常
	StoppingReasonTypeCC1StatusError                               StoppingReasonType = 31  // CC1状态错误
	StoppingReasonTypeOpenSolenoidLockError                        StoppingReasonType = 32  // 开启电磁锁错误
	StoppingReasonTypeAbnormalElectronicLock                       StoppingReasonType = 33  // 电子锁异常
	StoppingReasonTypeAdaptiveVoltageError                         StoppingReasonType = 34  // 适配电压错误
	StoppingReasonTypeChargingRelayOpeningError                    StoppingReasonType = 35  // 充电器开启错误
	StoppingReasonTypeChargingRelayClosingError                    StoppingReasonType = 36  // 充电器关闭错误
	StoppingReasonTypeInsulationDetectionError                     StoppingReasonType = 37  // 绝缘监测错误
	StoppingReasonTypePreChargeVoltageAndMessageFivePercentError   StoppingReasonType = 38  // 预充电池电压和报文差5%错误
	StoppingReasonTypePreChargeVoltageOverMaxVoltageError          StoppingReasonType = 39  // 预充电压大于允许最大电压
	StoppingReasonTypePreChargeVoltageUnderMinVoltageError         StoppingReasonType = 40  // 预充电压小于允许最小电压
	StoppingReasonTypePreChargeAdapterVoltageError                 StoppingReasonType = 41  // 预充适配电压错误
	StoppingReasonTypeStopCurrentOverFiveError                     StoppingReasonType = 42  // 停止充电电流大于5A
	StoppingReasonTypeConnectionRelayOpenError                     StoppingReasonType = 43  // 充电继电器开启错误
	StoppingReasonTypeConnectionRelayClosingError                  StoppingReasonType = 44  // 充电继电器关闭错误
	StoppingReasonTypeOutputOverVoltageError                       StoppingReasonType = 45  // 输出过压错误
	StoppingReasonTypeOutputOverCurrentError                       StoppingReasonType = 46  // 输出过压错误
	StoppingReasonTypeConnectorOvertemperature                     StoppingReasonType = 47  // 枪过温
	StoppingReasonTypeChargingStationOtherFault                    StoppingReasonType = 48  // 充电桩其他故障
	StoppingReasonTypeBMSNotReady                                  StoppingReasonType = 49  // BMS未就绪
	StoppingReasonTypeBHMTimeoutError                              StoppingReasonType = 50  // BHM超时错误
	StoppingReasonTypeBRMTimeoutError                              StoppingReasonType = 51  // BRM超时错误
	StoppingReasonTypeBCPTimeoutError                              StoppingReasonType = 52  // BCP超时错误
	StoppingReasonTypeBROTimeoutError                              StoppingReasonType = 53  // BRO超时错误
	StoppingReasonTypeBROAATimeoutError                            StoppingReasonType = 54  // BROAA超时错误
	StoppingReasonTypeBROSeriousFault                              StoppingReasonType = 55  // BROAA重大错误停止充电
	StoppingReasonTypeBCLTimeoutError                              StoppingReasonType = 56  // BCL超时错误
	StoppingReasonTypeBCSTimeoutError                              StoppingReasonType = 57  // BCS超时错误
	StoppingReasonTypeBSMTimeoutError                              StoppingReasonType = 58  // BSM超时错误
	StoppingReasonTypeBSTTimeoutError                              StoppingReasonType = 59  // BST超时错误
	StoppingReasonTypeBSDTimeoutError                              StoppingReasonType = 60  // BSD超时错误
	StoppingReasonTypeBEMTimeoutError                              StoppingReasonType = 61  // BEM充电错误报文超时
	StoppingReasonTypeBSTInsulationFault                           StoppingReasonType = 62  // BST绝缘故障
	StoppingReasonTypeBSTComponentOvertemperature                  StoppingReasonType = 63  // BST元件过温
	StoppingReasonTypeBSTConnectorFault                            StoppingReasonType = 64  // BST连接器故障
	StoppingReasonTypeBSTBatteryGroupOvertemperature               StoppingReasonType = 65  // BST电池组过温
	StoppingReasonTypeBSTHighVoltageRelayFault                     StoppingReasonType = 66  // BST高压继电器故障
	StoppingReasonTypeBSTCheckpointTwoVoltageFault                 StoppingReasonType = 67  // BST检测点2电压故障
	StoppingReasonTypeBSTOvercurrent                               StoppingReasonType = 68  // BST电流过大
	StoppingReasonTypeBSTOvervoltage                               StoppingReasonType = 69  // BST电压过大
	StoppingReasonTypeSinglePowerBatteryOvervoltage                StoppingReasonType = 70  // 单体动力蓄电池电压过高
	StoppingReasonTypeSinglePowerBatteryUndervoltage               StoppingReasonType = 71  // 单体动力蓄电池电压过低
	StoppingReasonTypeEntireVehiclePowerBatteryChargeStateOverSoC  StoppingReasonType = 72  // 整车动力蓄电池荷电状态SoC过高
	StoppingReasonTypeEntireVehiclePowerBatteryChargeStateUnderSoC StoppingReasonType = 73  // 整车动力蓄电池喝点状态SoC过低
	StoppingReasonTypeMaximumBatteryPackOvertemperature            StoppingReasonType = 74  // 最高电池组过温
	StoppingReasonTypePowerBatteryChargingOvercurrent              StoppingReasonType = 75  // 动力蓄电池充电过流
	StoppingReasonTypePowerBatteryChargingOvertemperature          StoppingReasonType = 76  // 动力蓄电池温度过高
	StoppingReasonTypePowerBatteryIsulationFault                   StoppingReasonType = 77  // 动力蓄电池绝缘故障
	StoppingReasonTypePowerBatteryConnectorFault                   StoppingReasonType = 78  // 动力蓄电池连接器故障
	StoppingReasonTypeChargeNotAllowedStatus                       StoppingReasonType = 79  // 充电不允许的状态
	StoppingReasonTypeBMSOhterFault                                StoppingReasonType = 80  // BMS其他故障
	StoppingReasonTypePriceSchemeException                         StoppingReasonType = 81  // 计费模板异常
	StoppingReasonTypeFullOfSelfStop                               StoppingReasonType = 101 // 充满自停(原22)
	StoppingReasonTypeAdminForceStop                               StoppingReasonType = 102 // 管理员强制结束(原23)
	StoppingReasonTypeEVOccupy                                     StoppingReasonType = 103 // 接口已被占用(原24)
	StoppingReasonTypeNoTariff                                     StoppingReasonType = 104 // 没有计费模版(原25)
	StoppingReasonTypeConnectorFault                               StoppingReasonType = 105 // 充电枪故障(原26)
	StoppingReasonTypeCC1NoConnect                                 StoppingReasonType = 106 // 充电枪未连接(充电中车辆控制引导异常/CC1连接异常)(原27)
	StoppingReasonTypeDeauthorized                                 StoppingReasonType = 107 // 取消授权
	StoppingReasonTypeEVConnectTimeout                             StoppingReasonType = 108 // 连接超时
	StoppingReasonTypeServer_Error                                 StoppingReasonType = 999 // 平台错误
	StoppingReasonTypeHardReset                                    StoppingReasonType = 82
	StoppingReasonTypeSoftReset                                    StoppingReasonType = 83
	StoppingReasonTypeUnlockCommand                                StoppingReasonType = 84
	StoppingReasonTypeEnergyLimitReached                           StoppingReasonType = 85
	StoppingReasonTypeGroundFault                                  StoppingReasonType = 86
	StoppingReasonTypeImmediateReset                               StoppingReasonType = 87
	StoppingReasonTypeLocal                                        StoppingReasonType = 88
	StoppingReasonTypeLocalOutOfCredit                             StoppingReasonType = 89
	StoppingReasonTypeMasterPass                                   StoppingReasonType = 90
	StoppingReasonTypeOvercurrentFault                             StoppingReasonType = 91
	StoppingReasonTypePowerQuality                                 StoppingReasonType = 92
	StoppingReasonTypeSocLimitReached                              StoppingReasonType = 93
	StoppingReasonTypeStoppedByEV                                  StoppingReasonType = 94
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
	IdTokenType     *IdTokenType       `json:"idTokenType,omitempty"`
	MeterStop       *int               `json:"meterStop"`
	EvseSerial      *string            `json:"evseSerial,omitempty"`
	ConnectorSerial *string            `json:"connectorSerial,omitempty"`
	ReservationId   *int64             `json:"reservationId,omitempty"`
	TransactionId   string             `json:"transactionId"`
	RemoteStartId   *int64             `json:"remoteStartId,omitempty"`
	Offline         bool               `json:"offline"`
	Timestamp       int64              `json:"timestamp"`
	MeterValue      *MeterValue        `json:"meterValue,omitempty"`
	Tariff          *Tariff            `json:"tariff,omitempty"`
	ChargingState   uint8              `json:"chargingState"`
	Vin             *string            `json:"vin,omitempty"`
	StopReason      StoppingReasonType `json:"stopReason"`
}

func (*equipStopTransactionRequest) GetName() string {
	return services.StopTransaction.String()
}

func NewEquipStopTransactionRequest(sn, pod, msgID string, p *services.Protocol,
	reason StoppingReasonType, transactionId string, isOffline bool, timestamp int64) *equipStopTransactionRequest {
	req := &equipStopTransactionRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.StopTransaction.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipStopTransactionRequestDetail{
			IdTokenType:   &IdTokenType{},
			StopReason:    reason,
			TransactionId: transactionId,
			Timestamp:     timestamp,
			Offline:       isOffline,
			MeterValue:    &MeterValue{},
		},
	}
	if !p.Equal(services.OCPP16()) {
		req.Data.Tariff = &Tariff{
			Id: -1,
		}
	}
	return req
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
