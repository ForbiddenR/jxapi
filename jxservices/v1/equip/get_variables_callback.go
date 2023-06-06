package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type equipGetVariablesCallbackRequest struct {
	services.Base
	Callback *equipGetVariablesCallbackRequestDetail `json:"callback"`
}

type equipGetVariablesCallbackRequestDetail struct {
	services.CB
	Variable   VariableAttribute `json:"variable"`
	UnknownKey *string           `json:"unknownKey,omitempty"`
}

func (g *equipGetVariablesCallbackRequest) GetName() string {
	return services.GetConfiguration.String()
}

func NewEquipGetVariablesCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipGetVariablesCallbackRequest {
	req := &equipGetVariablesCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.GetConfiguration.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: &equipGetVariablesCallbackRequestDetail{
			CB: services.NewCB(status),
		},
	}
	return req
}

func NewEquipGetVariablesRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipGetVariablesCallbackRequest {
	req := &equipGetVariablesCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.GetConfiguration.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: &equipGetVariablesCallbackRequestDetail{
			CB: services.NewCBError(err),
		},
	}
	return req
}

var _ services.Response = &equipGetVariablesCallbackResponse{}

type equipGetVariablesCallbackResponse struct {
	api.Response
	Data *equipGetVariablesCallbackResponse `json:"data"`
}

func (resp *equipGetVariablesCallbackResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipGetVariablesCallbackResponse) GetMsg() string {
	return resp.Msg
}

type equipGetVariablesResponseDetail struct {
}

func GetVariablesCallbackRequestWithGeneric(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.GetConfiguration)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipGetVariablesCallbackResponse{})
}
