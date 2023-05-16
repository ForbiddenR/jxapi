package equip

import (
	"context"
	"testing"
	"time"

	"gitee.com/csms/jxeu-ocpp/internal/config"
	"gitee.com/csms/jxeu-ocpp/pkg/api"
	"gitee.com/csms/jxeu-ocpp/pkg/api/services"
	"github.com/stretchr/testify/assert"
)

func TestChargingTimerNotification(t *testing.T) {
	config.TestConfig()
	api.Init()

	IdToken := "b361e850"

	ctx := context.TODO()
	p := services.OCPP16()
	now := time.Now().Unix()

	reqOnline := &equipOnlineRequest{
		Base: services.Base{
			EquipmentSn: services.TestSN,
			Protocol:    p,
			Category:    services.Online.FirstUpper(),
			AccessPod:   services.TestAccessPod,
			MsgID:       "1",
		},
		Data: &equipOnlineRequestDetail{
			RemoteAddress: nil,
		},
	}

	err := OnlineRequest(ctx, reqOnline)

	assert.Nil(t, err)

	// noError := StatusNotificationErrorCodeNoError
	reqStatusNotification := NewEquipStatusNotificationRequest(services.TestSN, services.TestAccessPod, "1",
		p, "1", ConnectorStatusPreparing, now)

	// reqStatusNotification.Data.ErrorCode = &noError

	err = StatusNotificationRequest(ctx, reqStatusNotification)

	assert.Nil(t, err)

	reqAuthorization := NewEquipAuthorizeTransactionRequest(services.TestSN, services.TestAccessPod, "1",
		p)

	reqAuthorization.Data.IdTokenType = IdTokenType{IdToken: "b361e850"}

	resp, err := AuthorizeTransactionRequest(ctx, reqAuthorization)

	assert.Nil(t, err)
	assert.NotNil(t, resp)
	version := int64(1)

	reqStartTransaction := NewEquipStartTransactionRequestOCPP16(services.TestSN, services.TestAccessPod, "1",
		IdToken, 0, "1", time.Now().Unix())

	respStartTransaction, err := StartTransactionOCPP16Request(ctx, reqStartTransaction)

	assert.Nil(t, err)
	assert.NotNil(t, respStartTransaction)

	transactionId := respStartTransaction.Data.TransactionId
	t.Log(transactionId)

	req := []*equipChargingTimerNotificationRequest{
		{
			Base: services.Base{
				EquipmentSn: services.TestSN,
				Protocol:    p,
				Category:    services.ChargingTimerNotification.FirstUpper(),
				AccessPod:   services.TestAccessPod,
			},
			Data: &equipChargingTimerNotificationRequestDetail{
				ConnectorSerial: "1",
				TransactionId:   &transactionId,
				TimerId:         3322112333,
				TriggerTime:     &now,
				Version:         &version,
				Status:          ChargingTimerStatusCharging,
			},
		},
	}

	for _, v := range req {
		err := ChargingTimerNotificationRequest(ctx, v)
		assert.Nil(t, err)
	}

	reqOffline := &equipOfflineRequest{
		Base: services.Base{
			EquipmentSn: services.TestSN,
			Protocol:    p,
			AccessPod:   services.TestAccessPod,
			MsgID:       "1",
		},
	}

	err = OfflineRequest(ctx, reqOffline)
	assert.Nil(t, err)
}
