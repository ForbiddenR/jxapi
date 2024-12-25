package equip

import (
	"encoding/json"
	"errors"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

type EquipGetVariableListRequest struct {
	services.Base
	Data []EquipGetVariablesRequestDetail `json:"data"`
}

func (r *EquipGetVariableListRequest) UnmarshalJSON(data []byte) error {
	type Alias EquipGetVariableListRequest
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
	if !aux.Protocol.Equal(services.OCPP16()) {
		if len(aux.Data) == 0 {
			return errors.New(aux.Protocol.String() + ":data is nil")
		}
	}
	return nil
}
