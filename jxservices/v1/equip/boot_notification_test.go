package equip

// import (
// 	"context"
// 	"testing"

// 	"github.com/Kotodian/gokit/id"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/ForbiddenR/jx-api/services"

// )

// func TestBootNotificationRequestWithGeneric(t *testing.T) {
// 	// config.TestConfig()
// 	// api.Init()
// 	// log.InitNopLogger()

// 	ctx := context.TODO()
// 	p := services.OCPP16()

// 	reqOnline := newTestEquipOnlineRequest(services.TestSN, p, services.TestAccessPod, id.Next().String())
// 	equipID, err := OnlineRequestWithGeneric(ctx, reqOnline)

// 	assert.Nil(t, err)
// 	assert.NotEmpty(t, equipID)

// 	//remoteAddress := "127.0.0.1"
// 	modelCode := "NJ4GP122A0"
// 	manufacturerCode := "MT01CNJQX0"
// 	firmwareVersion := "0.0.1"
// 	iccid := "898602YY8SSXXXXXXXXP"
// 	req := []*equipBootNotificationRequest{
// 		newTestEquipBootNotificationRequest(p, modelCode, manufacturerCode, firmwareVersion),
// 		newTestEquipBootNotificationRequestWithICCID(p, modelCode, manufacturerCode, firmwareVersion, iccid),
// 	}
// 	for _, v := range req {
// 		err = BootNotificationRequestWithGeneric(ctx, v)

// 		assert.Nil(t, err)
// 	}

// 	reqOffline := newTestEquipOfflineRequest(services.TestSN, p, services.TestAccessPod, id.Next().String(), EOF)
// 	err = OfflineRequestWithGeneric(ctx, reqOffline)
// 	assert.Nil(t, err)
// }

// func newTestEquipBootNotificationRequest(p *services.Protocol, modelCode, manufacturerCode, firmwareVersion string) *equipBootNotificationRequest {
// 	r := NewEquipBootNotificationRequest(services.TestSN, services.TestAccessPod, id.Next().String(), p)
// 	r.Data.ModelCode = modelCode
// 	r.Data.ManufacturerCode = manufacturerCode
// 	r.Data.FirmwareVersion = &firmwareVersion
// 	return r
// }

// func newTestEquipBootNotificationRequestWithICCID(p *services.Protocol, modelCode, manufacturerCode, firmwareVersion, iccid string) *equipBootNotificationRequest {
// 	r := NewEquipBootNotificationRequest(services.TestSN, services.TestAccessPod, id.Next().String(), p)
// 	r.Data.ModelCode = modelCode
// 	r.Data.ManufacturerCode = manufacturerCode
// 	r.Data.FirmwareVersion = &firmwareVersion
// 	r.Data.Iccid = &iccid
// 	return r
// }

// // func TestBootNotificationRequest(t *testing.T) {
// // 	config.TestConfig()
// // 	api.Init()
// // 	ctx := context.TODO()

// // 	p := services.OCPP16()
// // 	reqOnline := newTestEquipOnlineRequest(services.TestSN, p, services.TestAccessPod, id.Next().String())
// // 	err := OnlineRequest(ctx, reqOnline)

// // 	assert.Nil(t, err)

// // 	//remoteAddress := "127.0.0.1"
// // 	modelCode := "NJC0P121B0"
// // 	manufacturerCode := "MT01CNJQX0"
// // 	firmwareVersion := "0.0.1"
// // 	iccid := "123000000001"
// // 	req := []*equipBootNotificationRequest{
// // 		{
// // 			Base: services.Base{
// // 				EquipmentSn: services.TestSN,
// // 				Protocol:    p,
// // 				Category:    services.BootNotification.FirstUpper(),
// // 				AccessPod:   services.TestAccessPod,
// // 				MsgID:       "1",
// // 			},
// // 			Data: &equipBootNotificationRequestDetail{
// // 				ModelCode:        modelCode,
// // 				ManufacturerCode: manufacturerCode,
// // 				FirmwareVersion:  nil,
// // 				Iccid:            nil,
// // 			},
// // 		},
// // 		{
// // 			Base: services.Base{
// // 				EquipmentSn: services.TestSN,
// // 				Protocol:    p,
// // 				Category:    services.BootNotification.FirstUpper(),
// // 				AccessPod:   services.TestAccessPod,
// // 				MsgID:       "1",
// // 			},
// // 			Data: &equipBootNotificationRequestDetail{
// // 				ModelCode:        modelCode,
// // 				ManufacturerCode: manufacturerCode,
// // 				FirmwareVersion:  nil,
// // 				Iccid:            nil,
// // 			},
// // 		},
// // 		{
// // 			Base: services.Base{
// // 				EquipmentSn: services.TestSN,
// // 				AccessPod:   services.TestAccessPod,
// // 				Category:    services.BootNotification.FirstUpper(),
// // 				Protocol:    p,
// // 				MsgID:       "1",
// // 			},
// // 			Data: &equipBootNotificationRequestDetail{
// // 				ModelCode:        modelCode,
// // 				ManufacturerCode: manufacturerCode,
// // 				FirmwareVersion:  &firmwareVersion,
// // 				Iccid:            nil,
// // 			},
// // 		},
// // 		{
// // 			Base: services.Base{
// // 				EquipmentSn: services.TestSN,
// // 				AccessPod:   services.TestAccessPod,
// // 				Category:    services.BootNotification.FirstUpper(),
// // 				Protocol:    p,
// // 				MsgID:       "1",
// // 			},
// // 			Data: &equipBootNotificationRequestDetail{
// // 				ModelCode:        modelCode,
// // 				ManufacturerCode: manufacturerCode,
// // 				FirmwareVersion:  nil,
// // 				Iccid:            &iccid,
// // 			},
// // 		},
// // 		{
// // 			Base: services.Base{
// // 				EquipmentSn: services.TestSN,
// // 				AccessPod:   services.TestAccessPod,
// // 				Category:    services.BootNotification.FirstUpper(),
// // 				Protocol:    p,
// // 				MsgID:       "1",
// // 			},
// // 			Data: &equipBootNotificationRequestDetail{
// // 				ModelCode:        modelCode,
// // 				ManufacturerCode: manufacturerCode,
// // 				FirmwareVersion:  &firmwareVersion,
// // 				Iccid:            &iccid,
// // 			},
// // 		},
// // 	}

// // 	for _, v := range req {
// // 		err = BootNotificationRequest(ctx, v)

// // 		assert.Nil(t, err)
// // 	}

// // 	reqOffline := &equipOfflineRequest{
// // 		Base: services.Base{
// // 			EquipmentSn: services.TestSN,
// // 			Protocol:    p,
// // 			AccessPod:   services.TestAccessPod,
// // 			MsgID:       "1",
// // 		},
// // 	}

// // 	err = OfflineRequest(ctx, reqOffline)
// // 	assert.Nil(t, err)
// // }
