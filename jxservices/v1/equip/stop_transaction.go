package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type StoppingReasonType int

const (
	StoppingReasonTypeNormal                                       StoppingReasonType = 0  // 按照服务类型（按电量充电，按时长充电，按金额充电,
	StoppingReasonTypeLocalCard                                    StoppingReasonType = 1  // 用户本地刷卡停止充电
	StoppingReasonTypeLocalPIN                                     StoppingReasonType = 2  // 用户本地输入校验码停止充电
	StoppingReasonTypeRemoteUser                                   StoppingReasonType = 3  // 用户远程结束
	StoppingReasonTypeRemoteAdmin                                  StoppingReasonType = 4  // 管理员远程结束
	StoppingReasonTypeEmergencyStop                                StoppingReasonType = 5  // 急停按下，停止充电
	StoppingReasonTypeEVDisconnected                               StoppingReasonType = 6  // 检测到枪头断开，停止充电
	StoppingReasonTypeReboot                                       StoppingReasonType = 7  // 系统重启停止充电
	StoppingReasonTypeOffLine                                      StoppingReasonType = 8  // 充电桩断线，停止充电（对于不允许离线充电的充电桩）
	StoppingReasonTypePowerLoss                                    StoppingReasonType = 9  // 充电桩掉电，停止充电
	StoppingReasonTypeSystemFault                                  StoppingReasonType = 10 // 充电桩故障，停止充电
	StoppingReasonTypeBMSFault                                     StoppingReasonType = 11 // 电动车故障，停止充电
	StoppingReasonTypeOther                                        StoppingReasonType = 12 // 其它原因，停止充电
	StoppingReasonTypeNotEnoughMoney                               StoppingReasonType = 17 // 余额不足
	StoppingReasonTypeOverLoad                                     StoppingReasonType = 18 // 过载停止
	StoppingReasonTypeOverVoltage                                  StoppingReasonType = 19 // 过压停止
	StoppingReasonTypeUnderVoltage                                 StoppingReasonType = 20 // 欠压停止
	StoppingReasonTypeNetTimeout                                   StoppingReasonType = 21 // 网络同步超时停止充电
	StoppingReasonTypeCPError                                      StoppingReasonType = 25 // CP错误
	StoppingReasonTypeLeakage                                      StoppingReasonType = 26 // 漏电故障
	StoppingReasonTypeSocFull                                      StoppingReasonType = 16 // soc充满停止
	StoppingReasonTypeOverTime                                     StoppingReasonType = 22 // 超过设置时间
	StoppingReasonTypeOverMeter                                    StoppingReasonType = 23 // 超过设置电量
	StoppingReasonTypeOverSOC                                      StoppingReasonType = 24 // 超过设置SOC
	StoppingReasonTypeLightningProtection                          StoppingReasonType = 27 // 避雷
	StoppingReasonTypeAmbientOvertemperature                       StoppingReasonType = 28 // 环境过温
	StoppingReasonTypeMeterCommunicationFault                      StoppingReasonType = 13 // 电表通讯故障
	StoppingReasonTypeAccessControlFault                           StoppingReasonType = 29 // 门禁故障
	StoppingReasonTypeAbnormalCommunicationOfCardReader            StoppingReasonType = 30 // 刷卡器通讯异常
	StoppingReasonTypeCC1StatusError                               StoppingReasonType = 31 // CC1状态错误
	StoppingReasonTypeOpenSolenoidLockError                        StoppingReasonType = 32 // 开启电磁锁错误
	StoppingReasonTypeAbnormalElectronicLock                       StoppingReasonType = 33 // 电子锁异常
	StoppingReasonTypeAdaptiveVoltageError                         StoppingReasonType = 34 // 适配电压错误
	StoppingReasonTypeChargingRelayOpeningError                    StoppingReasonType = 35 // 充电器开启错误
	StoppingReasonTypeChargingRelayClosingError                    StoppingReasonType = 36 // 充电器关闭错误
	StoppingReasonTypeInsulationDetectionError                     StoppingReasonType = 37 // 绝缘监测错误
	StoppingReasonTypePreChargeVoltageAndMessageFivePercentError   StoppingReasonType = 38 // 预充电池电压和报文差5%错误
	StoppingReasonTypePreChargeVoltageOverMaxVoltageError          StoppingReasonType = 39 // 预充电压大于允许最大电压
	StoppingReasonTypePreChargeVoltageUnderMinVoltageError         StoppingReasonType = 40 // 预充电压小于允许最小电压
	StoppingReasonTypePreChargeAdapterVoltageError                 StoppingReasonType = 41 // 预充适配电压错误
	StoppingReasonTypeStopCurrentOverFiveError                     StoppingReasonType = 42 // 停止充电电流大于5A
	StoppingReasonTypeConnectionRelayOpenError                     StoppingReasonType = 43 // 充电继电器开启错误
	StoppingReasonTypeConnectionRelayClosingError                  StoppingReasonType = 44 // 充电继电器关闭错误
	StoppingReasonTypeOutputOverVoltageError                       StoppingReasonType = 45 // 输出过压错误
	StoppingReasonTypeOutputOverCurrentError                       StoppingReasonType = 46 // 输出过压错误
	StoppingReasonTypeConnectorOvertemperature                     StoppingReasonType = 47 // 枪过温
	StoppingReasonTypeChargingStationOtherFault                    StoppingReasonType = 48 // 充电桩其他故障
	StoppingReasonTypeBMSNotReady                                  StoppingReasonType = 49 // BMS未就绪
	StoppingReasonTypeBHMTimeoutError                              StoppingReasonType = 50 // BHM超时错误
	StoppingReasonTypeBRMTimeoutError                              StoppingReasonType = 51 // BRM超时错误
	StoppingReasonTypeBCPTimeoutError                              StoppingReasonType = 52 // BCP超时错误
	StoppingReasonTypeBROTimeoutError                              StoppingReasonType = 53 // BRO超时错误
	StoppingReasonTypeBROAATimeoutError                            StoppingReasonType = 54 // BROAA超时错误
	StoppingReasonTypeBROSeriousFault                              StoppingReasonType = 55 // BROAA重大错误停止充电
	StoppingReasonTypeBCLTimeoutError                              StoppingReasonType = 56 // BCL超时错误
	StoppingReasonTypeBCSTimeoutError                              StoppingReasonType = 57 // BCS超时错误
	StoppingReasonTypeBSMTimeoutError                              StoppingReasonType = 58 // BSM超时错误
	StoppingReasonTypeBSTTimeoutError                              StoppingReasonType = 59 // BST超时错误
	StoppingReasonTypeBSDTimeoutError                              StoppingReasonType = 60 // BSD超时错误
	StoppingReasonTypeBEMTimeoutError                              StoppingReasonType = 61 // BEM充电错误报文超时
	StoppingReasonTypeBSTInsulationFault                           StoppingReasonType = 62 // BST绝缘故障
	StoppingReasonTypeBSTComponentOvertemperature                  StoppingReasonType = 63 // BST元件过温
	StoppingReasonTypeBSTConnectorFault                            StoppingReasonType = 64 // BST连接器故障
	StoppingReasonTypeBSTBatteryGroupOvertemperature               StoppingReasonType = 65 // BST电池组过温
	StoppingReasonTypeBSTHighVoltageRelayFault                     StoppingReasonType = 66 // BST高压继电器故障
	StoppingReasonTypeBSTCheckpointTwoVoltageFault                 StoppingReasonType = 67 // BST检测点2电压故障
	StoppingReasonTypeBSTOvercurrent                               StoppingReasonType = 68 // BST电流过大
	StoppingReasonTypeBSTOvervoltage                               StoppingReasonType = 69 // BST电压过大
	StoppingReasonTypeSinglePowerBatteryOvervoltage                StoppingReasonType = 70 // 单体动力蓄电池电压过高
	StoppingReasonTypeSinglePowerBatteryUndervoltage               StoppingReasonType = 71 // 单体动力蓄电池电压过低
	StoppingReasonTypeEntireVehiclePowerBatteryChargeStateOverSoC  StoppingReasonType = 72 // 整车动力蓄电池荷电状态SoC过高
	StoppingReasonTypeEntireVehiclePowerBatteryChargeStateUnderSoC StoppingReasonType = 73 // 整车动力蓄电池喝点状态SoC过低
	StoppingReasonTypeMaximumBatteryPackOvertemperature            StoppingReasonType = 74 // 最高电池组过温
	StoppingReasonTypePowerBatteryChargingOvercurrent              StoppingReasonType = 75 // 动力蓄电池充电过流
	StoppingReasonTypePowerBatteryChargingOvertemperature          StoppingReasonType = 76 // 动力蓄电池温度过高
	StoppingReasonTypePowerBatteryIsulationFault                   StoppingReasonType = 77 // 动力蓄电池绝缘故障
	StoppingReasonTypePowerBatteryConnectorFault                   StoppingReasonType = 78 // 动力蓄电池连接器故障
	StoppingReasonTypeChargeNotAllowedStatus                       StoppingReasonType = 79 // 充电不允许的状态
	StoppingReasonTypeBMSOhterFault                                StoppingReasonType = 80 // BMS其他故障
	StoppingReasonTypePriceSchemeException                         StoppingReasonType = 81 // 计费模板异常
	// StoppingReasonTypeFullOfSelfStop                               StoppingReasonType = 101 // 充满自停(原22)
	// StoppingReasonTypeAdminForceStop                               StoppingReasonType = 102 // 管理员强制结束(原23)
	// StoppingReasonTypeEVOccupy                                     StoppingReasonType = 103 // 接口已被占用(原24)
	// StoppingReasonTypeNoTariff                                     StoppingReasonType = 104 // 没有计费模版(原25)
	// StoppingReasonTypeConnectorFault                               StoppingReasonType = 105 // 充电枪故障(原26)
	// StoppingReasonTypeCC1NoConnect                                 StoppingReasonType = 106 // 充电枪未连接(充电中车辆控制引导异常/CC1连接异常)(原27)
	// StoppingReasonTypeDeauthorized                                 StoppingReasonType = 107 // 取消授权
	StoppingReasonTypeEVConnectTimeout StoppingReasonType = 108 // 连接超时
	StoppingReasonTypeServer_Error     StoppingReasonType = 999 // 平台错误
	// StoppingReasonTypeHardReset                                    StoppingReasonType = 82
	// StoppingReasonTypeSoftReset                                    StoppingReasonType = 83
	// StoppingReasonTypeUnlockCommand                                StoppingReasonType = 84
	// StoppingReasonTypeEnergyLimitReached                           StoppingReasonType = 85
	// StoppingReasonTypeGroundFault                                  StoppingReasonType = 86
	// StoppingReasonTypeImmediateReset                               StoppingReasonType = 87
	// StoppingReasonTypeLocal                                        StoppingReasonType = 88
	// StoppingReasonTypeLocalOutOfCredit                             StoppingReasonType = 89
	// StoppingReasonTypeMasterPass                                   StoppingReasonType = 90
	// StoppingReasonTypeOvercurrentFault                             StoppingReasonType = 91
	// StoppingReasonTypePowerQuality                                 StoppingReasonType = 92
	// StoppingReasonTypeSocLimitReached                              StoppingReasonType = 93
	// StoppingReasonTypeStoppedByEV                                  StoppingReasonType = 94
	// StoppingReasonTypeStoppedByScreen                              StoppingReasonType = 95
	StoppingReasonScreenManualStop                                    StoppingReasonType = 111 // 屏幕手动停止
	StoppingReasonInsufficientBalance                                 StoppingReasonType = 112 // 余额不足
	StoppingReasonReachedSetChargingAmount                            StoppingReasonType = 113 // 达到设置的充电金额停止
	StoppingReasonInvalidCurrentStop                                  StoppingReasonType = 114 // 无有效电流停止
	StoppingReasonBMSAbnormalTermination                              StoppingReasonType = 115 // BMS异常终止充电
	StoppingReasonReachedTerminationCondition                         StoppingReasonType = 116 // 充电桩达到终止条件停止
	StoppingReasonGunNotConnected                                     StoppingReasonType = 117 // 枪未正确连接
	StoppingReasonReachedOfflineStop                                  StoppingReasonType = 118 // 达到离线充电停止
	StoppingReasonUnstoppedOrder                                      StoppingReasonType = 119 // 有未停止订单
	StoppingReasonCreateOrderAbnormal                                 StoppingReasonType = 120 // 创建订单异常
	StoppingReasonPileNotExistOrDisabled                              StoppingReasonType = 121 // 桩不存在或已禁用
	StoppingReasonPlatformUnableToStartCharging                       StoppingReasonType = 122 // 平台无法启动充电
	StoppingReasonPileStartupTimeout                                  StoppingReasonType = 123 // 桩启动响应超时
	StoppingReasonPileStopTimeout                                     StoppingReasonType = 124 // 桩停止响应超时
	StoppingReasonPilePlatformCommunicationEstablished                StoppingReasonType = 125 // 桩与平台通讯建立
	StoppingReasonPilePlatformCommunicationDisconnected               StoppingReasonType = 126 // 桩与平台通讯断开
	StoppingReasonVINAuthFailure                                      StoppingReasonType = 127 // 车辆VIN鉴权失败
	StoppingReasonFirmwareUpgradeFailure                              StoppingReasonType = 128 // 固件升级失败
	StoppingReasonUpgradePackageException                             StoppingReasonType = 129 // 升级包异常
	StoppingReasonUnpaidOrderStartFailure                             StoppingReasonType = 130 // 订单未支付，启动失败
	StoppingReasonPileConnectionChargingCloudAuthFailure              StoppingReasonType = 131 // 桩连接充电云鉴权失败
	StoppingReasonUserBlacklisted                                     StoppingReasonType = 132 // 户在黑名单中，禁止充电
	StoppingReasonOrderTerminationCommandInvalid                      StoppingReasonType = 133 // 指令要求终止的订单号不存在或者和目标充电口当前订单不一致
	StoppingReasonSystemUpgrading                                     StoppingReasonType = 134 // 系统正在升级中，请稍后再试
	StoppingReasonSystemMaintenance                                   StoppingReasonType = 135 // 系统维护中，请稍后再试
	StoppingReasonBackendParameterError                               StoppingReasonType = 136 // 后台下发参数异常
	StoppingReasonDeviceSelfCheckTimeout                              StoppingReasonType = 137 // 设备自检超时
	StoppingReasonPileOffline                                         StoppingReasonType = 138 // 桩离线
	StoppingReasonSystemFanFault                                      StoppingReasonType = 139 // 系统风扇故障
	StoppingReasonModuleFanFault                                      StoppingReasonType = 140 // 模块风扇故障
	StoppingReasonChargingGunNotReturned                              StoppingReasonType = 141 // 充电枪未归位告警
	StoppingReasonModuleCommunicationFault                            StoppingReasonType = 142 // 模块通讯故障
	StoppingReasonPowerModuleAddressConflict                          StoppingReasonType = 143 // 电源模块地址冲突
	StoppingReasonPowerModuleFault                                    StoppingReasonType = 144 // 电源模块故障
	StoppingReasonPowerModuleOverTemperature                          StoppingReasonType = 145 // 电源模块过温告警
	StoppingReasonNoIdleModuleAvailable                               StoppingReasonType = 146 // 无空闲模块可用（限智能分配功率）
	StoppingReasonMeterDataAbnormal                                   StoppingReasonType = 147 // 电表数据异常
	StoppingReasonOutputContactorSticking                             StoppingReasonType = 148 // 输出接触器粘连
	StoppingReasonDCContactorFault                                    StoppingReasonType = 149 // 直流接触器故障
	StoppingReasonDCFuseFault                                         StoppingReasonType = 150 // 直流熔断器故障
	StoppingReasonIntermediateRelayFault                              StoppingReasonType = 151 // 中间继电器故障
	StoppingReasonAuxiliaryPowerFault                                 StoppingReasonType = 152 // 辅助电源故障
	StoppingReasonDischargeCircuitFault                               StoppingReasonType = 153 // 泄放回路故障
	StoppingReasonLiquidLevelAlarm                                    StoppingReasonType = 154 // 液位报警
	StoppingReasonManualStopCharging                                  StoppingReasonType = 155 // 手动停止充电
	StoppingReasonChargingConnectionFault                             StoppingReasonType = 156 // 充电连接故障
	StoppingReasonGunMouthAbnormal                                    StoppingReasonType = 157 // 枪口异常
	StoppingReasonParkingLockFault                                    StoppingReasonType = 158 // 车位锁故障
	StoppingReasonParkingLockBatteryDepleted                          StoppingReasonType = 159 // 车位锁电池耗尽
	StoppingReasonParkingLockLockingFailed                            StoppingReasonType = 160 // 车位锁落锁失败
	StoppingReasonRemotePowerDistributionFailed                       StoppingReasonType = 162 // 充电桩执行远程功率分配策略失败
	StoppingReasonTooManyFailures                                     StoppingReasonType = 163 // 当前用户失败次数过多（>=3次），请更换充电桩
	StoppingReasonDeviceDamaged                                       StoppingReasonType = 164 // 充电桩设备损坏，请更换充电桩
	StoppingReasonDeviceSuspended                                     StoppingReasonType = 165 // 充电桩暂停使用，请更换充电桩
	StoppingReasonACContactorFault                                    StoppingReasonType = 166 // 交流接触器故障
	StoppingReasonGunHeadInsertionRemovalWarning                      StoppingReasonType = 167 // 枪头插拔次数预警
	StoppingReasonSelfCheckPowerDistributionTimeout                   StoppingReasonType = 168 // 自检功率分配超时
	StoppingReasonMainContactSticking                                 StoppingReasonType = 169 // 母联粘连故障
	StoppingReasonPrechargeCompletionTimeout                          StoppingReasonType = 170 // 预充完成超时
	StoppingReasonChargingStartTimeout                                StoppingReasonType = 171 // 启动充电超时
	StoppingReasonChargingStartResponseFailed                         StoppingReasonType = 172 // 启动完成应答失败
	StoppingReasonModulePowerOnTimeout                                StoppingReasonType = 173 // 模块开机超时
	StoppingReasonBillingControlUnitCommFault                         StoppingReasonType = 174 // 计费控制单元通讯故障
	StoppingReasonEnvironmentMonitoringBoardCommFault                 StoppingReasonType = 175 // 环境监控板通讯故障
	StoppingReasonAirConditioningCommFault                            StoppingReasonType = 176 // 空调通讯故障
	StoppingReasonPassiveOutputBoxCommFault                           StoppingReasonType = 177 // 无源开出盒通讯故障
	StoppingReasonPassiveInputBoxCommFault                            StoppingReasonType = 178 // 无源开入盒通讯故障
	StoppingReasonInsulationSamplingBoxCommFault                      StoppingReasonType = 179 // 绝缘采样盒通讯故障
	StoppingReasonDCSamplingBoxCommFault                              StoppingReasonType = 180 // 直流采样盒通讯故障
	StoppingReasonGuidanceBoardCommFault                              StoppingReasonType = 181 // 导引板通讯故障
	StoppingReasonLightBoardCommFault                                 StoppingReasonType = 182 // 灯板通讯故障
	StoppingReasonPrechargePowerDistributionTimeout                   StoppingReasonType = 183 // 预充功率分配超时
	StoppingReasonSynchronizedGunFault                                StoppingReasonType = 184 // 并充枪同步过来的故障
	StoppingReasonConcentratorSystemFault                             StoppingReasonType = 185 // 集中器系统故障
	StoppingReasonZHM13SCommFault                                     StoppingReasonType = 186 // zhm13s通讯故障
	StoppingReasonTHSB02CommFault                                     StoppingReasonType = 187 // thsb02通讯故障
	StoppingReasonZHIM03CommFault                                     StoppingReasonType = 188 // zhim03通讯故障
	StoppingReasonChargingPileWaterIngress                            StoppingReasonType = 189 // 充电桩水浸故障
	StoppingReasonChargingCabinetWaterIngress                         StoppingReasonType = 190 // 充电柜水浸故障
	StoppingReasonSyncChargeCommTimeout                               StoppingReasonType = 191 // 并充通讯超时
	StoppingReasonSyncChargeOtherGunFault                             StoppingReasonType = 192 // 并充的另一把枪故障
	StoppingReasonSyncChargeTimeout                                   StoppingReasonType = 193 // 并充同步超时
	StoppingReasonSyncChargeStartMethodAbnormal                       StoppingReasonType = 194 // 并充启动方式异常
	StoppingReasonChargingCabinetDoorFault                            StoppingReasonType = 195 // 充电机柜门禁故障
	StoppingReasonOutputVoltageUndervoltageFault                      StoppingReasonType = 196 // 输出电压欠压故障
	StoppingReasonOutputShortCircuitFault                             StoppingReasonType = 197 // 输出短路故障
	StoppingReasonACBreakerFault                                      StoppingReasonType = 198 // 交流断路器故障
	StoppingReasonRelayExternalVoltageAbove10V                        StoppingReasonType = 199 // 继电器外侧电压大于10V
	StoppingReasonTestPointVoltageDetectionFault                      StoppingReasonType = 200 // 检测点电压检测故障
	StoppingReasonPileGroupCapacityExceedsRatedLimit                  StoppingReasonType = 201 // 桩群电容量超过额定限制
	StoppingReasonInputPhaseLossAlarm                                 StoppingReasonType = 202 // 输入缺相报警
	StoppingReasonElectricalLeakageProtection                         StoppingReasonType = 203 // 漏电保护
	StoppingReasonGroundWireAlarm                                     StoppingReasonType = 204 // 地线报警
	StoppingReasonACSurgeProtectionAlarm                              StoppingReasonType = 205 // 交流防雷报警
	StoppingReasonOtherPowerFailure                                   StoppingReasonType = 206 // 其他电源故障
	StoppingReasonCarPileVoltageAbnormal                              StoppingReasonType = 207 // 车/桩电压异常
	StoppingReasonModuleProtection                                    StoppingReasonType = 208 // 模块保护
	StoppingReasonThreePhaseImbalance                                 StoppingReasonType = 209 // 三相不平衡
	StoppingReasonACInputOvervoltage                                  StoppingReasonType = 210 // 交流输入过压
	StoppingReasonACInputUndervoltage                                 StoppingReasonType = 211 // 交流输入欠压
	StoppingReasonModuleInputFailure                                  StoppingReasonType = 212 // 模块输入故障
	StoppingReasonBMSCommunicationException                           StoppingReasonType = 213 // BMS通讯异常
	StoppingReasonBHMOutputMismatch                                   StoppingReasonType = 214 // BHM桩的输出能力不匹配
	StoppingReasonBMSVoltageDemandTooLow                              StoppingReasonType = 215 // BMS需求电压过低
	StoppingReasonBMSTemperatureOverlimit                             StoppingReasonType = 216 // BMS元件过温
	StoppingReasonBatteryReverseConnectionFailure                     StoppingReasonType = 217 // 电池反接故障
	StoppingReasonBatteryVoltageAbnormal                              StoppingReasonType = 218 // 电池电压异常
	StoppingReasonCROOutputReadyTimeout                               StoppingReasonType = 219 // CRO充电机输出就绪超时
	StoppingReasonCCSStatusMessageTimeout                             StoppingReasonType = 220 // CCS充电机状态报文超时
	StoppingReasonCSTTerminationMessageTimeout                        StoppingReasonType = 221 // CST充电机终止充电报文超时
	StoppingReasonCSDStatisticsMessageTimeout                         StoppingReasonType = 222 // CSD充电统计数据报文超时
	StoppingReasonVehicleCurrentMismatch                              StoppingReasonType = 223 // 车辆电流不匹配 电流过大
	StoppingReasonVehicleChargeNotTransferable                        StoppingReasonType = 224 // 车辆电量无法传送
	StoppingReasonVehicleOccupationTimeout                            StoppingReasonType = 225 // 车辆占位超时
	StoppingReasonBMSNewOldStandardDetectionTimeout                   StoppingReasonType = 226 // BMS新老国标探测超时
	StoppingReasonBMSStartChargeTimeout                               StoppingReasonType = 227 // BMS启动充电超时--未发送BCS和BCL
	StoppingReasonBMSHighVoltageRelayFailure                          StoppingReasonType = 228 // BMS高压继电器故障
	StoppingReasonBMSMonitorPoint2VoltageFailure                      StoppingReasonType = 229 // BMS监测点2电压检测故障
	StoppingReasonBROReadyCancelled                                   StoppingReasonType = 230 // BRO准备就绪后取消
	StoppingReasonBMSPauseTimeout                                     StoppingReasonType = 231 // BMS暂停超时
	StoppingReasonVehicleBCPVoltageMismatch                           StoppingReasonType = 232 // 车辆BCP报文和实际电压不符
	StoppingReasonPressureAdjustmentFailure                           StoppingReasonType = 233 // 预充阶段调压失败
	StoppingReasonPreChargeWaitTimeout                                StoppingReasonType = 234 // 预充阶段等待BCL和BCS超时
	StoppingReasonBatterySoftStartFailure                             StoppingReasonType = 235 // 电池软起失败
	StoppingReasonBSTOutputConnectorOvertemp                          StoppingReasonType = 236 // BST输出连接器过温
	StoppingReasonBSTChargeConnectorFailure                           StoppingReasonType = 237 // BST充电连接器故障
	StoppingReasonBSTOtherFailure                                     StoppingReasonType = 238 // BST其他故障
	StoppingReasonBSTMonitorPoint2VoltageFailure                      StoppingReasonType = 239 // BST监测点2电压检测故障
	StoppingReasonBatteryMaxVoltageLessThanMinPileVoltage             StoppingReasonType = 241 // 电池最高电压小于桩最小输出电压
	StoppingReasonBatteryCurrentVoltageLessThanMinPileVoltage         StoppingReasonType = 242 // 电池当前电压小于桩最小输出电压
	StoppingReasonBatteryVoltageGreaterThanMaxPileVoltage             StoppingReasonType = 243 // 电池电压大于桩最大输出电压
	StoppingReasonBatteryVoltageGreaterThanBCPMaxAllowedChargeVoltage StoppingReasonType = 244 // 电池电压大于BCP最高允许充电电压
	StoppingReasonBCPDataAbnormal                                     StoppingReasonType = 246 // BCP数据异常
	StoppingReasonInternalCommunicationFailure                        StoppingReasonType = 247 // 充电桩内部通讯故障
	StoppingReasonBMSVINAbnormalStop                                  StoppingReasonType = 248 // BMS VIN异常停止
	StoppingReasonPrechargeStartModuleFailure                         StoppingReasonType = 249 // 预充电启动模块失败
	StoppingReasonPileJudgedBMSOtherFailure                           StoppingReasonType = 250 // 充电桩判断BMS其他故障停止充电
	StoppingReasonParallelContactorFailure                            StoppingReasonType = 251 // 并联接触器据动/误动故障
	StoppingReasonOutputCurrentExceedsMaxAllowed                      StoppingReasonType = 252 // 输出电流大于最高允许充电电流
	StoppingReasonMonitorCommunicationError                           StoppingReasonType = 253 // 监控之间通信出错
	StoppingReasonOutputOvercurrentReverse                            StoppingReasonType = 254 // 输出过流倒送
	StoppingReasonInsulationCheckBatteryVoltageLow                    StoppingReasonType = 255 // 绝缘检查电池电压未达预设值
	StoppingReasonSmokeSensorFailure                                  StoppingReasonType = 256 // 烟感故障
	StoppingReasonBSMBatteryTempExceedsMax                            StoppingReasonType = 257 // BSM电池温度大于最高允许
	StoppingReasonPartialChargingModuleFailure                        StoppingReasonType = 258 // 充电模块故障（部分）
	StoppingReasonChargingModuleACOvervoltage                         StoppingReasonType = 259 // 充电模块交流过压
	StoppingReasonChargingModuleACUndervoltage                        StoppingReasonType = 260 // 充电模块交流欠压
	StoppingReasonChargingModuleShortCircuit                          StoppingReasonType = 261 // 充电模块短路故障
	StoppingReasonInsulationMonitorAlarm                              StoppingReasonType = 262 // 绝缘监视告警
	StoppingReasonOutputContactorExternalVoltage                      StoppingReasonType = 263 // 输出接触器外侧电压＞10V
	StoppingReasonDischargeTimeout                                    StoppingReasonType = 264 // 泄放超时
	StoppingReasonPrechargeBatteryVoltageLow                          StoppingReasonType = 265 // 预充电电池电压过低
	StoppingReasonChargingVehicleControlPilotFault                    StoppingReasonType = 266 // 充电中车辆控制导引故障
	StoppingReasonPrechargeK5K6PositionAbnormal                       StoppingReasonType = 267 // 预充电K5K6位置异常
	StoppingReasonSystemReset                                         StoppingReasonType = 268 // 系统复位
	StoppingReasonBatteryVoltageLowOrMismatch                         StoppingReasonType = 269 // 电池电压过低或与上送值不符
	StoppingReasonVehicleProactiveStop                                StoppingReasonType = 270 // 车辆主动停止
	StoppingReasonBMSMaxAllowedChargeVoltageLow                       StoppingReasonType = 271 // BMS最高允许充电电压过低
	StoppingReasonBSTBMSComponentFailure                              StoppingReasonType = 272 // BST-BMS元件故障
	StoppingReasonBSTVoltageAbnormal                                  StoppingReasonType = 273 // BST电压异常
	StoppingReasonWaterImmersionFault                                 StoppingReasonType = 274 // 水浸故障
	StoppingReasonBSTReachedSOCGoal                                   StoppingReasonType = 275 // BST达到所需求SOC目标值
	StoppingReasonBSTReachedTotalVoltageSetpoint                      StoppingReasonType = 276 // BST达到总电压设定值
	StoppingReasonBSTReachedSingleVoltageSetpoint                     StoppingReasonType = 277 // BST达到单体电压设定值
	StoppingReasonBCLChargeModeAbnormal                               StoppingReasonType = 278 // BCL充电模式异常
	StoppingReasonTemperatureHumidityFault                            StoppingReasonType = 279 // 温湿度故障
	StoppingReasonChargeSOCCompletion                                 StoppingReasonType = 280 // 充电SOC完成 按设定的SOC充电
	StoppingReasonGunHeadVoltageGreaterThan60V                        StoppingReasonType = 281 // 停机后枪头电压大于60V
	StoppingReasonBCSSingleBatteryVoltageTooHigh                      StoppingReasonType = 282 // BCS单体电池电压过高
	StoppingReasonBSMBatteryTemperatureAbnormal                       StoppingReasonType = 283 // BSM电池温度异常
	StoppingReasonCPAbnormal                                          StoppingReasonType = 284 // CP异常
	StoppingReasonTipoverFault                                        StoppingReasonType = 285 // 倾倒故障

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

var _ services.Request = &equipStopTransactionRequest{}

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

type StopTransactionRequestConfig struct {
	services.ReusedConfig
	TransactionId string
	Offline       bool
	StopReason    StoppingReasonType
	ChargingState uint8
	Timestamp     int64
}

func NewEquipStopTransactionRequestWithConfig(config *StopTransactionRequestConfig) *equipStopTransactionRequest {
	return &equipStopTransactionRequest{
		Base: services.Base{
			EquipmentSn: config.Sn,
			Protocol:    config.Protocol,
			Category:    services.StopTransaction.FirstUpper(),
			AccessPod:   config.Pod,
			MsgID:       config.MsgID,
		},
		Data: &equipStopTransactionRequestDetail{
			IdTokenType:   &IdTokenType{},
			StopReason:    config.StopReason,
			TransactionId: config.TransactionId,
			Timestamp:     config.Timestamp,
			Offline:       config.Offline,
			MeterValue:    &MeterValue{},
			Tariff: &Tariff{
				Id: -1,
			},
		},
	}
}

func (equipStopTransactionRequest) GetName() services.Request2ServicesNameType {
	return services.StopTransaction
}

func (e *equipStopTransactionRequest) TraceId() string {
	return e.MsgID
}

func (equipStopTransactionRequest) IsCallback() bool {
	return false
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
	req.Data.Tariff = &Tariff{
		Id: -1,
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

func StopTransactionRequest(ctx context.Context, req *equipStopTransactionRequest) (*equipStopTransactionResponse, error) {
	header := services.GetSimpleHeaderValue(services.StopTransaction)

	url := services.GetSimpleURL(req)

	return services.RequestWithResponse(ctx, req, url, header, &equipStopTransactionResponse{})
}
