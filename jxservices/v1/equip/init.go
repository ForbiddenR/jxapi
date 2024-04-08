package equip

import (
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

func init() {
	services.UnsupportedFeatures.Set(services.SendQRCodeFeatureName, func(base services.Base, err *apierrors.CallbackError) services.Request {
		return NewSendQRCodeCallbackRequest(base, services.WithError(err))
	})
	// services.UnsupportedFeatures.Set(services.SendQRCodeFeatureName, func(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) services.Request {
	// 	return NewEquipSendQRCodeCallbackRequestError(sn, pod, msgID, p, err)
	// })
	// services.UnsupportedFeatures.Set(services.SetLoadBalanceFeatureName, func(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) services.Request {
	// 	return NewEquipSetLoadBalanceRequestError(sn, pod, msgID, p, err)
	// })
	// services.UnsupportedFeatures.Set(services.CancelReservationFeatureName, func(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) services.Request {
	// 	return NewEquipCancelReservationCallbackRequestError(sn, pod, msgID, p, err)
	// })
	// services.UnsupportedFeatures.Set(services.SetIntellectChargeFeatureName, func(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) services.Request {
	// 	return NewEquipSetIntellectChargeCallbackRequestError(sn, pod, msgID, p, err)
	// })
	// services.UnsupportedFeatures.Set(services.SetFactoryResetFeatureName, func(sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) services.Request {
	// 	return NewEquipSetFactoryResetRequestError(sn, pod, msgID, p, err)
	// })
}
