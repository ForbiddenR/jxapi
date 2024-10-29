package equip

import (
	"encoding/json"
	"errors"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type EquipClearOfflineCardsRequest struct {
	services.Base
	Data *EquipClearOfflineCardsRequestDetail `json:"data"`
}

type EquipClearOfflineCardsRequestDetail struct {
	OfflineCards []string `json:"offlineCards"`
}

func (r *EquipClearOfflineCardsRequest) UnmarshalJSON(data []byte) error {
	type Alias EquipClearOfflineCardsRequest
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
