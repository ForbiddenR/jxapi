package equip

import (
	"encoding/json"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

type EquipClearChargingProfileRequest struct {
	services.Base
}

func (r *EquipClearChargingProfileRequest) UnmarshalJSON(data []byte) error {
	type Alias EquipSetChargingProfileRequest
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