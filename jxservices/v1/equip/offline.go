package equip

import (
	"context"
	"strings"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type OfflineReason string

const (
	// Cannot received any message in the stated period.
	Timeout = "Timeout disconnected"
	// Be closed by charger station.
	EOF = "Active disconnected"
	// Closing the ws connection.
	Initiative = "Passive disconnected"
)

func GetOfflineReason(err error) string {
	if err == nil {
		return Initiative
	}
	if strings.Contains(strings.ToLower(err.Error()), "eof") {
		return EOF
	}
	return Timeout
}

type equipOfflineRequest struct {
	services.Base
	Data *equipOfflineRequestDetail `json:"data"`
}

type equipOfflineRequestDetail struct {
	OfflineReason string `json:"offlineReason"`
}

func NewEquipOfflineRequest(sn string, protocol *services.Protocol, pod, msgID string, reason string) *equipOfflineRequest {
	return &equipOfflineRequest{
		Base: services.Base{
			Category:    services.Offline.FirstUpper(),
			EquipmentSn: sn,
			Protocol:    protocol,
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipOfflineRequestDetail{
			OfflineReason: reason,
		},
	}
}

func (e equipOfflineRequest) GetName() string {
	return services.Offline.String()
}

var _ services.Response = &equipOfflineResponse{}

type equipOfflineResponse struct {
	api.Response
}

func (resp *equipOfflineResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipOfflineResponse) GetMsg() string {
	return resp.Msg
}

func OfflineRequestWithGeneric(ctx context.Context, req *equipOfflineRequest) error {
	header := services.GetSimpleHeaderValue(services.Offline)

	url := services.GetSimpleURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipOfflineResponse{})
}