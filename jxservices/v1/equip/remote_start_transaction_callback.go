package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type equipRemoteStartTransactionCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (r *equipRemoteStartTransactionCallbackRequest) GetName() string {
	return services.RemoteStartTransaction.String()
}

func NewEquipRemoteStartTransactionCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipRemoteStartTransactionCallbackRequest {
	req := &equipRemoteStartTransactionCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.RemoteStartTransaction.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipRemoteStartTransactionCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipRemoteStartTransactionCallbackRequest {
	req := &equipRemoteStartTransactionCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.RemoteStartTransaction.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
	}
	return req
}

var _ services.Response = &equipRemoteStartTransactionCallbackResponse{}

type equipRemoteStartTransactionCallbackResponse struct {
	api.Response
}

func (resp *equipRemoteStartTransactionCallbackResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipRemoteStartTransactionCallbackResponse) GetMsg() string {
	return resp.Msg
}

func RemoteStartTransactionCallbackRequest(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.RemoteStartTransaction)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipRemoteStartTransactionCallbackResponse{})
}
