package equip

import (
	"encoding/json"
	"errors"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

type EquipUpdateFirmwareRequest struct {
	services.Base
	Data *EquipUpdateFirmwareRequestDetail `json:"data"`
}

type EquipUpdateFirmwareRequestDetail struct {
	Location        string   `json:"location"`
	RequestID       *int64   `json:"requestId,omitempty"`
	Retries         *int     `json:"retries,omitempty"`
	RetryInterval   *int     `json:"retryInterval,omitempty"`
	InstallDateTime *int64   `json:"installDateTime,omitempty"`
	Params          []string `json:"params,omitempty"`
}

func (r *EquipUpdateFirmwareRequest) UnmarshalJSON(b []byte) error {
	type Alias EquipUpdateFirmwareRequest
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	if aux.Data == nil {
		return errors.New("data is nil")
	}
	if !aux.Protocol.Equal(services.OCPP16()) {
		if aux.Data.RequestID == nil {
			return errors.New(aux.Protocol.String() + ":request id is nil")
		}
	}

	return nil
}
