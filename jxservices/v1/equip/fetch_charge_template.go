package equip

import (
	"encoding/json"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type EquipFetchChargeTemplateRequest struct {
	services.Base
}

func (r *EquipFetchChargeTemplateRequest) UnmarshalJSON(data []byte) error {
	type Alias EquipFetchChargeTemplateRequest
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
