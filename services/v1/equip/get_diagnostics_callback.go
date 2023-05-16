package equip

import (
	"context"

	callback "gitee.com/csms/jxeu-ocpp/internal/errors"
	"gitee.com/csms/jxeu-ocpp/pkg/api"
	"gitee.com/csms/jxeu-ocpp/pkg/api/services"
)

type equipGetDiagnosticsCallbackRequest struct {
	services.Base
	Callback *equipGetDiagnosticsCallbackRequestDetail `json:"callback"`
}

func (s *equipGetDiagnosticsCallbackRequest) GetName() string {
	return services.GetDiagnostics.String()
}

type equipGetDiagnosticsCallbackRequestDetail struct {
	services.CB
	Filename *string `json:"filename"`
}

func NewEquipGetDiagnosticsCallbackRequest(sn, pod, msgId string, p *services.Protocol, status int) *equipGetDiagnosticsCallbackRequest {
	req := &equipGetDiagnosticsCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.GetDiagnostics.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: &equipGetDiagnosticsCallbackRequestDetail{
			CB: services.NewCB(status),
		},
	}
	return req
}

func NewEquipGetDiagnosticsCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *callback.CallbackError) *equipGetDiagnosticsCallbackRequest {
	req := &equipGetDiagnosticsCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.GetDiagnostics.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: &equipGetDiagnosticsCallbackRequestDetail{
			CB: services.NewCBError(err),
		},
	}
	return req
}

var _ services.Response = &equipGetDiagnosticsCallbackResponse{}

type equipGetDiagnosticsCallbackResponse struct {
	api.Response
	Data *equipGetDiagnosticsResponseDetail `json:"data"`
}

type equipGetDiagnosticsResponseDetail struct {
}

func (resp *equipGetDiagnosticsCallbackResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipGetDiagnosticsCallbackResponse) GetMsg() string {
	return resp.Msg
}

func GetDiagnosticsCallbackRequestG(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.GetDiagnostics)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipGetDiagnosticsCallbackResponse{})
}
