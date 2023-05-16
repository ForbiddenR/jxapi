package equip

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"gitee.com/csms/jxeu-ocpp/internal/config"
	callbackError "gitee.com/csms/jxeu-ocpp/internal/errors"
	"gitee.com/csms/jxeu-ocpp/pkg/api"
	"gitee.com/csms/jxeu-ocpp/pkg/api/services"
)

type equipGetVariablesCallbackRequest struct {
	services.Base
	Callback *equipGetVariablesCallbackRequestDetail `json:"callback"`
}

type equipGetVariablesCallbackRequestDetail struct {
	services.CB
	Value      VariableAttribute `json:"value"`
	UnknownKey []string          `json:"unknownKey,omitempty"`
}

func (g *equipGetVariablesCallbackRequest) GetName() string {
	return services.GetConfiguration.String()
}

func NewEquipGetVariablesCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipGetVariablesCallbackRequest {
	req := &equipGetVariablesCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.GetConfiguration.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: &equipGetVariablesCallbackRequestDetail{
			CB: services.NewCB(status),
		},
	}
	return req
}

func NewEquipGetVariablesRequestError(sn, pod, msgID string, p *services.Protocol, err *callbackError.CallbackError) *equipGetVariablesCallbackRequest {
	req := &equipGetVariablesCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.GetConfiguration.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: &equipGetVariablesCallbackRequestDetail{
			CB: services.NewCBError(err),
		},
	}
	return req
}

var _ services.Response = &equipGetVariablesCallbackResponse{}

type equipGetVariablesCallbackResponse struct {
	api.Response
	Data *equipGetVariablesCallbackResponse `json:"data"`
}

func (resp *equipGetVariablesCallbackResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipGetVariablesCallbackResponse) GetMsg() string {
	return resp.Msg
}

type equipGetVariablesResponseDetail struct {
}

func GetVariablesCallbackRequest(ctx context.Context, req services.CallbackRequest) error {
	headerValue := make([]string, 0)
	//headerValue = append(headerValue, api.Services, "get", "variables", services.Callback)
	headerValue = append(headerValue, api.Services)
	headerValue = append(headerValue, services.GetConfiguration.Split()...)
	headerValue = append(headerValue, services.Callback)

	header := map[string]string{api.Perms: strings.Join(headerValue, ":")}

	url := config.App.ServicesUrl + services.Equip + "/" + services.Callback + "/" + req.GetName() + "Callback"

	message, err := api.SendRequest(ctx, url, req, header)

	if err != nil {
		return err
	}

	resp := &equipGetVariablesCallbackResponse{}
	err = json.Unmarshal(message, resp)
	if err != nil {
		return err
	}

	if resp.Status == 1 {
		return errors.New(resp.Msg)
	}

	return nil
}

func GetVariablesCallbackRequestWithGeneric(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.GetConfiguration)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipGetVariablesCallbackResponse{})
}
