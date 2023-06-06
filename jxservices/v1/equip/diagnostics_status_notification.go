package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type DiagnosticsStatusNotificationType int

const (
	DiagnosticsStatusNotificationTypeIdle         DiagnosticsStatusNotificationType = 1
	DiagnosticsStatusNotificationTypeUploadFailed DiagnosticsStatusNotificationType = 2
	DiagnosticsStatusNotificationTypeUploaded     DiagnosticsStatusNotificationType = 3
	DiagnosticsStatusNotificationTypeUploading    DiagnosticsStatusNotificationType = 4
)

// func OCPP16GetDiagnosticsStatus(status protocol.DiagnosticsStatusNotificationJsonStatus) DiagnosticsStatusNotificationType {
// 	var result DiagnosticsStatusNotificationType
// 	switch status {
// 	case protocol.DiagnosticsStatusNotificationJsonStatusIdle:
// 		result = DiagnosticsStatusNotificationTypeIdle
// 	case protocol.DiagnosticsStatusNotificationJsonStatusUploadFailed:
// 		result = DiagnosticsStatusNotificationTypeUploadFailed
// 	case protocol.DiagnosticsStatusNotificationJsonStatusUploaded:
// 		result = DiagnosticsStatusNotificationTypeUploaded
// 	case protocol.DiagnosticsStatusNotificationJsonStatusUploading:
// 		result = DiagnosticsStatusNotificationTypeUploading
// 	}
// 	return result
// }

type equipDiagnosticsStatusNotificationRequest struct {
	services.Base
	Data *equipDiagnosticsStatusNotificationRequestDetail `json:"data"`
}

func (*equipDiagnosticsStatusNotificationRequest) GetName() string {
	return services.DiagnosticsStatusNotification.String()
}

type equipDiagnosticsStatusNotificationRequestDetail struct {
	RequestId *int64                            `json:"requestId,omitempty"`
	Status    DiagnosticsStatusNotificationType `json:"status"`
}

func NewEquipDiagnosticsStatusNotificationRequestOCPP16(sn, pod, msgID string, status DiagnosticsStatusNotificationType) *equipDiagnosticsStatusNotificationRequest {
	req := &equipDiagnosticsStatusNotificationRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    services.OCPP16(),
			Category:    services.DiagnosticsStatusNotification.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipDiagnosticsStatusNotificationRequestDetail{
			Status: status,
		},
	}
	return req
}

func NewEquipDiagnosticsStatusNotificationRequest(sn, pod, msgID string, p *services.Protocol, requestId int64, status DiagnosticsStatusNotificationType) *equipDiagnosticsStatusNotificationRequest {
	req := &equipDiagnosticsStatusNotificationRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.DiagnosticsStatusNotification.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipDiagnosticsStatusNotificationRequestDetail{
			RequestId: &requestId,
			Status:    status,
		},
	}
	return req
}

var _ services.Response = &equipDiagnosticsStatusNotificationResponse{}

type equipDiagnosticsStatusNotificationResponse struct {
	api.Response
	Data *equipDiagnosticsStatusNotificationResponseDetail `json:"data"`
}

func (resp *equipDiagnosticsStatusNotificationResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipDiagnosticsStatusNotificationResponse) GetMsg() string {
	return resp.Msg
}

type equipDiagnosticsStatusNotificationResponseDetail struct {
}

func DiagnosticsStatusNotificationRequestG(ctx context.Context, req services.Request) error {
	header := services.GetSimpleHeaderValue(services.DiagnosticsStatusNotification)

	url := services.GetSimpleURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipDiagnosticsStatusNotificationResponse{})
}
