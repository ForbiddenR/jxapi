package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
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

// func OCPP16ConfigurationKeyToVariable(keys []ocpp16.GetConfigurationResponseJsonConfigurationKeyElem) []equipGetBaseReportCallbackRequestDetailVariable {
// 	variable := make([]equipGetBaseReportCallbackRequestDetailVariable, 0)
// 	for _, key := range keys {
// 		v := NewEquipGetBaseReportCallbackRequestDetailVariable(key.Key, key.Readonly)
// 		if key.Value != nil {
// 			v.Value = key.Value
// 		}
// 		variable = append(variable, v)
// 	}
// 	return variable
// }

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

func NewEquipGetBaseReportRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipGetBaseReportCallbackRequest {
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

func GetBaseReportCallbackRequestWithGeneric(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.GetConfiguration)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipGetBaseReportCallbackResponse{})
}
