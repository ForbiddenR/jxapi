package equip

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	api "github.com/ForbiddenR/jx-api"
	"github.com/ForbiddenR/jx-api/apierrors"
	"github.com/ForbiddenR/jx-api/services"
)

type equipNotifyEventRequest struct {
	services.Base
	Data *equipNotifyEventRequestData `json:"data"`
}

func (r *equipNotifyEventRequest) GetName() string {
	return services.NotifyEvent.String()
}

type equipNotifyEventRequestData struct {
	Code          int64  `json:"code"`
	Time          int64  `json:"time"`
	Clean         bool   `json:"clean"`
	EventID       int64  `json:"eventId"`
	RemoteAddress string `json:"remoteAddress"`
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

func NewNotifyEventRequest(sn, pod, msgID string, p *services.Protocol, code int64, time int64, clean bool, eventID int64, remoteAddress string) *equipNotifyEventRequest {
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
		},
	}
}

func NotifyEventRequest(ctx context.Context, req *equipNotifyEventRequest) error {
	headerValue := make([]string, 0)
	headerValue = append(headerValue, api.Services, services.Equip)
	headerValue = append(headerValue, services.NotifyEvent.Split()...)

	header := map[string]string{api.Perms: strings.Join(headerValue, ":")}

	url := config.App.ServicesUrl + services.Equip + "/" + req.GetName()

	message, err := api.SendRequest(ctx, url, req, header)
	if err != nil {
		return err
	}

	resp := &equipNotifyEventResponse{}

	err = json.Unmarshal(message, resp)

	if err != nil {
		return err
	}

	if resp.Status == 1 {
		return errors.New(resp.Msg)
	}

	return nil

}
