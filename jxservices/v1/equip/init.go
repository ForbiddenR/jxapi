package equip

import (
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

func init() {
	services.InitFC()
	services.RegisterFC(services.SendQRCodeFeatureName, func(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) services.CallbackRequest {
		return NewEquipSendQRCodeCallbackRequestError(sn, pod, msgID, p, err)
	})
}