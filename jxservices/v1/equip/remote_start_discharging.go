package equip

import (
	"encoding/json"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type EquipRequestStartDischargingRequest struct {
	services.Base
	Data *EquipRequestStartDischargingRequestDetail `json:"data"`
}

type EquipRequestStartDischargingRequestDetail struct {
	VendorId    string `json:"vendorId"`
	ConnectorId string `json:"connectorSerial"`
}

func (r *EquipRequestStartDischargingRequest) UnmarshalJSON(data []byte) error {
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
