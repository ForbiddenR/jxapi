package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi/v2"
	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipSetChargeTemplateCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (r *equipSetChargeTemplateCallbackRequest) GetName() services.Request2ServicesNameType {
	return services.SetChargeTemplate
}

func (r *equipSetChargeTemplateCallbackRequest) TraceId() string {
	return r.MsgID
}

func (r *equipSetChargeTemplateCallbackRequest) IsCallback() bool {
	return true
}

func NewEquipSetChargeTemplateCallbackRequest(sn, pod, msgId string, p *services.Protocol, status int) *equipSetChargeTemplateCallbackRequest {
	req := &equipSetChargeTemplateCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.SetChargeTemplate.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipSetChargeTemplateCallbackRequestError(sn, pod, msgId string, p *services.Protocol, status int, err *apierrors.CallbackError) *equipSetChargeTemplateCallbackRequest {
	req := &equipSetChargeTemplateCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.SetChargeTemplate.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCBError(err),
	}
	return req
}

type equipSetChargeTemplateCallbackResponse struct {
	api.Response
}

func (r *equipSetChargeTemplateCallbackResponse) GetStatus() int {
	return r.Status
}

func (r *equipSetChargeTemplateCallbackResponse) GetMsg() string {
	return r.Msg
}

func SetChargeTemplateCallbackRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
