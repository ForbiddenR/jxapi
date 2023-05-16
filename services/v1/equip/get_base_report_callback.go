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
	ocpp16 "gitee.com/csms/jxeu-ocpp/pkg/ocpp1.6/protocol"
)

type equipGetBaseReportCallbackRequest struct {
	services.Base
	Callback *equipGetBaseReportCallbackRequestDetail `json:"callback"`
}

type equipGetBaseReportCallbackRequestDetail struct {
	services.CB
	Variable   []equipGetBaseReportCallbackRequestDetailVariable `json:"variable"`
	UnknownKey []string                                          `json:"unknownKey,omitempty"`
}

type equipGetBaseReportCallbackRequestDetailVariable struct {
	Key      string  `json:"key"`
	Readonly bool    `json:"readonly"`
	Value    *string `json:"value"`
}

func NewEquipGetBaseReportCallbackRequestDetailVariable(key string, readonly bool) equipGetBaseReportCallbackRequestDetailVariable {
	return equipGetBaseReportCallbackRequestDetailVariable{
		Key:      key,
		Readonly: readonly,
	}
}

func OCPP16ConfigurationKeyToVariable(keys []ocpp16.GetConfigurationResponseJsonConfigurationKeyElem) []equipGetBaseReportCallbackRequestDetailVariable {
	variable := make([]equipGetBaseReportCallbackRequestDetailVariable, 0)
	for _, key := range keys {
		v := NewEquipGetBaseReportCallbackRequestDetailVariable(key.Key, key.Readonly)
		if key.Value != nil {
			v.Value = key.Value
		}
		variable = append(variable, v)
	}
	return variable
}

func (g *equipGetBaseReportCallbackRequest) GetName() string {
	return services.GetBaseReport.String()
}

func NewEquipGetBaseReportCallbackRequestOCPP16(sn, pod, msgID string, status int, variable []equipGetBaseReportCallbackRequestDetailVariable, unknownKey []string) *equipGetBaseReportCallbackRequest {
	req := &equipGetBaseReportCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    services.OCPP16(),
			Category:    services.GetBaseReport.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: &equipGetBaseReportCallbackRequestDetail{
			CB: services.NewCB(status),
		},
	}
	if len(variable) > 0 {
		req.Callback.Variable = variable
	}

	if len(unknownKey) > 0 {
		req.Callback.UnknownKey = unknownKey
	}

	return req
}

func NewEquipGetBaseReportCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipGetBaseReportCallbackRequest {
	req := &equipGetBaseReportCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.GetBaseReport.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: &equipGetBaseReportCallbackRequestDetail{
			CB: services.NewCB(status),
		},
	}
	return req
}

func NewEquipGetBaseReportRequestError(sn, pod, msgID string, p *services.Protocol, err *callbackError.CallbackError) *equipGetBaseReportCallbackRequest {
	req := &equipGetBaseReportCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.GetBaseReport.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: &equipGetBaseReportCallbackRequestDetail{
			CB: services.NewCBError(err),
		},
	}
	return req
}

var _ services.Response = &equipGetBaseReportCallbackResponse{}

type equipGetBaseReportCallbackResponse struct {
	api.Response
	Data *equipGetBaseReportCallbackResponse `json:"data"`
}

func (resp *equipGetBaseReportCallbackResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipGetBaseReportCallbackResponse) GetMsg() string {
	return resp.Msg
}

type equipGetBaseReportResponseDetail struct {
}

func GetBaseReportCallbackRequest(ctx context.Context, req services.CallbackRequest) error {
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

	resp := &equipGetBaseReportCallbackResponse{}
	err = json.Unmarshal(message, resp)
	if err != nil {
		return err
	}

	if resp.Status == 1 {
		return errors.New(resp.Msg)
	}

	return nil
}

func GetBaseReportCallbackRequestWithGeneric(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.GetConfiguration)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipGetBaseReportCallbackResponse{})
}
