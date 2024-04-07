package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

var _ services.Request = &equipUpdateFirmwareCallbackRequest{}

type equipUpdateFirmwareCallbackRequest struct {
	services.Base
	Callback services.CB                               `json:"callback"`
	Data     *equipUpdateFirmwareCallbackRequestDetail `json:"data"`
}

type equipUpdateFirmwareCallbackRequestDetail struct {
}

func (*equipUpdateFirmwareCallbackRequest) GetName() services.Request2ServicesNameType {
	return services.UpdateFirmware
}

func (equipUpdateFirmwareCallbackRequest) IsCallback() bool {
	return true
}

func NewEquipUpdateFirmwareCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipUpdateFirmwareCallbackRequest {
	return &equipUpdateFirmwareCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.UpdateFirmware.FirstUpper(),
			AccessPod:   pod,

			MsgID: msgID,
		},
		Callback: services.NewCB(status),
	}
}

func NewEquipUpdateFirmwareCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipUpdateFirmwareCallbackRequest {
	return &equipUpdateFirmwareCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.UpdateFirmware.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
		Data:     nil,
	}
}

var _ services.Response = &equipUpdateFirmwareCallbackResponse{}

type equipUpdateFirmwareCallbackResponse struct {
	api.Response
}

func (resp *equipUpdateFirmwareCallbackResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipUpdateFirmwareCallbackResponse) GetMsg() string {
	return resp.Msg
}

func UpdateFirmwareCallbackRequestWithGeneric(ctx context.Context, req services.Request) error {
	header := services.GetCallbackHeaderValue(services.UpdateFirmware)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipUpdateFirmwareCallbackResponse{})
}
