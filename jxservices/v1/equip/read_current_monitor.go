package equip

import (
	"encoding/json"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type EquipReadCurrentMonitorRequest struct {
	services.Base
	Data *EquipReadCurrentMonitorRequestDetail `json:"data"`
}

type EquipReadCurrentMonitorRequestDetail struct {
	ConnectorId string `json:"connectorSerial" validate:"required"`
}

func (r *EquipReadCurrentMonitorRequest) UnmarshalJSON(b []byte) error {
	type Alias EquipReadCurrentMonitorRequest
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	return nil
}
