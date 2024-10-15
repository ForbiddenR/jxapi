package equip

import (
	"encoding/json"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type EquipClearChargingProfileRequest struct {
	services.Base
	Data *EquipClearchargingProfileRequestDetail `json:"data"`
}

type EquipClearchargingProfileRequestDetail struct {
	ID int64 `json:"id"`
}

func (r *EquipClearChargingProfileRequest) UnmarshalJSON(data []byte) error {
	type Alias EquipClearChargingProfileRequest
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
