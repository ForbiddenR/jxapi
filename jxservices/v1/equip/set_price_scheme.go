package equip

import (
	"encoding/json"
	"errors"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

type EquipSetPriceSchemeRequest struct {
	services.Base
	Data *EquipSetPriceSchemeRequestDetail `json:"data"`
}

type EquipSetPriceSchemeRequestDetail struct {
}

func (s *EquipSetPriceSchemeRequest) Unmarshal(data []byte) error {
	type Alias EquipSetPriceSchemeRequest
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
