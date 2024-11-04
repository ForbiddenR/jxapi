package equip

import (
	"encoding/json"
	"errors"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type EquipQueryOfflineCardsRequest struct {
	services.Base
	Data *EquipQueryOfflineCardsRequestDetail `json:"data"`
}

type EquipQueryOfflineCardsRequestDetail struct {
	OfflineCards []string `json:"offlineCards"`
}

func (r *EquipQueryOfflineCardsRequest) UnmarshalJSON(data []byte) error {
	type Alias EquipQueryOfflineCardsRequest
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
