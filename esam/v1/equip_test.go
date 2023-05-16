package v1

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"gitee.com/csms/jxeu-ocpp/internal/config"
	"gitee.com/csms/jxeu-ocpp/internal/log"
	"gitee.com/csms/jxeu-ocpp/pkg/api"
)

func TestAccessVerifyRequest(t *testing.T) {
	config.TestConfig()
	api.Init()
	log.InitNopLogger()
	requests := []*accessVerifyRequest{
		{
			EquipmentSn:      "JK000000006",
			RequestPort:      "33887",
			RemoteAddress:    nil,
			CertSerialNumber: nil,
			AccountPassword:  nil,
			Protocol:         "ocpp",
			ProtocolVersion:  "1.6",
		},
	}
	ctx := context.TODO()
	for _, request := range requests {
		resp, err := AccessVerifyRequest(ctx, "lYMFB!X#87,7woq?$C#W~z", request)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		t.Log(resp)
	}
}
