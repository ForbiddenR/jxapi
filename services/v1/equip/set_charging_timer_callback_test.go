package equip

import (
	"context"
	"testing"

	"gitee.com/csms/jxeu-ocpp/internal/config"
	"gitee.com/csms/jxeu-ocpp/pkg/api"
	"gitee.com/csms/jxeu-ocpp/pkg/api/services"
	"github.com/stretchr/testify/assert"
)

func TestSetChargingTimerRequest(t *testing.T) {
	config.TestConfig()
	api.Init()

	p := services.OCPP16()
	ctx := context.TODO()
	req := []*equipSetChargingTimerCallbackRequest{
		{
			Base: services.Base{
				EquipmentSn: services.TestSN,
				Protocol:    p,
				Category:    services.SetChargingTimer.FirstUpper(),
				AccessPod:   services.TestAccessPod,
				MsgID:       "1",
			},
			Callback: services.NewCB(services.Successful),
		},
	}

	for _, value := range req {
		err := SetChargingTimerCallbackRequest(ctx, value)
		assert.Nil(t, err)
	}
}

func TestSetChargingTimerRequestWithGeneric(t *testing.T) {
	config.TestConfig()
	api.Init()

	p := services.OCPP16()
	ctx := context.TODO()
	req := []*equipSetChargingTimerCallbackRequest{
		{
			Base: services.Base{
				EquipmentSn: services.TestSN,
				Protocol:    p,
				Category:    services.SetChargingTimer.FirstUpper(),
				AccessPod:   services.TestAccessPod,
				MsgID:       "1",
			},
			Callback: services.NewCB(services.Successful),
		},
	}

	for _, value := range req {
		err := SetChargingTimerCallbackRequestWithGeneric(ctx, value)
		assert.Nil(t, err)
	}
}
