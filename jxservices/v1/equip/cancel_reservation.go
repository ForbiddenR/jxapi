package equip

import (
	"encoding/json"
	"errors"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

type EquipCancelReservationRequest struct {
	services.Base
	Data *EquipCancelReservationRequestDetail `json:"data"`
}

func (r *EquipCancelReservationRequest) UnmarshalJSON(data []byte) error {
	type Alias EquipCancelReservationRequest
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

type EquipCancelReservationRequestDetail struct {
	ID int64 `json:"id" validate:"required"`
}
