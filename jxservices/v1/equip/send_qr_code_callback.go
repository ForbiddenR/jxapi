package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

var _ services.Request = &equipSendQRCodeRequest{}

type equipSendQRCodeRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (*equipSendQRCodeRequest) GetName() services.Request2ServicesNameType {
	return services.SendQRCode
}

func (equipSendQRCodeRequest) IsCallback() bool {
	return true
}

func NewEquipSendQRCodeCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipSendQRCodeRequest {
	req := &equipSendQRCodeRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.SendQRCode.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipSendQRCodeCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipSendQRCodeRequest {
	req := &equipSendQRCodeRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.SendQRCode.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
	}
	return req
}

var _ services.Response = &equipSendQRCodeResponse{}

type equipSendQRCodeResponse struct {
	api.Response
}

func (resp *equipSendQRCodeResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipSendQRCodeResponse) GetMsg() string {
	return resp.Msg
}

func SendQRCodeRequest(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.SendQRCode)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipSendQRCodeResponse{})
}
