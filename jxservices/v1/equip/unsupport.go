package equip

import (
	"context"
	"fmt"

	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

func UnspportedCallbackRequest(ctx context.Context, name, sn, pod, msgID string, p *services.Protocol, err *apierrors.CallbackError) error {
	fn, ok := services.FetchFC(name)
	if !ok {
		return fmt.Errorf("this feature hasn't implemented. name: %s, id: %s", name, msgID)
	}
	req := fn(sn, pod, msgID, p, err)

	header := services.GetSimpleHeaderValue(services.Request2ServicesNameType(req.GetName()))

	url := services.GetSimpleURL(req)

	return services.RequestGeneral(ctx, req, url, header)
}
