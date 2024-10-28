package equip

import (
	"encoding/json"
	"errors"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type EquipRemoteUpdateAccountBalanceRequest struct {
	services.Base
	Data *EquipRemoteUpdateAccountBalanceRequestDetail `json:"data"`
}

func (r *EquipRemoteUpdateAccountBalanceRequest) UnmarshalJSON(data []byte) error {
	type Alias EquipRemoteUpdateAccountBalanceRequest
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
	return nil
}

type EquipRemoteUpdateAccountBalanceRequestDetail struct {
	ConnectorId string  `json:"connectorSerial"`
	Card        string  `json:"card"`
	Balance     float64 `json:"balance"`
}
