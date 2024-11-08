package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi/v2"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipFetchChargeTemplateRequest struct {
	services.Base
	Data *equipFetchChargeTemplateRequestDetail `json:"data"`
}

func (r *equipFetchChargeTemplateRequest) GetName() services.Request2ServicesNameType {
	return services.FetchChargeTemplate
}

func (r *equipFetchChargeTemplateRequest) TraceId() string {
	return r.MsgID
}

func (r *equipFetchChargeTemplateRequest) IsCallback() bool {
	return true
}

type equipFetchChargeTemplateRequestDetail struct {
	TemplateId string `json:"template_id"`
}

func NewequipFetchChargeTemplateRequest(sn, pod, msgId string, p *services.Protocol) *equipFetchChargeTemplateRequest {
	req := &equipFetchChargeTemplateRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.FetchChargeTemplate.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Data: &equipFetchChargeTemplateRequestDetail{},
	}
	return req
}

type equipFetchChargeTemplateResponse struct {
	api.Response
	Data *equipFetchChargeTemplateResponseData `json:"data"`
}

type equipFetchChargeTemplateResponseData struct {
	PriceScheme
}

func (r *equipFetchChargeTemplateResponse) GetStatus() int {
	return r.Status
}

func (r *equipFetchChargeTemplateResponse) GetMsg() string {
	return r.Msg
}

func FetchChargeTemplateRequest(ctx context.Context, req services.Request) (*equipFetchChargeTemplateResponse, error) {
	resp := &equipFetchChargeTemplateResponse{}
	err := services.TransportWithResp(ctx, req, resp)
	return resp, err
}
