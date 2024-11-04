package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi/v2"
	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipFetchChargeTemplateCallbackRequest struct {
	services.Base
	Callback services.CB                                    `json:"callback"`
	Data     *equipFetchChargeTemplateCallbackRequestDetail `json:"data"`
}

func (r *equipFetchChargeTemplateCallbackRequest) GetName() services.Request2ServicesNameType {
	return services.FetchChargeTemplate
}

func (r *equipFetchChargeTemplateCallbackRequest) TraceId() string {
	return r.MsgID
}

func (r *equipFetchChargeTemplateCallbackRequest) IsCallback() bool {
	return true
}

type equipFetchChargeTemplateCallbackRequestDetail struct {
	TemplateId string `json:"template_id"`
}

func NewEquipFetchChargeTemplateCallbackRequest(sn, pod, msgId string, p *services.Protocol, status int) *equipFetchChargeTemplateCallbackRequest {
	req := &equipFetchChargeTemplateCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.FetchChargeTemplate.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Data:     &equipFetchChargeTemplateCallbackRequestDetail{},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipFetchChargeTemplateCallbackRequestError(sn, pod, msgId string, p *services.Protocol, err *apierrors.CallbackError) *equipFetchChargeTemplateCallbackRequest {
	req := &equipFetchChargeTemplateCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.FetchChargeTemplate.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Data:     &equipFetchChargeTemplateCallbackRequestDetail{},
		Callback: services.NewCBError(err),
	}
	return req
}

type equipFetchChargeTemplateCallbackResponse struct {
	api.Response
}

func (r *equipFetchChargeTemplateCallbackResponse) GetStatus() int {
	return r.Status
}

func (r *equipFetchChargeTemplateCallbackResponse) GetMsg() string {
	return r.Msg
}

func FetchChargeTemplateCallbackRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
