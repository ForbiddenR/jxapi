package equip

import (
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

func Init() {
	services.RegisterFC(services.QRCode.String(), func(s1, s2, s3 string, p *services.Protocol, ce *apierrors.CallbackError) services.CallbackRequest {
		return NewEquipSendQRCodeCallbackRequestError(s1, s2, s3, p, ce)
	})
}