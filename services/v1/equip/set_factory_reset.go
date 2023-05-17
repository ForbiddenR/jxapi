package equip

import (
	"context"
	
	api "github.com/ForbiddenR/jx-api"
	"github.com/ForbiddenR/jx-api/apierrors"
	"github.com/ForbiddenR/jx-api/services"
)

type equipSetFactoryResetRequest struct {
	services.Base
	Callback services.CB                        `json:"callback"`
	Data     *equipSetFactoryResetRequestDetail `json:"data"`
}

type equipSetFactoryResetRequestDetail struct {
}

func (*equipSetFactoryResetRequest) GetName() string {
	return services.SetFactoryReset.String()
}

func NewEquipSetFactoryResetRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipSetFactoryResetRequest {
	req := &equipSetFactoryResetRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.SetFactoryReset.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
		Data:     &equipSetFactoryResetRequestDetail{},
	}
	return req
}

func NewEquipSetFactoryResetRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipSetFactoryResetRequest {
	req := &equipSetFactoryResetRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.SetFactoryReset.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
		Data:     &equipSetFactoryResetRequestDetail{},
	}
	return req
}

type equipSetFactoryResetResponse struct {
	api.Response
	Data *equipSetFactoryResetResponseDetail `json:"data"`
}

var _ services.Response = &equipSetFactoryResetResponse{}

func (e *equipSetFactoryResetResponse) GetStatus() int {
	return e.Status
}

func (e *equipSetFactoryResetResponse) GetMsg() string {
	return e.Msg
}

type equipSetFactoryResetResponseDetail struct {
}

func SetFactoryResetRequestWithG(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.SetFactoryReset)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipSetFactoryResetResponse{})
}
