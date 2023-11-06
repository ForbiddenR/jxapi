package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type equipReserveNowCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (r *equipReserveNowCallbackRequest) GetName() string {
	return services.ReserveNow.String()
}

func NewEquipReserveNowCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipReserveNowCallbackRequest {
	req := &equipReserveNowCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.ReserveNow.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipReserveNowCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipReserveNowCallbackRequest {
	req := &equipReserveNowCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.ReserveNow.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
	}
	return req
}

type equipReserveNowCallbackResponse struct {
	api.Response
}

func (resp *equipReserveNowCallbackResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipReserveNowCallbackResponse) GetMsg() string {
	return resp.Msg
}

func ReserveNowCallbackRequest(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.ReserveNow)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipReserveNowCallbackResponse{})
}
