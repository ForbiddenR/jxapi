package equip

import services "github.com/ForbiddenR/jxapi/jxservices"

type EquipRemoteStopTransactionRequest struct {
	services.Base
	Data *EquipRemoteStopTransactionRequestDetail `json:"data"`
}

type EquipRemoteStopTransactionRequestDetail struct {
	TransactionId string `json:"transactionId" validate:"required"`
}
