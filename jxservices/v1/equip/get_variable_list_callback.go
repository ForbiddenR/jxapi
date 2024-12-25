package equip

import (
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

var _ services.Request = &equipGetVariablesCallbackRequest{}

type equipGetVariableListCallbackRequest struct {
	services.Base
	Callback services.CB                              `json:"callback"`
	Data     []equipGetVariablesCallbackRequestDetail `json:"data"`
}

func (equipGetVariableListCallbackRequest) GetName() services.Request2ServicesNameType {
	return services.GetConfiguration
}

func (e *equipGetVariableListCallbackRequest) TraceId() string {
	return e.MsgID
}

func (equipGetVariableListCallbackRequest) IsCallback() bool {
	return true
}

func NewEquipGetVariableListCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int) *equipGetVariableListCallbackRequest {
	req := &equipGetVariableListCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.GetConfiguration.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipGetVariableListRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipGetVariableListCallbackRequest {
	req := &equipGetVariableListCallbackRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.GetConfiguration.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCBError(err),
	}
	return req
}
