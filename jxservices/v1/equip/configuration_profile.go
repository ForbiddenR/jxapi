package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi/v2"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipConfigurationProfileRequest struct {
	services.Base
}

func (r *equipConfigurationProfileRequest) GetName() services.Request2ServicesNameType {
	return services.ConfigurationProfile
}

func (r *equipConfigurationProfileRequest) TraceId() string {
	return r.MsgID
}

func (r *equipConfigurationProfileRequest) IsCallback() bool {
	return false
}

func NewEquipConfigurationProfileRequest(sn, pod, msgId string, p *services.Protocol) *equipConfigurationProfileRequest {
	return &equipConfigurationProfileRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.ConfigurationProfile.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
	}
}

type equipConfigurationProfileResponse struct {
	api.Response
}

func (r *equipConfigurationProfileResponse) GetMsg() string {
	return r.Msg
}

func (r *equipConfigurationProfileResponse) GetStatus() int {
	return r.Status
}

func ConfigurationProfileRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
