package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

var _ services.Request = &equipStatusNotificationRequest{}

type equipStatusNotificationRequest struct {
	services.Base
	Data *equipStatusNotificationRequestDetail `json:"data"`
}

type equipStatusNotificationRequestDetail struct {
	EvseSerial      *int    `json:"evseSerial,omitempty"`
	ConnectorSerial string  `json:"connectorSerial"`
	Status          int     `json:"status"`
	ErrorCode       *string `json:"errorCode,omitempty"`
	VendorErrorCode *string `json:"vendorErrorCode"`
	Timestamp       int64   `json:"timestamp"`
}

type StatusNotificationRequestConfig struct {
	ConnectorSerial string
	Status          int
	Timestamp       int64
}

func NewEquipStatusNotificationRequestOCPP16(sn, pod, msgID string, connectorId string, status int, errorCode StatusNotificationErrorCodeEnum, timestamp int64) *equipStatusNotificationRequest {
	ec := string(errorCode)
	return &equipStatusNotificationRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    services.OCPP16(),
			Category:    services.StatusNotification.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipStatusNotificationRequestDetail{
			ConnectorSerial: connectorId,
			Status:          status,
			ErrorCode:       &ec,
			Timestamp:       timestamp,
		},
	}
}

func NewStatusNotification(base services.Base, config *StatusNotificationRequestConfig) *equipStatusNotificationRequest {
	req := &equipStatusNotificationRequest{
		Base: base,
		Data: &equipStatusNotificationRequestDetail{
			ConnectorSerial: config.ConnectorSerial,
			Status: config.Status,
			Timestamp: config.Timestamp,
		},
	}
	return req
}

func NewEquipStatusNotificationRequest(sn, pod, msgID string, p *services.Protocol, connectorId string, status int, timestamp int64) *equipStatusNotificationRequest {
	return &equipStatusNotificationRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.StatusNotification.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipStatusNotificationRequestDetail{
			ConnectorSerial: connectorId,
			Status:          status,
			Timestamp:       timestamp,
		},
	}
}

func (equipStatusNotificationRequest) GetName() services.Request2ServicesNameType {
	return services.StatusNotification
}

func (equipStatusNotificationRequest) IsCallback() bool {
	return false
}

type ConnectorStatusTypeEnum int

const (
	ConnectorStatusUnavailable ConnectorStatusTypeEnum = iota
	ConnectorStatusAvailable
	ConnectorStatusPreparing
	ConnectorStatusCharging
	ConnectorStatusSuspendedEV
	ConnectorStatusSuspendedEVSE
	ConnectorStatusFinishing
	ConnectorStatusFaulted
	ConnectorStatusReserved
	ConnectorStatusOccupied
)

type ConnectorStatus201TypeEnum int

const (
	ConnectorStatus201Unavailable   ConnectorStatus201TypeEnum = iota // 不可用
	ConnectorStatus201Available                                       // 空闲可用
	ConnectorStatus201Occupied                                        // 占用
	ConnectorStatus201Reserved                                        // 预约
	ConnectorStatus201Faulted                                         // 故障
	ConnectorStatus201Preparing                                       // 准备中
	ConnectorStatus201Charging                                        // 充电中
	ConnectorStatus201SuspendedEV                                     // 车端挂起，不输入电能
	ConnectorStatus201SuspendedEVSE                                   // 桩端挂起，不输出电能
	ConnectorStatus201Finishing                                       // 结束中
)

// func OCPP16ConnectorStatus(s ocpp16.StatusNotificationJsonStatus) ConnectorStatusTypeEnum {
// 	switch s {
// 	case ocpp16.StatusNotificationJsonStatusAvailable:
// 		return ConnectorStatusAvailable
// 	case ocpp16.StatusNotificationJsonStatusPreparing:
// 		return ConnectorStatusPreparing
// 	case ocpp16.StatusNotificationJsonStatusCharging:
// 		return ConnectorStatusCharging
// 	case ocpp16.StatusNotificationJsonStatusSuspendedEV:
// 		return ConnectorStatusSuspendedEV
// 	case ocpp16.StatusNotificationJsonStatusSuspendedEVSE:
// 		return ConnectorStatusSuspendedEVSE
// 	case ocpp16.StatusNotificationJsonStatusFinishing:
// 		return ConnectorStatusFinishing
// 	case ocpp16.StatusNotificationJsonStatusFaulted:
// 		return ConnectorStatusFaulted
// 	case ocpp16.StatusNotificationJsonStatusReserved:
// 		return ConnectorStatusReserved
// 	default:
// 		return ConnectorStatusUnavailable
// 	}
// }

// func OCPP201ConnectorStatus(s ocpp201.ConnectorStatusEnumType) ConnectorStatusTypeEnum {
// 	switch s {
// 	case ocpp201.ConnectorStatusEnumTypeAvailable:
// 		return ConnectorStatusAvailable
// 	case ocpp201.ConnectorStatusEnumTypeOccupied:
// 		return ConnectorStatusOccupied
// 	case ocpp201.ConnectorStatusEnumTypeFaulted:
// 		return ConnectorStatusFaulted
// 	case ocpp201.ConnectorStatusEnumTypeReserved:
// 		return ConnectorStatusReserved
// 	default:
// 		return ConnectorStatusUnavailable
// 	}
// }

type StatusNotificationErrorCodeEnum string

const (
	StatusNotificationErrorCodeNoError              StatusNotificationErrorCodeEnum = "NoError"
	StatusNotificationErrorCodeConnectorLockFailure StatusNotificationErrorCodeEnum = "ConnectorLockFailure"
	StatusNotificationErrorCodeEVCommunicationError StatusNotificationErrorCodeEnum = "EVCommunicationError"
	StatusNotificationErrorCodeGroundFailure        StatusNotificationErrorCodeEnum = "GroundFailure"
	StatusNotificationErrorCodeHighTemperature      StatusNotificationErrorCodeEnum = "HighTemperature"
	StatusNotificationErrorCodeInternalError        StatusNotificationErrorCodeEnum = "InternalError"
	StatusNotificationErrorCodeLocalListConflict    StatusNotificationErrorCodeEnum = "LocalListConflict"
	StatusNotificationErrorCodeOverCurrentFailure   StatusNotificationErrorCodeEnum = "OverCurrentFailure"
	StatusNotificationErrorCodeOverVoltage          StatusNotificationErrorCodeEnum = "OverVoltage"
	StatusNotificationErrorCodePowerMeterFailure    StatusNotificationErrorCodeEnum = "PowerMeterFailure"
	StatusNotificationErrorCodePowerSwitchFailure   StatusNotificationErrorCodeEnum = "PowerSwitchFailure"
	StatusNotificationErrorCodeReaderFailure        StatusNotificationErrorCodeEnum = "ReaderFailure"
	StatusNotificationErrorCodeResetFailure         StatusNotificationErrorCodeEnum = "ResetFailure"
	StatusNotificationErrorCodeUnderVoltage         StatusNotificationErrorCodeEnum = "UnderVoltage"
	StatusNotificationErrorCodeWeakSignal           StatusNotificationErrorCodeEnum = "WeakSignal"
	StatusNotificationErrorCodeOtherError           StatusNotificationErrorCodeEnum = "OtherError"
)

// func OCPP16StatusNotificationErrorCode(s ocpp16.StatusNotificationJsonErrorCode) StatusNotificationErrorCodeEnum {
// 	switch s {
// 	case ocpp16.StatusNotificationJsonErrorCodeNoError:
// 		return StatusNotificationErrorCodeNoError
// 	case ocpp16.StatusNotificationJsonErrorCodeConnectorLockFailure:
// 		return StatusNotificationErrorCodeConnectorLockFailure
// 	case ocpp16.StatusNotificationJsonErrorCodeEVCommunicationError:
// 		return StatusNotificationErrorCodeEVCommunicationError
// 	case ocpp16.StatusNotificationJsonErrorCodeGroundFailure:
// 		return StatusNotificationErrorCodeGroundFailure
// 	case ocpp16.StatusNotificationJsonErrorCodeHighTemperature:
// 		return StatusNotificationErrorCodeHighTemperature
// 	case ocpp16.StatusNotificationJsonErrorCodeInternalError:
// 		return StatusNotificationErrorCodeInternalError
// 	case ocpp16.StatusNotificationJsonErrorCodeLocalListConflict:
// 		return StatusNotificationErrorCodeLocalListConflict
// 	case ocpp16.StatusNotificationJsonErrorCodeOverCurrentFailure:
// 		return StatusNotificationErrorCodeOverCurrentFailure
// 	case ocpp16.StatusNotificationJsonErrorCodeOverVoltage:
// 		return StatusNotificationErrorCodeOverVoltage
// 	case ocpp16.StatusNotificationJsonErrorCodePowerMeterFailure:
// 		return StatusNotificationErrorCodePowerMeterFailure
// 	case ocpp16.StatusNotificationJsonErrorCodePowerSwitchFailure:
// 		return StatusNotificationErrorCodePowerSwitchFailure
// 	case ocpp16.StatusNotificationJsonErrorCodeReaderFailure:
// 		return StatusNotificationErrorCodeReaderFailure
// 	case ocpp16.StatusNotificationJsonErrorCodeResetFailure:
// 		return StatusNotificationErrorCodeResetFailure
// 	case ocpp16.StatusNotificationJsonErrorCodeWeakSignal:
// 		return StatusNotificationErrorCodeWeakSignal
// 	case ocpp16.StatusNotificationJsonErrorCodeUnderVoltage:
// 		return StatusNotificationErrorCodeUnderVoltage
// 	default:
// 		return StatusNotificationErrorCodeOtherError
// 	}
// }

var _ services.Response = &equipStatusNotificationResponse{}

type equipStatusNotificationResponse struct {
	api.Response
}

func (resp *equipStatusNotificationResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipStatusNotificationResponse) GetMsg() string {
	return resp.Msg
}

func StatusNotificationRequest(ctx context.Context, req *equipStatusNotificationRequest) error {
	header := services.GetSimpleHeaderValue(services.Register)

	url := services.GetSimpleURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipStatusNotificationResponse{})
}
