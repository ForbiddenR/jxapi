package equip

import (
	"context"

	api "github.com/ForbiddenR/jx-api"
	"github.com/ForbiddenR/jx-api/services"
	// ocpp16 "gitee.com/csms/jxeu-ocpp/pkg/ocpp1.6/protocol"
	// ocpp201 "gitee.com/csms/jxeu-ocpp/pkg/ocpp2.0.1/protocol"
)

type equipStatusNotificationRequest struct {
	services.Base
	Data *equipStatusNotificationRequestDetail `json:"data"`
}

type equipStatusNotificationRequestDetail struct {
	EvseSerial      *int                    `json:"evseSerial,omitempty"`
	ConnectorSerial string                  `json:"connectorSerial"`
	Status          ConnectorStatusTypeEnum `json:"status"`
	ErrorCode       *string                 `json:"errorCode,omitempty"`
	VendorErrorCode *string                 `json:"vendorErrorCode"`
	Timestamp       int64                   `json:"timestamp"`
}

func NewEquipStatusNotificationRequestOCPP16(sn, pod, msgID string, connectorId string, status ConnectorStatusTypeEnum, errorCode StatusNotificationErrorCodeEnum, timestamp int64) *equipStatusNotificationRequest {
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

func NewEquipStatusNotificationRequest(sn, pod, msgID string, p *services.Protocol, connectorId string, status ConnectorStatusTypeEnum, timestamp int64) *equipStatusNotificationRequest {
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

func (o *equipStatusNotificationRequest) GetName() string {
	return services.StatusNotification.String()
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

func StatusNotificationRequestWithGeneric(ctx context.Context, req *equipStatusNotificationRequest) error {
	header := services.GetSimpleHeaderValue(services.Register)

	url := services.GetSimpleURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipStatusNotificationResponse{})
}
