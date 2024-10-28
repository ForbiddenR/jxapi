package equip

import (
	"encoding/json"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type EquipResetOfflineCardsRequest struct {
	services.Base
	Data *EquipResetOfflineCardsRequestDetail `json:"data"`
}

type EquipResetOfflineCardsRequestDetail struct {
	OfflineCard []string `json:"offlineCard"`
}

func (r *EquipResetOfflineCardsRequest) UnmarshalJSON(data []byte) error {
	type Alias EquipResetOfflineCardsRequest
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
