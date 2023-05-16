package v1

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"gitee.com/csms/jxeu-ocpp/internal/config"
	"gitee.com/csms/jxeu-ocpp/pkg/api"
	"gitee.com/csms/jxeu-ocpp/pkg/api/esam"
	"gitee.com/csms/jxeu-ocpp/pkg/utils"
)

const readWait = 80 * time.Second

type accessVerifyRequest struct {
	EquipmentSn      string  `json:"equipmentSn"`
	Protocol         string  `json:"deviceProtocol,omitempty"`
	ProtocolVersion  string  `json:"deviceProtocolVersion,omitempty"`
	RequestPort      string  `json:"requestPort"`
	RemoteAddress    *string `json:"remoteAddress"`
	CertSerialNumber *string `json:"certSerialNumber"`
	AccountPassword  *string `json:"accountPassword"`
}

func NewAccessVerifyRequest(sn, requestPort, protocol, protocolVersion string) *accessVerifyRequest {
	return &accessVerifyRequest{
		EquipmentSn:     sn,
		Protocol:        protocol,
		ProtocolVersion: protocolVersion,
		RequestPort:     requestPort,
	}
}

func (a *accessVerifyRequest) getName() string {
	return "accessVerify"
}

type accessVerifyResponse struct {
	api.Response
	Data *accessVerifyResponseData `json:"data"`
}

type accessVerifyResponseData struct {
	ID               string         `json:"id"`
	BaseUrl          string         `json:"baseUrl"`
	HearbeatInterval int            `json:"keepaliveInterval"`
	Registered       bool           `json:"registered"`
	ReadWait         utils.Duration `json:"readWait"`
}

func AccessVerifyRequest(ctx context.Context, ticket string, request *accessVerifyRequest) (*accessVerifyResponse, error) {
	headerValue := make([]string, 0)
	headerValue = append(headerValue, api.Esam, esam.Equip)
	headerValue = append(headerValue, esam.Access.Split()...)

	header := map[string]string{api.Perms: strings.Join(headerValue, ":"), esam.TicketKey: ticket}
	//header := map[string]string{"Perms": "esam:equip:access:verify"}
	url := config.App.EsamUrl + esam.Equip + "/verify"
	resp, err := api.SendRequest(ctx, url, request, header)
	if err != nil {
		return nil, err
	}
	accessResponse := new(accessVerifyResponse)
	err = json.Unmarshal(resp, accessResponse)
	if err != nil {
		return nil, err
	}
	if accessResponse.Status == 1 {
		return accessResponse, errors.New(accessResponse.Msg)
	}
	if accessResponse.Data == nil {
		return accessResponse, errors.New("data is empty")
	}

	if accessResponse.Data.BaseUrl == "" {
		return accessResponse, errors.New("baseUrl is empty")
	}

	if accessResponse.Data.ReadWait == 0 {
		accessResponse.Data.ReadWait = utils.Duration(readWait)
	}

	return accessResponse, nil
}
