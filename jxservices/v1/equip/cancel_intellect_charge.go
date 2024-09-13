package equip

import (
	"encoding/json"
	"errors"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

type EquipCancelIntellectChargeRequest struct {
	services.Base
	Data *EquipCancelIntellectChargeRequestDetail `json:"data"`
}

type EquipCancelIntellectChargeRequestDetail struct {
	VendorId    string      `json:"vendorId"`
	TimingType  int         `json:"timingType"`
	Evse        EVSE        `json:"evse"`
	IdTokenType IdTokenType `json:"idTokenType"`
}

func (s *EquipCancelIntellectChargeRequest) Unmarshal(data []byte) error {
	type Alias EquipCancelIntellectChargeRequest
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
