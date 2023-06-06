package equip

// import (
// 	"context"
// 	"gitee.com/csms/jxeu-ocpp/internal/config"
// 	"gitee.com/csms/jxeu-ocpp/internal/errors"
// 	"gitee.com/csms/jxeu-ocpp/pkg/api"
// 	"gitee.com/csms/jxeu-ocpp/pkg/api/services"
// 	"gitee.com/csms/jxeu-ocpp/pkg/ocpp1.6/protocol"
// 	"github.com/stretchr/testify/assert"
// 	"testing"
// )

// func TestEquipSetLoadBalanceRequest(t *testing.T) {
// 	config.TestConfig()
// 	api.Init()
// 	// logOpt := &log.Options{}
// 	// logOpt.Development = true
// 	// log.InitLogger(logOpt)
// 	ctx := context.TODO()
// 	p := services.OCPP16()

// 	err := errors.NewCallbackErrorGenericError(services.TestSN, protocol.DataTransferFeatureName, "test err")
// 	req := []*equipSetLoadBalanceRequest{
// 		NewEquipSetLoadBalanceRequest(services.TestSN, services.TestAccessPod, "1", p, services.Successful),
// 		NewEquipSetLoadBalanceRequestError(services.TestSN, services.TestAccessPod, "1", p, err),
// 	}

// 	for _, v := range req {
// 		err := SetLoadBalanceRequestWithG(ctx, v)
// 		assert.Nil(t, err)
// 	}
// }
