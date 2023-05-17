package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type equipSetLoadBalanceRequest struct {
	services.Base
	Callback services.CB                       `json:"callback"`
	Data     *equipSetLoadBalanceRequestDetail `json:"data"`
}

type equipSetLoadBalanceRequestDetail struct {
}

func (*equipSetLoadBalanceRequest) GetName() string {
	return services.SetLoadBalance.String()
}

func NewEquipSetLoadBalanceRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipSetLoadBalanceRequest {
	req := &equipSetLoadBalanceRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.SetLoadBalance.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
		Data:     &equipSetLoadBalanceRequestDetail{},
	}
	return req
}

func NewEquipSetLoadBalanceRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipSetLoadBalanceRequest {
	req := &equipSetLoadBalanceRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.SetLoadBalance.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
		Data:     &equipSetLoadBalanceRequestDetail{},
	}
	return req
}

var _ services.Response = &equipSetLoadBalanceResponse{}

type equipSetLoadBalanceResponse struct {
	api.Response
	Data *equipSetLoadBalanceResponseDetail `json:"data"`
}

func (e *equipSetLoadBalanceResponse) GetStatus() int {
	return e.Status
}

func (e *equipSetLoadBalanceResponse) GetMsg() string {
	return e.Msg
}

type equipSetLoadBalanceResponseDetail struct {
}

func SetLoadBalanceRequestWithG(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.SetLoadBalance)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipSetLoadBalanceResponse{})
}
