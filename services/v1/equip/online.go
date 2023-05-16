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

type equipOnlineRequest struct {
	services.Base
	Data *equipOnlineRequestDetail `json:"data"`
}

type equipOnlineRequestDetail struct {
	RemoteAddress *string `json:"remoteAddress"`
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
	headerValue := make([]string, 0)
	headerValue = append(headerValue, api.Services, services.Equip)
	headerValue = append(headerValue, services.Online.Split()...)
	header := map[string]string{api.Perms: strings.Join(headerValue, ":")}

	url := config.App.ServicesUrl + services.Equip + "/" + req.GetName()

	message, err := api.SendRequest(ctx, url, req, header)
	if err != nil {
		return err
	}

	resp := &equipOnlineResponse{}
	err = json.Unmarshal(message, resp)
	if err != nil {
		return err
	}

	if resp.Status == 1 {
		return errors.New(resp.Msg)
	}

	return nil
}

func OnlineRequestWithGeneric(ctx context.Context, req *equipOnlineRequest) (id string, err error) {
	header := services.GetSimpleHeaderValue(services.Online)

	url := services.GetSimpleURL(req)

	resp := equipOnlineResponse{}
	err = services.RequestWithoutResponse(ctx, req, url, header, &resp)
	if err != nil {
		return "", err
	}
	if resp.Data == nil {
		return "", errors.New("response data is nil")
	}

	return resp.Data.EquipmentID, nil
}
