package equip

import (
	"encoding/json"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type EquipSetChargeTemplateRequest struct {
	services.Base
}

func (r *EquipSetChargeTemplateRequest) UnmarshalJSON(data []byte) error {
	type Alias EquipSetChargeTemplateRequest
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
