package equip

import "github.com/ForbiddenR/jx-api/services"

type EquipRemoteStopTransactionRequest struct {
	services.Base
	Data *EquipRemoteStopTransactionRequestDetail `json:"data"`
}

type EquipRemoteStopTransactionRequestDetail struct {
	TransactionId string `json:"transactionId" validate:"required"`
}
