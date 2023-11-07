package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type equipOnlineRequest struct {
	services.Base
	Data *equipOnlineRequestDetail `json:"data"`
}

type equipOnlineRequestDetail struct {
	RemoteAddress *string `json:"remoteAddress"`
}

type OnlineConfig struct {
	services.ReusedConfig
}

func NewEquipOnlineRequestWithConfig(config OnlineConfig) *equipOnlineRequest {
	return NewEquipOnlineRequest(config.Sn, config.Protocol, config.Pod, config.MsgID)
}

func NewEquipOnlineRequest(sn string, protocol *services.Protocol, pod, msgID string) *equipOnlineRequest {
	return &equipOnlineRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    protocol,
			Category:    services.Online.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},

		Data: &equipOnlineRequestDetail{},
	}
}

func (o *equipOnlineRequest) GetName() string {
	return services.Online.String()
}

var _ services.Response = &equipOnlineResponse{}

type equipOnlineResponse struct {
	api.Response
	Data *equipOnlineResponseDetail `json:"data"`
}

type equipOnlineResponseDetail struct {
	EquipmentID string `json:"equipmentId" validate:"required"`
	EquipmentSN string `json:"equipmentSN" validate:"required"`
}

func (resp *equipOnlineResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipOnlineResponse) GetMsg() string {
	return resp.Msg
}

func OnlineRequest(ctx context.Context, req *equipOnlineRequest) error {
	header := services.GetSimpleHeaderValue(services.Online)

	url := services.GetSimpleURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipOnlineResponse{})
}
