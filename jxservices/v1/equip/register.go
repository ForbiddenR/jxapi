package equip

import (
	"context"
	"errors"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
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
