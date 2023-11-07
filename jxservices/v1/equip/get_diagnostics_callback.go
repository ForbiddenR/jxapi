package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type equipGetDiagnosticsCallbackRequest struct {
	services.Base
	Callback services.CB                               `json:"callback"`
	Data     *equipGetDiagnosticsCallbackRequestDetail `json:"data"`
}

func (s *equipGetDiagnosticsCallbackRequest) GetName() string {
	return services.GetDiagnostics.String()
}

type equipGetDiagnosticsCallbackRequestDetail struct {
	// services.CB
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
		Callback: services.NewCB(status),
		Data:     &equipGetDiagnosticsCallbackRequestDetail{},
	}
	return req
}

func NewEquipGetDiagnosticsCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipGetDiagnosticsCallbackRequest {
	req := &equipGetDiagnosticsCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.GetDiagnostics.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
		Data:     &equipGetDiagnosticsCallbackRequestDetail{},
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

func GetDiagnosticsCallbackRequest(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.GetDiagnostics)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipGetDiagnosticsCallbackResponse{})
}
