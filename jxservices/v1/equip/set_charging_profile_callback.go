package equip

import (
	"context"

	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

var _ services.Request = &equipSetChargingProfileRequest{}

type equipSetChargingProfileRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (equipSetChargingProfileRequest) GetName() services.Request2ServicesNameType {
	return services.SetChargingProfile
}

func (e *equipSetChargingProfileRequest) TraceId() string {
	return e.MsgID
}

func (equipSetChargingProfileRequest) IsCallback() bool {
	return true
}

func NewSetChargingProfileCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipSetChargingProfileRequest {
	req := &equipSetChargingProfileRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.SetChargingProfile.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewSetChargingProfileCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipSetChargingProfileRequest {
	req := &equipSetChargingProfileRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.SetChargingProfile.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
	}

	return req
}

// var _ services.Response = &equipSetChargingProfileResponse{}

// type equipSetChargingProfileResponse struct {
// 	api.Response
// }

// func (resp *equipSetChargingProfileResponse) GetStatus() int {
// 	return resp.Status
// }

// func (resp *equipSetChargingProfileResponse) GetMsg() string {
// 	return resp.Msg
// }

func SetChargingProfileRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
