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

type equipRegisterRequest struct {
	services.Base
	Data *equipRegisterRequestDetail `json:"data"`
}

type equipRegisterRequestDetail struct {
	RemoteAddress *string `json:"remoteAddress"`
}

func NewEquipRegisterRequest(sn string, protocol *services.Protocol, pod, msgID string) *equipRegisterRequest {
	return &equipRegisterRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    protocol,
			Category:    services.Register.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},

		Data: &equipRegisterRequestDetail{},
	}
}

func (r *equipRegisterRequest) GetName() string {
	return services.Register.String()
}

var _ services.Response = &equipRegisterResponse{}

type equipRegisterResponse struct {
	api.Response
	Data *equipRegisterResponseDetail `json:"data"`
}

type equipRegisterResponseDetail struct {
	EquipmentID string `json:"equipmentId" validate:"required"`
	EquipmentSN string `json:"equipmentSN" validate:"required"`
}

func (resp *equipRegisterResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipRegisterResponse) GetMsg() string {
	return resp.Msg
}

func RegisterRequest(ctx context.Context, req *equipRegisterRequest) error {
	headerValue := make([]string, 0)
	headerValue = append(headerValue, api.Services, services.Equip)
	headerValue = append(headerValue, services.Register.Split()...)
	header := map[string]string{api.Perms: strings.Join(headerValue, ":")}

	url := config.App.ServicesUrl + services.Equip + "/" + req.GetName()

	message, err := api.SendRequest(ctx, url, req, header)
	if err != nil {
		return err
	}

	resp := &equipRegisterResponse{}
	err = json.Unmarshal(message, resp)
	if err != nil {
		return err
	}

	if resp.Status == 1 {
		return errors.New(resp.Msg)
	}

	return nil
}

func RegisterRequestWithGeneric(ctx context.Context, req *equipRegisterRequest) (id string, err error) {
	header := services.GetSimpleHeaderValue(services.Register)

	url := services.GetSimpleURL(req)

	resp := equipRegisterResponse{}
	err = services.RequestWithoutResponse(ctx, req, url, header, &resp)
	if err != nil {
		return
	}

	if resp.Data == nil {
		err = errors.New("response data is nil")
		return
	}

	id = resp.Data.EquipmentID
	return
}
