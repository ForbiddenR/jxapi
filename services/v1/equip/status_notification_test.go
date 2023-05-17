package equip

// import (
// 	"context"
// 	"testing"
// 	"time"

// 	"github.com/Kotodian/gokit/id"
// 	"github.com/stretchr/testify/assert"

// 	"gitee.com/csms/jxeu-ocpp/internal/config"
// 	"gitee.com/csms/jxeu-ocpp/internal/log"
// 	"gitee.com/csms/jxeu-ocpp/pkg/api"
// 	"gitee.com/csms/jxeu-ocpp/pkg/api/services"
// )

// func TestStatusNotificationRequestWithGeneric(t *testing.T) {
// 	config.TestConfig()
// 	api.Init()
// 	ctx := context.TODO()
// 	p := services.OCPP16()
// 	log.InitNopLogger()

// 	reqOnline := newTestEquipOnlineRequest(services.TestSN, p, services.TestAccessPod, id.Next().String())
// 	_, err := OnlineRequestWithGeneric(ctx, reqOnline)

// 	assert.Nil(t, err)
// 	cid := "1"

// 	req := []*equipStatusNotificationRequest{
// 		// ocpp1.6正常枪状态
// 		newTestEquipStatusNotificationRequest(p, cid, ConnectorStatusAvailable),
// 		// ocpp1.6发生一次枪状态变化
// 		newTestEquipStatusNotificationRequest(p, cid, ConnectorStatusOccupied),
// 		// ocpp1.6上报自定义告警
// 		newTestEquipStatusNotificationRequestWithErrorVendorCode(cid, ConnectorStatusOccupied, "0xff0c01"),
// 		// ocpp1.6异常枪状态
// 		newTestEquipStatusNotificationRequestWithErrorCode(cid, ConnectorStatusFaulted, StatusNotificationErrorCodeEVCommunicationError),
// 		// ocpp1.6枪恢复正常状态
// 		newTestEquipStatusNotificationRequest(p, cid, ConnectorStatusAvailable),
// 	}

// 	for _, v := range req {
// 		err = StatusNotificationRequestWithGeneric(ctx, v)

// 		assert.Nil(t, err)
// 	}

// 	reqOffline := newTestEquipOfflineRequest(services.TestSN, p, services.TestAccessPod, id.Next().String(), EOF)
// 	err = OfflineRequestWithGeneric(ctx, reqOffline)
// 	assert.Nil(t, err)
// }

// func newTestEquipStatusNotificationRequest(p *services.Protocol, connectorId string, status ConnectorStatusTypeEnum) *equipStatusNotificationRequest {
// 	now := time.Now().Unix()
// 	switch p {
// 	case services.OCPP16():
// 		return NewEquipStatusNotificationRequestOCPP16(services.TestSN, services.TestAccessPod, id.Next().String(), connectorId, status, StatusNotificationErrorCodeNoError, now)
// 	default:
// 		return NewEquipStatusNotificationRequest(services.TestSN, services.TestAccessPod, id.Next().String(), p, connectorId, status, now)
// 	}
// }

// func newTestEquipStatusNotificationRequestWithErrorCode(connectorId string, status ConnectorStatusTypeEnum, errorCode StatusNotificationErrorCodeEnum) *equipStatusNotificationRequest {
// 	return NewEquipStatusNotificationRequestOCPP16(services.TestSN, services.TestAccessPod, id.Next().String(), connectorId, status, errorCode, time.Now().Unix())
// }

// func newTestEquipStatusNotificationRequestWithErrorVendorCode(connectorId string, status ConnectorStatusTypeEnum, vendorErrorCode string) *equipStatusNotificationRequest {
// 	r := newTestEquipStatusNotificationRequestWithErrorCode(connectorId, status, StatusNotificationErrorCodeOtherError)
// 	r.Data.VendorErrorCode = &vendorErrorCode
// 	return r
// }
