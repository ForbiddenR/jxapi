package equip

import (
	"encoding/json"
	"errors"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

type EquipSendQRCodeRequest struct {
	services.Base
	Data *EquipSendQRCodeRequestDetail `json:"data"`
}

type EquipSendQRCodeRequestDetail struct {
	
}

func (s *EquipSendQRCodeRequest) UnmarshalJSON(data []byte) error {
	type Alias EquipSendQRCodeRequest
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
