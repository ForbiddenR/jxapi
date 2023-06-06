package jxservices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceRequestNameServices(t *testing.T) {
	assert.Equal(t, Register.FirstUpper(), "EquipRegister")
	assert.Equal(t, Register.Split(), []string{"register"})

	assert.Equal(t, Online.FirstUpper(), "EquipOnline")
	assert.Equal(t, Online.Split(), []string{"online"})

	assert.Equal(t, BootNotification.FirstUpper(), "BootNotification")
	assert.Equal(t, BootNotification.Split(), []string{"boot", "notification"})

	assert.Equal(t, StatusNotification.FirstUpper(), "StatusNotification")
	assert.Equal(t, StatusNotification.Split(), []string{"status", "notification"})

	assert.Equal(t, Authorize.FirstUpper(), "AuthorizeTransaction")
	assert.Equal(t, Authorize.Split(), []string{"authorize"})

	assert.Equal(t, RemoteStartTransaction.Split(), []string{"remote", "start", "transaction"})

	assert.Equal(t, RemoteStopTransaction.Split(), []string{"remote", "stop", "transaction"})

	assert.Equal(t, RemoteStartTransaction.FirstUpper(), "RemoteStartTransaction")

	assert.Equal(t, UpdateFirmware.Split(), []string{"push", "firmware", "equipment"})

	assert.Equal(t, FirmwareStatusNotification.Split(), []string{"push", "firmware", "notification"})
}

func TestURL(t *testing.T) {
	assert.Equal(t, "ac/firmwareStatusNotification",
		Equip+"/"+"firmwareStatusNotification")

	assert.Equal(t, "ac/callback/pushFirmwareCallback",
		Equip+"/"+Callback+"/"+UpdateFirmware.String()+"Callback")
}

//func TestTransferFeatureName(t *testing.T) {
//	remoteStart := protocol.RemoteStartTransactionFeatureName
//	assert.Equal(t, transferFeatureName(remoteStart), "remote start transaction")
//
//	remoteStop := protocol.RemoteStopTransactionFeatureName
//	assert.Equal(t, transferFeatureName(remoteStop), "remote stop transaction")
//
//	reset := protocol.ResetFeatureName
//	assert.Equal(t, transferFeatureName(reset), "reset")
//
//	changeConfiguration := protocol.ChangeConfigurationFeatureName
//	assert.Equal(t, transferFeatureName(changeConfiguration), "change configuration")
//
//	getConfiguration := protocol.GetConfigurationFeatureName
//	assert.Equal(t, transferFeatureName(getConfiguration), "get configuration")
//
//	sendLocalList := protocol.SendLocalListFeatureName
//	assert.Equal(t, transferFeatureName(sendLocalList), "send local list")
//}

func TestRequestWithoutResponse(t *testing.T) {
	//req := equip.NewEquipResetRequest("sn", "pod", api.Ocpp16, 1)
	//url := ""
	//header := make(map[string]string)
	//err := RequestWithoutResponse(req, url, header)
	//assert.NotNil(t, err)
	//t.Log(err)

	//resp, err :=RequestWithResponse(req, url, header)
	//assert.Nil(t, err)
	//t.Log(resp)
}
