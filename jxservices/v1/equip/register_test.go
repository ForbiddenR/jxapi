package equip

// func TestEquipRegisterRequest(t *testing.T) {
// 	config.TestConfig()
// 	api.Init()
// 	log.InitNopLogger()
// 	ctx := context.TODO()
// 	p := services.OCPP16()

// 	remoteAddress := "127.0.0.1"
// 	req := []*equipRegisterRequest{
// 		newTestEquipRegisterRequest(services.TestSN, p, services.TestAccessPod, id.Next().String()),
// 		newTestEquipRegisterRequestWithRemoteAddress(services.TestSN, p, services.TestAccessPod, id.Next().String(), remoteAddress),
// 	}

// 	for _, v := range req {
// 		err := RegisterRequest(ctx, v)
// 		assert.Nil(t, err)
// 	}
// }

// func TestAddEquipRegisterRequestWithGeneric(t *testing.T) {
// 	config.TestConfig()
// 	api.Init()
// 	log.InitNopLogger()

// 	ctx := context.TODO()
// 	p := services.OCPP16()
// 	raddr := "127.0.0.1"

// 	req := newTestEquipRegisterRequest(services.TestSN, p, services.TestAccessPod, id.Next().String())

// 	equipid, err := RegisterRequestWithGeneric(ctx, req)
// 	assert.Nil(t, err)
// 	assert.NotEmpty(t, equipid)

// 	reqOffline := newTestEquipOfflineRequest(services.TestSN, p, services.TestAccessPod, id.Next().String(), EOF)

// 	err = OfflineRequestWithGeneric(ctx, reqOffline)
// 	assert.Nil(t, err)

// 	req = newTestEquipRegisterRequestWithRemoteAddress(services.TestSN, p, services.TestAccessPod, id.Next().String(), raddr)

// 	equipid, err = RegisterRequestWithGeneric(ctx, req)
// 	assert.Nil(t, err)
// 	assert.NotEmpty(t, equipid)

// 	reqOffline = newTestEquipOfflineRequest(services.TestSN, p, services.TestAccessPod, id.Next().String(), EOF)

// 	err = OfflineRequestWithGeneric(ctx, reqOffline)
// 	assert.Nil(t, err)
// }

// func newTestEquipRegisterRequest(sn string, p *services.Protocol, pod, msgID string) *equipRegisterRequest {
// 	return NewEquipRegisterRequest(sn, p, pod, msgID)
// }

// func newTestEquipRegisterRequestWithRemoteAddress(sn string, p *services.Protocol, pod, msgID, remoteAddress string) *equipRegisterRequest {
// 	req := NewEquipRegisterRequest(sn, p, pod, msgID)
// 	req.Data.RemoteAddress = &remoteAddress
// 	return req
// }
