package equip

import (
	"context"
	"testing"

	"gitee.com/csms/jxeu-ocpp/internal/config"
	"gitee.com/csms/jxeu-ocpp/internal/errors"
	"gitee.com/csms/jxeu-ocpp/pkg/api"
	"gitee.com/csms/jxeu-ocpp/pkg/api/services"
	"gitee.com/csms/jxeu-ocpp/pkg/ocpp1.6/protocol"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestCallStatusNotificationRequest(t *testing.T) {
	config.TestConfig()
	api.Init()
	p := &services.Protocol{Name: "OCPP", Version: "1.6"}
	// logOpt := &log.Options{}
	// logOpt.Development = true
	// log.InitLogger(logOpt)
	ctx := context.TODO()

	err := errors.NewCallbackErrorGenericError(services.TestSN, protocol.RemoteStopTransactionFeatureName, "test err")
	req := []*equipCallStatusNotificationCallbackRequest{
		NewEquipCallStatusNotificationCallbackRequest(services.TestSN, services.TestAccessPod, "1", p, services.Successful),
		NewEquipCallStatusNotificationCallbackRequestError(services.TestSN, services.TestAccessPod, "1", p, err),
	}

	lo.ForEach(req, func(item *equipCallStatusNotificationCallbackRequest, _ int) {
		e := CallStatusNotificationCallbackRequestG(ctx, item)
		assert.Nil(t, e)
	})
}
