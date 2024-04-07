package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

var _ services.Request = &equipNotifyEventRequest{}

type equipNotifyEventRequest struct {
	services.Base
	Data *equipNotifyEventRequestData `json:"data"`
}

func (r *equipNotifyEventRequest) GetName() services.Request2ServicesNameType {
	return services.NotifyEvent
}

func (equipNotifyEventRequest) IsCallback() bool {
	return false
}

type equipNotifyEventRequestData struct {
	Code          int64  `json:"code"`
	Time          int64  `json:"time"`
	Clean         bool   `json:"clean"`
	EventID       int64  `json:"eventId"`
	RemoteAddress string `json:"remoteAddress"`
	ConnectorId   string `json:"connectorSerial"`
}

type equipNotifyEventResponse struct {
	api.Response
}

func (r *equipNotifyEventResponse) GetStatus() int {
	return r.Status
}

func (r *equipNotifyEventResponse) GetMsg() string {
	return r.Msg
}

func NewNotifyEventRequest(sn, pod, msgID string, p *services.Protocol, code, time int64, clean bool, eventID int64, remoteAddress, connectorId string) *equipNotifyEventRequest {
	return &equipNotifyEventRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.NotifyEvent.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipNotifyEventRequestData{
			Code:          code,
			Time:          time,
			Clean:         clean,
			EventID:       eventID,
			RemoteAddress: remoteAddress,
			ConnectorId:   connectorId,
		},
	}
}

func NotifyEventRequest(ctx context.Context, req *equipNotifyEventRequest) error {
	header := services.GetSimpleHeaderValue(services.NotifyEvent)

	url := services.GetSimpleURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipNotifyEventResponse{})
}
