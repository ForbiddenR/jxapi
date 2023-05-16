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

type equipAuthorizeTransactionRequest struct {
	services.Base
	Data *equipAuthorizeTransactionRequestDetail `json:"data"`
}

type equipAuthorizeTransactionRequestDetail struct {
	IdTokenType IdTokenType `json:"idTokenType"`
}

func (r equipAuthorizeTransactionRequest) GetName() string {
	return services.Authorize.String()
}

func NewEquipAuthorizeTransactionRequest(sn, pod, msgID string, p *services.Protocol) *equipAuthorizeTransactionRequest {
	return &equipAuthorizeTransactionRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.Authorize.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipAuthorizeTransactionRequestDetail{},
	}
}

var _ services.Response = &equipAuthorizeTransactionResponse{}

type equipAuthorizeTransactionResponse struct {
	api.Response
	Data *equipAuthorizeTransactionResponseDetail `json:"data"`
}

func (resp *equipAuthorizeTransactionResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipAuthorizeTransactionResponse) GetMsg() string {
	return resp.Msg
}

type equipAuthorizeTransactionResponseDetail struct {
	IdTokenInfo IdTokenInfo `json:"idTokenInfo" validate:"required"`
}

// AuthorizeTransactionRequest needs to return the necessary response.
func AuthorizeTransactionRequest(ctx context.Context, req *equipAuthorizeTransactionRequest) (*equipAuthorizeTransactionResponse, error) {
	headerValue := make([]string, 0)
	headerValue = append(headerValue, api.Services, services.Equip)
	headerValue = append(headerValue, services.Authorize.Split()...)
	header := map[string]string{api.Perms: strings.Join(headerValue, ":")}

	url := config.App.ServicesUrl + services.Equip + "/" + req.GetName()

	message, err := api.SendRequest(ctx, url, req, header)

	if err != nil {
		return nil, err
	}

	resp := &equipAuthorizeTransactionResponse{}

	err = json.Unmarshal(message, resp)

	if err != nil {
		return nil, err
	}

	if resp.Status == 1 {
		return nil, errors.New(resp.Msg)
	}

	return resp, nil
}

func AuthorizeTransactionRequestWithGeneric(ctx context.Context, req *equipAuthorizeTransactionRequest) (*equipAuthorizeTransactionResponse, error) {
	header := services.GetSimpleHeaderValue(services.Authorize)

	url := services.GetSimpleURL(req)

	return services.RequestWithResponse(ctx, req, url, header, &equipAuthorizeTransactionResponse{})
}
