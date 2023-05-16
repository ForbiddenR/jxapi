package equip

import (
	"encoding/json"
	"errors"

	"gitee.com/csms/jxeu-ocpp/pkg/api/services"
)

type EquipSetVariablesRequest struct {
	services.Base
	Data *EquipSetVariablesRequestDetail `json:"data"`
}

func (r *EquipSetVariablesRequest) UnmarshalJSON(data []byte) error {
	type Alias EquipSetVariablesRequest
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
		if aux.Data.Component == nil {
			return errors.New(aux.Protocol.String() + ":component is nil")
		}
	}
	return nil
}

type EquipSetVariablesRequestDetail struct {
	Component *Component `json:"component,omitempty"`
	Key       string     `json:"key" validate:"required"`
	Value     string     `json:"value" validate:"requried"`
}
