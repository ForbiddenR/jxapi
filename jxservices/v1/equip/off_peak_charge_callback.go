package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type equipOffPeakChargeCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (*equipOffPeakChargeCallbackRequest) GetName() string {
	return services.OffPeakCharge.String()
}

type OffPeakChargeConfig struct {
}

func NewOffPeakChargeCallbackRequest(base services.Base, config *OffPeakChargeConfig, status int) *equipOffPeakChargeCallbackRequest {
	req := &equipOffPeakChargeCallbackRequest{
		Base:     base,
		Callback: services.NewCB(status),
	}
	return req
}

func NewOffPeakChargeCallbackRequestError(base services.Base, err *apierrors.CallbackError) *equipOffPeakChargeCallbackRequest {
	return &equipOffPeakChargeCallbackRequest{
		Base:     base,
		Callback: services.NewCBError(err),
	}
}

var _ services.Response = &equipOffPeakChargeCallbackResponse{}

type equipOffPeakChargeCallbackResponse struct {
	api.Response
}

func OffPeakChargeCallbackRequest(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.OffPeakCharge)
	url := services.GetCallbackURL(req)
	return services.Transport(ctx, req, url, header)
}
