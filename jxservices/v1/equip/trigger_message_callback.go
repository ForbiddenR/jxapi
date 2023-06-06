package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type equipCallStatusNotificationCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (equipCallStatusNotificationCallbackRequest) GetName() string {
	return services.CallStatusNotification.String()
}

func NewEquipCallStatusNotificationCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipCallStatusNotificationCallbackRequest {
	req := &equipCallStatusNotificationCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.CallStatusNotification.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipCallStatusNotificationCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipCallStatusNotificationCallbackRequest {
	req := &equipCallStatusNotificationCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.CallStatusNotification.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
	}

	return req
}

var _ services.Response = &equipCallStatusNotificationCallbackResponse{}

type equipCallStatusNotificationCallbackResponse struct {
	api.Response
}

func (resp *equipCallStatusNotificationCallbackResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipCallStatusNotificationCallbackResponse) GetMsg() string {
	return resp.Msg
}

func CallStatusNotificationCallbackRequestG(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.CallStatusNotification)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipCallStatusNotificationCallbackResponse{})
}
