package equip

// import (
// 	"context"
// 	"testing"

// 	"gitee.com/csms/jxeu-ocpp/internal/config"
// 	"gitee.com/csms/jxeu-ocpp/pkg/api"
// 	"gitee.com/csms/jxeu-ocpp/pkg/api/services"
// 	"github.com/samber/lo"
// 	"github.com/stretchr/testify/assert"
// )

// func TestDiagnosticsStatusNotificationRequest(t *testing.T) {
// 	config.TestConfig()
// 	api.Init()

// 	ctx := context.TODO()
// 	// logOpt := &log.Options{}
// 	// logOpt.Development = true
// 	// log.InitLogger(logOpt)

// 	p := services.OCPP16()
// 	reqOnline := NewEquipOnlineRequest(services.TestSN, p, "1", services.TestAccessPod)

// 	err := OnlineRequest(ctx, reqOnline)

// 	assert.Nil(t, err)

// 	req := []*equipDiagnosticsStatusNotificationRequest{
// 		NewEquipDiagnosticsStatusNotificationRequestOCPP16(services.TestSN, services.TestAccessPod, "1",
// 			DiagnosticsStatusNotificationTypeIdle),
// 		NewEquipDiagnosticsStatusNotificationRequestOCPP16(services.TestSN, services.TestAccessPod, "1",
// 			DiagnosticsStatusNotificationTypeUploaded),
// 	}

// 	lo.ForEach(req, func(item *equipDiagnosticsStatusNotificationRequest, _ int) {
// 		err := DiagnosticsStatusNotificationRequestG(ctx, item)
// 		assert.Nil(t, err)
// 	})

// 	reqOffline := NewEquipOfflineRequest(services.TestSN, p, services.TestAccessPod, "1", EOF)

// 	err = OfflineRequest(ctx, reqOffline)
// 	assert.Nil(t, err)
// }
