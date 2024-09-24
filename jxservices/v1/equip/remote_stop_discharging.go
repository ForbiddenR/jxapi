package equip

import (
	"encoding/json"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

type EquipRequestStopDischargingRequest struct {
	services.Base
	Data *EquipRequestStartDischargingRequestDetail `json:"data"`
}

type EquipRequestStopDischargingRequestDetail struct {
	ConnectorId string `json:"connectorSerial"`
}

func (r *EquipRequestStopDischargingRequest) UnmarshalJSON(data []byte) error {
	type Alias EquipRequestStartDischargingRequest
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
