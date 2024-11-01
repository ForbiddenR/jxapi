package equip

import (
	"encoding/json"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type EquipSyncTimeRequest struct {
	services.Base
}

func (r *EquipSyncTimeRequest) UnmarshalJSON(data []byte) error {
	type Alias EquipSyncTimeRequest
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
