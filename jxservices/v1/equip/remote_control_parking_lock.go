package equip

import (
	"encoding/json"
	"errors"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type EquipRemoteControlParkingLockRequest struct {
	services.Base
	Data *EquipRemoteControlParkingLockRequestDetail `json:"data"`
}

type EquipRemoteControlParkingLockRequestDetail struct {
	ConnectorId string `json:"connectorSerial"`
	Status      uint8  `json:"status"`
}

func (r *EquipRemoteControlParkingLockRequest) UnmarshalJSON(data []byte) error {
	type Alias EquipRemoteControlParkingLockRequest
	aux := &struct {
		*Alias
	}{Alias: (*Alias)(r)}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if aux.Data == nil {
		return errors.New("data is nil")
	}
	return nil
}
