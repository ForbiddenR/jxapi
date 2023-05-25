package equip

import (
	"encoding/json"
	"errors"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

type EquipSetIntellectChargeRequest struct {
	services.Base
	Data *EquipSetIntellectChargeRequestDetail `json:"data"`
}

type EquipSetIntellectChargeRequestDetail struct {
}

func (s *EquipSetIntellectChargeRequest) Unmarshal(data []byte) error {
	type Alias EquipSetIntellectChargeRequest
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if aux.Data == nil {
		return errors.New("data is nil")
	}
	return nil
}
