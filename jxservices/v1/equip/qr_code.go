package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type equipQRCodeRequest struct {
	services.Base
	Data *equipQRCoeRequestDetail `json:"data"`
}

type equipQRCoeRequestDetail struct {
}

func NewEquipQRCodeRequest(sn, pod, msgID string, protocol *services.Protocol) *equipQRCodeRequest {
	return &equipQRCodeRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    protocol,
			Category:    services.QRCode.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipQRCoeRequestDetail{},
	}
}

func (q *equipQRCodeRequest) GetName() string {
	return services.QRCode.String()
}

var _ services.Response = &equipQRCodeResponse{}

type equipQRCodeResponse struct {
	api.Response
	Data *equipQRCodeResponseDetail `json:"data"`
}

type equipQRCodeResponseDetail struct {
	Qrcode string `json:"qrcode" validate:"required"`
}

func (q *equipQRCodeResponse) GetStatus() int {
	return q.Status
}

func (q *equipQRCodeResponse) GetMsg() string {
	return q.Msg
}

func QRCode(ctx context.Context, req services.Request) (*equipQRCodeResponse, error) {
	header := services.GetSimpleHeaderValue(services.QRCode)

	url := services.GetSimpleURL(req)

	return services.RequestWithResponse(ctx, req, url, header, &equipQRCodeResponse{})
}
