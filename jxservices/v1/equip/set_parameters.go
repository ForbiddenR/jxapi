package equip

import (
	"encoding/json"
	"errors"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type EquipSetParametersRequest struct {
	services.Base
	Data *EquipSetParametersRequestDetail `json:"data"`
}

type EquipSetParametersRequestDetail struct {
	Avaliability   uint8 `json:"avaliability"`
	MaxOutputPower uint8 `json:"maxOutputPower"`
}

func (r *EquipSetParametersRequest) UnmarshalJSON(data []byte) error {
	type Alias EquipSetParametersRequest
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if aux.Data == nil {
		return errors.New("data is nil")
	}
	return nil
}
