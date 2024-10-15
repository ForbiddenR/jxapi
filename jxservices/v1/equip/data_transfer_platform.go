package equip

import (
	"encoding/json"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type EquipDataTransferRequest struct {
	services.Base
	Data *EquipDataTransferRequestDetail `json:"data"`
}

type EquipDataTransferRequestDetail struct {
	VendorId  string `json:"vendorId"`
	Data      string `json:"data"`
	MessageId string `json:"messageId"`
}

func (r *equipDataTransferRequest) UnmarshalJSON(b []byte) error {
	type Alias equipDataTransferRequest
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}

	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	return nil
}
