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
	services.RegisterFC(services.SetLoadBalanceFeatureName, func(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) services.CallbackRequest {
		return NewEquipSetLoadBalanceRequestError(sn, pod, msgID, p, err)
	})
	services.RegisterFC(services.CancelReservationFeatureName,func(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) services.CallbackRequest {
		return NewEquipCancelReservationCallbackRequestError(sn, pod, msgID, p, err)
	})
	services.RegisterFC(services.SetIntellectChargeFeatureName,func(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) services.CallbackRequest {
		return NewEquipSetIntellectChargeCallbackRequestError(sn, pod, msgID, p, err)
	})
}