package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type equipLoginRequest struct {
	services.Base
	Data *equipLoginRequestDetail `json:"data"`
}

type equipLoginRequestDetail struct {
	RemoteAddress      *string `json:"remoteAddress"`
	ModelCode          string  `json:"modelCode"`
	ManufacturerCode   string  `json:"manufacturerCode"`
	FirmwareVersion    *string `json:"firmwareVersion"`
	Iccid              *string `json:"iccid"`
	Imsi               *string `json:"imsi"`
	ReconnectingReason *string `json:"reason"`
}

func (equipLoginRequest) GetName() string {
	return services.Login.String()
}

type LoginRequestConfig struct {
	services.ReusedConfig
	ModelCode        string
	ManufacturerCode string
}

func NewLogin(base services.Base, modelCode, manufacturerCode string) *equipLoginRequest {
	req := &equipLoginRequest{
		Base: base,
		Data: &equipLoginRequestDetail{
			ModelCode:          modelCode,
			ManufacturerCode:   manufacturerCode,
		},
	}
	return req
}

func NewEquipLoginRequestWithConfig(config *LoginRequestConfig) *equipLoginRequest {
	return NewEquipLoginRequest(config.Sn, config.Pod, config.MsgID, config.Protocol,
		config.ModelCode, config.ManufacturerCode)
}

func NewEquipLoginRequest(sn, pod, msgID string, p *services.Protocol,
	modelCode, manufacturerCode string) *equipLoginRequest {
	request := &equipLoginRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.Login.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
	}
	request.Data = &equipLoginRequestDetail{
		ModelCode:        modelCode,
		ManufacturerCode: manufacturerCode,
	}
	return request
}

type equipLoginResponse struct {
	api.Response
}

func (resp *equipLoginResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipLoginResponse) GetMsg() string {
	return resp.Msg
}

func LoginRequest(ctx context.Context, req services.Request) error {
	header := services.GetSimpleHeaderValue(services.Login)

	url := services.GetSimpleURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipLoginResponse{})
}
