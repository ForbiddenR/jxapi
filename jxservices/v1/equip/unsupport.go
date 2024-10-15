package equip

import (
	"context"
	"fmt"

	"github.com/ForbiddenR/jxapi/v2/apierrors"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

func UnspportedCallbackRequest(ctx context.Context, name string, base services.Base) error {
	fn, ok := services.UnsupportedFeatures.Get(name)
	if !ok {
		return fmt.Errorf("this feature hasn't implemented. name: %s, id: %s", name, base.MsgID)
	}
	err := apierrors.NewCallbackErrorNotSupported(base.EquipmentSn, name)
	req := fn(base, err)
	return services.Transport(ctx, req)
}
