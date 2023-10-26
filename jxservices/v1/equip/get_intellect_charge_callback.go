package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type equipGetIntellectChargeRequest struct {
	services.Base
	Callback *equipGetIntellectChargeRequestDetail `json:"callback"`
}

type equipGetIntellectChargeRequestDetail struct {
	services.CB
	EVSE           EVSE     `json:"evse"`
	IntellectType  uint8    `json:"type"`
	IntellectId    string   `json:"intellectId"`
	StartTime      int64    `json:"startTime"`
	EndTime        *int64   `json:"endTime"`
	EndElectricity *float64 `json:"endElectricity"`
	EndSoc         *float64 `json:"endSOC"`
	Status         uint8    `json:"status"`
}

func (*equipGetIntellectChargeRequest) GetName() string {
	return services.GetIntellectCharge.String()
}

func NewEquipGetIntellectChargeCallbackRequest(sn, pod, msgID string, p *services.Protocol, status int,
	cid, evseId string, intellectType uint8, intellectId string, startTime int64, intellectStatus uint8) *equipGetIntellectChargeRequest {
	req := &equipGetIntellectChargeRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.GetIntellectCharge.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: &equipGetIntellectChargeRequestDetail{
			CB: services.NewCB(status),
			EVSE: EVSE{
				Id:          evseId,
				ConnectorId: cid,
			},
			IntellectType:  intellectType,
			IntellectId:    intellectId,
			StartTime:      startTime,
			Status:         intellectStatus,
		},
	}
	return req
}

func NewEquipGetIntellectChargeCallbackRequestError(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) *equipGetIntellectChargeRequest {
	req := &equipGetIntellectChargeRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.GetConfiguration.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: &equipGetIntellectChargeRequestDetail{
			CB: services.NewCBError(err),
		},
	}
	return req
}

var _ services.Response = &equipGetIntellectChargeResponse{}

type equipGetIntellectChargeResponse struct {
	api.Response
	Data *equipGetIntellectChargeResponseDetail `json:"data"`
}

func (resp *equipGetIntellectChargeResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipGetIntellectChargeResponse) GetMsg() string {
	return resp.Msg
}

type equipGetIntellectChargeResponseDetail struct {
}

func GetIntellectChargeCallbackRequest(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.GetIntellectCharge)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx ,req, url, header, &equipGetIntellectChargeResponse{})
} 