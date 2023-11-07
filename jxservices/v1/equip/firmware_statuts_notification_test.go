package equip

// import (
// 	"context"
// 	"testing"

// 	"gitee.com/csms/jxeu-ocpp/internal/config"
// 	"gitee.com/csms/jxeu-ocpp/pkg/api"
// 	"gitee.com/csms/jxeu-ocpp/pkg/api/services"
// 	"github.com/stretchr/testify/assert"
// )

// func TestFirmwareStatusNotificationRequest(t *testing.T) {
// 	config.TestConfig()
// 	api.Init()

// 	ctx := context.TODO()
// 	// logOpt := &log.Options{}
// 	// logOpt.Development = true
// 	// log.InitLogger(logOpt)

// 	p := services.OCPP16()
// 	reqOnline := &equipOnlineRequest{
// 		Base: services.Base{
// 			EquipmentSn: services.TestSN,
// 			Protocol:    p,
// 			Category:    services.Online.FirstUpper(),
// 			AccessPod:   services.TestAccessPod,
// 			MsgID:       "1",
// 		},
// 		Data: &equipOnlineRequestDetail{
// 			RemoteAddress: nil,
// 		},
// 	}

// 	err := OnlineRequest(ctx, reqOnline)

// 	assert.Nil(t, err)

// 	req := []*equipFirmwareStatusNotificationRequest{
// 		{
// 			Base: services.Base{
// 				EquipmentSn: services.TestSN,
// 				Protocol:    p,
// 				Category:    services.FirmwareStatusNotification.FirstUpper(),
// 				AccessPod:   services.TestAccessPod,
// 				MsgID:       "1",
// 			},
// 			Data: &equipFirmwareStatusNotificationRequestDetail{
// 				Status: Idle,
// 			},
// 		},
// 	}

// 	for _, v := range req {
// 		err := FirmwareStatusNotificationRequest(ctx, v)
// 		assert.Nil(t, err)
// 	}

// 	reqOffline := &equipOfflineRequest{
// 		Base: services.Base{
// 			EquipmentSn: services.TestSN,
// 			Protocol:    p,
// 			AccessPod:   services.TestAccessPod,
// 			MsgID:       "1",
// 		},
// 	}

// 	err = OfflineRequest(ctx, reqOffline)
// 	assert.Nil(t, err)
// }

// func TestFirmwareStatusNotificationRequestWithGeneric(t *testing.T) {
// 	config.TestConfig()
// 	api.Init()

// 	// logOpt := &log.Options{}
// 	// logOpt.Development = true
// 	// log.InitLogger(logOpt)
// 	p := services.OCPP16()

// 	ctx := context.TODO()
// 	reqOnline := &equipOnlineRequest{
// 		Base: services.Base{
// 			EquipmentSn: services.TestSN,
// 			Protocol:    p,
// 			Category:    services.Online.FirstUpper(),
// 			AccessPod:   services.TestAccessPod,
// 			MsgID:       "1",
// 		},
// 		Data: &equipOnlineRequestDetail{
// 			RemoteAddress: nil,
// 		},
// 	}

// 	equipID, err := OnlineRequestWithGeneric(ctx, reqOnline)

// 	assert.Nil(t, err)
// 	assert.NotEmpty(t, equipID)

// 	req := []*equipFirmwareStatusNotificationRequest{
// 		{
// 			Base: services.Base{
// 				EquipmentSn: services.TestSN,
// 				Protocol:    p,
// 				Category:    services.FirmwareStatusNotification.FirstUpper(),
// 				AccessPod:   services.TestAccessPod,
// 				MsgID:       "1",
// 			},
// 			Data: &equipFirmwareStatusNotificationRequestDetail{
// 				Status: Idle,
// 			},
// 		},
// 	}

// 	for _, v := range req {
// 		err := FirmwareStatusNotificationRequestWithGeneric(ctx, v)
// 		assert.Nil(t, err)
// 	}

// 	reqOffline := &equipOfflineRequest{
// 		Base: services.Base{
// 			EquipmentSn: services.TestSN,
// 			Protocol:    p,
// 			AccessPod:   services.TestAccessPod,
// 			MsgID:       "1",
// 		},
// 	}

// 	err = OfflineRequestWithGeneric(ctx, reqOffline)
// 	assert.Nil(t, err)
// }
