package equip

import (
	"encoding/json"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

type EquipSetChargingProfileRequest struct {
	services.Base
}

func (r *EquipSetChargingProfileRequest) UnmarshalJSON(data []byte) error {
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