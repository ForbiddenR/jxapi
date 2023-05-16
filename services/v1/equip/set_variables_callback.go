package equip

import (
	"context"

	api "github.com/ForbiddenR/jx-api"
	"github.com/ForbiddenR/jx-api/apierrors"
	"github.com/ForbiddenR/jx-api/services"
)

const (
	SetVariablesAccept         = 0
	SetVariablesRejected       = 1
	SetVariablesRebootRequired = 2
	SetVariablesNotSupported   = 3
)

// func OCPP16SetVariablesStatus(status protocol.ChangeConfigurationResponseJsonStatus) int {
// 	switch status {
// 	case protocol.ChangeConfigurationResponseJsonStatusAccepted:
// 		return SetVariablesAccept
// 	case protocol.ChangeConfigurationResponseJsonStatusRejected:
// 		return SetVariablesRejected
// 	case protocol.ChangeConfigurationResponseJsonStatusRebootRequired:
// 		return SetVariablesRebootRequired
// 	default:
// 		return SetVariablesNotSupported
// 	}
// }

type equipSetVariablesCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (s *equipSetVariablesCallbackRequest) GetName() string {
	return services.ChangeConfiguration.String()
}

func NewEquipSetVariablesCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipSetVariablesCallbackRequest {
	req := &equipSetVariablesCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.ChangeConfiguration.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipSetVariablesRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipSetVariablesCallbackRequest {
	req := &equipSetVariablesCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.ChangeConfiguration.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
	}
	return req
}

var _ services.Response = &equipSetVariablesCallbackResponse{}

type equipSetVariablesCallbackResponse struct {
	api.Response
}

func (resp *equipSetVariablesCallbackResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipSetVariablesCallbackResponse) GetMsg() string {
	return resp.Msg
}

func SetVariablesRequestWithGeneric(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.ChangeConfiguration)

	url := services.GetCallbackURL(req)
	return services.RequestWithoutResponse(ctx, req, url, header, &equipSetVariablesCallbackResponse{})
}
