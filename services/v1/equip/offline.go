package equip

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"gitee.com/csms/jxeu-ocpp/internal/config"
	"gitee.com/csms/jxeu-ocpp/pkg/api"
	"gitee.com/csms/jxeu-ocpp/pkg/api/services"
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

func OfflineRequest(ctx context.Context, req *equipOfflineRequest) error {
	headerValue := make([]string, 0)
	headerValue = append(headerValue, api.Services, services.Equip)
	headerValue = append(headerValue, services.Offline.Split()...)

	header := map[string]string{api.Perms: strings.Join(headerValue, ":")}

	url := config.App.ServicesUrl + services.Equip + "/" + req.GetName()

	message, err := api.SendRequest(ctx, url, req, header)
	if err != nil {
		return err
	}

	resp := &equipOfflineResponse{}

	err = json.Unmarshal(message, resp)

	if err != nil {
		return err
	}

	if resp.Status == 1 {
		return errors.New(resp.Msg)
	}

	return nil
}

func OfflineRequestWithGeneric(ctx context.Context, req *equipOfflineRequest) error {
	header := services.GetSimpleHeaderValue(services.Offline)

	url := services.GetSimpleURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipOfflineResponse{})
}
