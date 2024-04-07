package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

const (
	SendLocalListAccept          = 0
	SendLocalListFailed          = 1
	SendLocalListVersionMismatch = 2
	SendLocalListNotSupported    = 3
)

// func OCPP16SendLocalListStatus(status protocol.SendLocalListResponseJsonStatus) int {
// 	switch status {
// 	case protocol.SendLocalListResponseJsonStatusAccepted:
// 		return SendLocalListAccept
// 	case protocol.SendLocalListResponseJsonStatusFailed:
// 		return SendLocalListFailed
// 	case protocol.SendLocalListResponseJsonStatusVersionMismatch:
// 		return SendLocalListVersionMismatch
// 	default:
// 		return SendLocalListNotSupported
// 	}
// }

var _ services.Request = &equipSendLocalListCallbackRequest{}

type equipSendLocalListCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (equipSendLocalListCallbackRequest) GetName() services.Request2ServicesNameType {
	return services.SendLocalList
}

func (equipSendLocalListCallbackRequest) IsCallback() bool {
	return true
}

func NewEquipSendLocalListCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipSendLocalListCallbackRequest {
	req := &equipSendLocalListCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.SendLocalList.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipSendLocalListCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipSendLocalListCallbackRequest {
	req := &equipSendLocalListCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.SendLocalList.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
	}
	return req
}

var _ services.Response = &equipSendLocalListResponse{}

type equipSendLocalListResponse struct {
	api.Response
}

func (resp *equipSendLocalListResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipSendLocalListResponse) GetMsg() string {
	return resp.Msg
}

func SendLocalListCallbackRequest(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.SendLocalList)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipSendLocalListResponse{})
}
