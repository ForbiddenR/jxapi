package equip

import (
	"context"

	services "github.com/ForbiddenR/jxapi/jxservices"
	"github.com/Kotodian/gokit/datasource/rabbitmq"
)

const updateTransactionQueue = services.QueuePrefix + "transaction"

type equipUpdateTransactionRequest struct {
	services.Base
	Data *equipUpdateTransactionReqeustDetail `json:"data"`
}

func (equipUpdateTransactionRequest) GetName() string {
	return services.UpdateTransaction.String()
}

type equipUpdateTransactionReqeustDetail struct {
}

func NewUpdateTransactionRequest(sn, pod, msgID string, p *services.Protocol) *equipUpdateTransactionRequest {
	updateTransaction := &equipUpdateTransactionRequest{}

	return updateTransaction
}

func UpdateTransactionReqeust(req *services.Request) error {
	ctx := context.Background()
	err := rabbitmq.Publish(ctx, updateTransactionQueue, nil, req)
	if err != nil {
		return err
	}
	return nil
}
