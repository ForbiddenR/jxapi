package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type equipSetIntellectChargeRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (*equipSetIntellectChargeRequest) GetName() string {
	return services.SetIntellectCharge.String()
}

func NewEquipSetIntellectChargeCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipSetIntellectChargeRequest {
	req := &equipSetIntellectChargeRequest{
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

func NewEquipSetIntellectChargeCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipSetIntellectChargeRequest {
	req := &equipSetIntellectChargeRequest{
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

var _ services.Response = &equipSetIntellectChargeResponse{}

type equipSetIntellectChargeResponse struct {
	api.Response
}

func (resp *equipSetIntellectChargeResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipSetIntellectChargeResponse) GetMsg() string {
	return resp.Msg
}

func SetIntellectChargeRequest(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.SendQRCode)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipSetIntellectChargeResponse{})
}
