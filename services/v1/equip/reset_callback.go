package equip

import (
	"context"

	api "github.com/ForbiddenR/jx-api"
	"github.com/ForbiddenR/jx-api/apierrors"
	"github.com/ForbiddenR/jx-api/services"
)

type equipResetCallbackRequest struct {
	services.Base
	Callback services.CB                      `json:"callback"`
	Data     *equipResetCallbackRequestDetail `json:"data"`
}

type equipResetCallbackRequestDetail struct {
}

func (r *equipResetCallbackRequest) GetName() string {
	return services.Reset.String()
}

func NewEquipResetCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipResetCallbackRequest {
	req := &equipResetCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.Reset.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
	}

	return req
}

func NewEquipResetCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipResetCallbackRequest {
	req := &equipResetCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.Reset.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
	}
	return req
}

var _ services.Response = &equipResetResponse{}

type equipResetResponse struct {
	api.Response
}

func (resp *equipResetResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipResetResponse) GetMsg() string {
	return resp.Msg
}

type equipResetResponseDetail struct {
}

func ResetCallbackRequestWithGeneric(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.Reset)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipResetResponse{})
}
