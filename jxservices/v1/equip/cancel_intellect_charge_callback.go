package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type equipCancelIntellectChargeRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (*equipCancelIntellectChargeRequest) GetName() string {
	return services.SetIntellectCharge.String()
}

func NewEquipCancelIntellectChargeCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipCancelIntellectChargeRequest {
	req := &equipCancelIntellectChargeRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.SetIntellectCharge.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipCancelIntellectChargeCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipCancelIntellectChargeRequest {
	req := &equipCancelIntellectChargeRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.SetIntellectCharge.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
	}
	return req
}

var _ services.Response = &equipCancelIntellectChargeResponse{}

type equipCancelIntellectChargeResponse struct {
	api.Response
}

func (resp *equipCancelIntellectChargeResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipCancelIntellectChargeResponse) GetMsg() string {
	return resp.Msg
}

func CancelIntellectChargeRequest(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.SendQRCode)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipCancelIntellectChargeResponse{})
}
