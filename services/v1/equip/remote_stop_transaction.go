package equip

import "gitee.com/csms/jxeu-ocpp/pkg/api/services"

type EquipRemoteStopTransactionRequest struct {
	services.Base
	Data *EquipRemoteStopTransactionRequestDetail `json:"data"`
}

type EquipRemoteStopTransactionRequestDetail struct {
	TransactionId string `json:"transactionId" validate:"required"`
}
