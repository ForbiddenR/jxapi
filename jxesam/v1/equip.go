package v1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	api "github.com/ForbiddenR/jxapi/v2"
	esam "github.com/ForbiddenR/jxapi/v2/jxesam"
	utils "github.com/ForbiddenR/jxapi/v2/jxutils"
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

type accessVerifyResponse struct {
	api.Response
	Data *accessVerifyResponseData `json:"data"`
}

type accessVerifyResponseData struct {
	ID                  string         `json:"id"`
	BaseUrl             string         `json:"baseUrl"`
	HearbeatInterval    int            `json:"keepalive"`
	TransactionInterval int            `json:"transactionInterval"`
	Registered          bool           `json:"registered"`
	ReadWait            utils.Duration `json:"readWait"`
	Blocked             bool           `json:"blocked"`
}

func AccessVerifyRequest(ctx context.Context, ticket string, traceId string, request *accessVerifyRequest) (*accessVerifyResponse, error) {
	headerValue := make([]string, 0)
	headerValue = append(headerValue, api.Esam, esam.Equip)
	headerValue = append(headerValue, esam.Access.Split()...)

	header := map[string]string{api.Perms: strings.Join(headerValue, ":"), esam.TicketKey: ticket, "TraceId": traceId}
	//header := map[string]string{"Perms": "esam:equip:access:verify"}
	url := api.EsamUrl + esam.Equip + "/verify"
	resp, err := api.SendRequest(ctx, url, request, header)
	if err != nil {
		return nil, err
	}
	accessResponse := new(accessVerifyResponse)
	err = json.Unmarshal(resp, accessResponse)
	if err != nil {
		return nil, fmt.Errorf("bad json format: %s", resp)
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
