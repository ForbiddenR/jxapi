package equip

import (
	"context"
	"testing"

	"github.com/Kotodian/gokit/id"
	"github.com/stretchr/testify/assert"

	// api "github.com/ForbiddenR/jx-api"
	"gitee.com/csms/jxeu-ocpp/pkg/api/services"
)

func TestAuthorizeTransactionRequestWithGeneric(t *testing.T) {
	// config.TestConfig()
	// api.Init()
	// log.InitNopLogger()

	ctx := context.TODO()
	p := services.OCPP16()

	reqOnline := newTestEquipOnlineRequest(services.TestSN, p, services.TestAccessPod, id.Next().String())
	equipID, err := OnlineRequestWithGeneric(ctx, reqOnline)

	assert.Nil(t, err)
	assert.NotEmpty(t, equipID)

	req := map[string]*equipAuthorizeTransactionRequest{
		"valid id token":   newTestEquipAuthorizeTransactionRequest(p, "fbf8e1faa42c18ba98330b77ca97d29a"),
		"invalid id token": newTestEquipAuthorizeTransactionRequest(p, "111111"),
		"default id Token": newTestEquipAuthorizeTransactionRequest(p, "ffffffff"),
		"expired id Token": newTestEquipAuthorizeTransactionRequest(p, "bbbbbbbb"),
	}

	for n, v := range req {
		t.Run(n, func(t *testing.T) {
			resp, err := AuthorizeTransactionRequestWithGeneric(ctx, v)
			assert.Nil(t, err)
			assert.NotNil(t, resp)
			t.Log(resp.Data.IdTokenInfo.Status)
		})

	}

	reqOffline := newTestEquipOfflineRequest(services.TestSN, p, services.TestAccessPod, id.Next().String(), EOF)
	err = OfflineRequestWithGeneric(ctx, reqOffline)
	assert.Nil(t, err)
}

func newTestEquipAuthorizeTransactionRequest(p *services.Protocol, idToken string) *equipAuthorizeTransactionRequest {
	req := NewEquipAuthorizeTransactionRequest(services.TestSN, services.TestAccessPod, id.Next().String(), p)
	req.Data.IdTokenType = IdTokenType{
		IdToken: idToken,
	}
	return req
}

// func TestAuthorizeTransactionRequest(t *testing.T) {
// 	// config.TestConfig()
// 	// api.Init()
// 	// log.InitNopLogger()

// 	ctx := context.TODO()
// 	p := services.OCPP16()

// 	reqOnline := &equipOnlineRequest{
// 		Base: services.Base{
// 			EquipmentSn: services.TestSN,
// 			Protocol:    p,
// 			Category:    services.Online.FirstUpper(),
// 			AccessPod:   services.TestAccessPod,
// 			MsgID:       id.Next().String(),
// 		},
// 		Data: &equipOnlineRequestDetail{
// 			RemoteAddress: nil,
// 		},
// 	}

// 	err := OnlineRequest(ctx, reqOnline)

// 	assert.Nil(t, err)

// 	idToken := "b361e850"
// 	req := []*equipAuthorizeTransactionRequest{
// 		{
// 			Base: services.Base{
// 				EquipmentSn: services.TestSN,
// 				Protocol:    p,
// 				Category:    services.Authorize.FirstUpper(),
// 				AccessPod:   services.TestAccessPod,
// 			},
// 			Data: &equipAuthorizeTransactionRequestDetail{
// 				IdTokenType: IdTokenType{
// 					IdToken: idToken,
// 				},
// 			},
// 		},
// 	}

// 	for _, v := range req {
// 		resp, err := AuthorizeTransactionRequest(ctx, v)
// 		assert.Nil(t, err)
// 		assert.NotNil(t, resp)
// 		t.Log(resp.Data)
// 	}

// 	reqOffline := &equipOfflineRequest{
// 		Base: services.Base{
// 			EquipmentSn: services.TestSN,
// 			Protocol:    p,
// 			AccessPod:   services.TestAccessPod,
// 			MsgID:       id.Next().String(),
// 		},
// 	}

// 	err = OfflineRequest(ctx, reqOffline)
// 	assert.Nil(t, err)
// }