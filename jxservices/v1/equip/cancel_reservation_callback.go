package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type equipCancelReservationCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (r *equipCancelReservationCallbackRequest) GetName() string {
	return services.CancelReservation.String()
}

func NewEquipCancelReseravtionCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipCancelReservationCallbackRequest {
	req := &equipCancelReservationCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.CancelReservation.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipCancelReservationCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipCancelReservationCallbackRequest {
	req := &equipCancelReservationCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.CancelReservation.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
	}
	return req
}

var _ services.Response = &equipCancelReservationCallbackResponse{}

type equipCancelReservationCallbackResponse struct {
	api.Response
}

func (resp *equipCancelReservationCallbackResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipCancelReservationCallbackResponse) GetMsg() string {
	return resp.Msg
}

func CancelReservationCallbackRequestWithGeneric(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.CancelReservation)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipCancelReservationCallbackResponse{})
}
