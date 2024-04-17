package equip

import (
	"encoding/json"
	"errors"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

type EquipTriggerMessageRequest struct {
	services.Base
	Data *EquipTriggerMessageRequestDetail `json:"data"`
}

type EquipTriggerMessageRequestDetail struct {
	RequestedMessage TriggerMessageEnumType `json:"requestedMessage" validate:"required"`
	ConnectorID      string                 `json:"connectorSerial" validate:"required"`
	EvseID           *string                `json:"evseSerial,omitempty"`
}

func (t *EquipTriggerMessageRequest) UnmarshalJSON(data []byte) error {
	type Alias EquipTriggerMessageRequest
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(t),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if aux.Data == nil {
		return errors.New("data is nil")
	}
	return nil
}

type EquipCallStatusNotificationRequest struct {
	services.Base
	Data *EquipCallStatusNotificationRequestDetail `json:"data"`
}

func (r *EquipCallStatusNotificationRequest) UnmarshalJSON(data []byte) error {
	type Alias EquipCallStatusNotificationRequest
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
		if aux.Data.EvseID == nil {
			return errors.New(aux.Protocol.String() + ":evseSerial is nil")
		}
	}
	return nil
}

type EquipCallStatusNotificationRequestDetail struct {
	ConnectorID string  `json:"connectorSerial" validate:"required"`
	EvseID      *string `json:"evseSerial,omitempty"`
}
