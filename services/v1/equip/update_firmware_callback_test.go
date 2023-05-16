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

func TestUpdateFirmwareRequest(t *testing.T) {
	config.TestConfig()
	api.Init()
	ctx := context.TODO()

	err := errors.NewCallbackErrorGenericError(services.TestSN, protocol.RemoteStopTransactionFeatureName, "test err")
	req := []*equipUpdateFirmwareCallbackRequest{
		{
			Base: services.Base{
				EquipmentSn: services.TestSN,
				Protocol:    &services.Protocol{Name: "OCPP", Version: "1.6"},
				Category:    services.UpdateFirmware.FirstUpper(),
				AccessPod:   services.TestAccessPod,
				MsgID:       "1",
			},
			Callback: services.NewCB(services.Successful),
		},
		{
			Base: services.Base{
				EquipmentSn: services.TestSN,
				Protocol:    &services.Protocol{Name: "OCPP", Version: "2.0.1"},
				Category:    services.UpdateFirmware.FirstUpper(),
				AccessPod:   services.TestAccessPod,
				MsgID:       "1",
			},
			Callback: services.NewCBError(err),
		},
	}

	for _, v := range req {
		err := UpdateFirmwareCallbackRequest(ctx, v)
		assert.Nil(t, err)
	}
}

func TestUpdateFirmwareCallbackRequestWithGeneric(t *testing.T) {
	config.TestConfig()
	api.Init()
	// logOpt := &log.Options{}
	// logOpt.Development = true
	// log.InitLogger(logOpt)
	ctx := context.TODO()

	err := errors.NewCallbackErrorGenericError(services.TestSN, protocol.RemoteStopTransactionFeatureName, "test err")
	req := []*equipUpdateFirmwareCallbackRequest{
		{
			Base: services.Base{
				EquipmentSn: services.TestSN,
				Protocol: &services.Protocol{
					Name:    "OCPP",
					Version: "1.6",
				},
				Category:  services.UpdateFirmware.FirstUpper(),
				AccessPod: services.TestAccessPod,
			},
			Callback: services.NewCB(services.Successful),
		},
		{
			Base: services.Base{
				EquipmentSn: services.TestSN,
				Protocol: &services.Protocol{
					Name:    "OCPP",
					Version: "1.6",
				},
				Category:  services.UpdateFirmware.FirstUpper(),
				AccessPod: services.TestAccessPod,
			},
			Callback: services.NewCBError(err),
		},
	}

	lo.ForEach(req, func(v *equipUpdateFirmwareCallbackRequest, _ int) {
		err := UpdateFirmwareCallbackRequestWithGeneric(ctx, v)
		assert.Nil(t, err)
	})
}
