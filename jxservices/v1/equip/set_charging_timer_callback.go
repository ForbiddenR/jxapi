package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

var _ services.Request = &equipSetChargingTimerCallbackRequest{}

type equipSetChargingTimerCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

//func (e *equipSetChargingTimerRequest) SetError(err *callback.CallbackError) {
//	code := err.Code()
//	msg := err.Error()
//	e.Data.Code = &code
//	e.Data.Msg = &msg
//}

func (equipSetChargingTimerCallbackRequest) GetName() services.Request2ServicesNameType {
	return services.SetChargingTimer
}

func (equipSetChargingTimerCallbackRequest) IsCallback() bool {
	return true
}

func NewEquipSetChargingTimerCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipSetChargingTimerCallbackRequest {
	req := &equipSetChargingTimerCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.SetChargingTimer.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipSetChargingTimerCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipSetChargingTimerCallbackRequest {
	req := &equipSetChargingTimerCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.SetChargingTimer.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
	}
	return req
}

var _ services.Response = &equipSetChargingTimerCallbackResponse{}

type equipSetChargingTimerCallbackResponse struct {
	api.Response
}

func (resp *equipSetChargingTimerCallbackResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipSetChargingTimerCallbackResponse) GetMsg() string {
	return resp.Msg
}

func SetChargingTimerCallbackRequest(ctx context.Context, req services.Request) error {
	header := services.GetCallbackHeaderValue(services.SetChargingTimer)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipSetChargingTimerCallbackResponse{})
}
