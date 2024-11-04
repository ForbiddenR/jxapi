package equip

import (
	"context"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipDataTransferRequest struct {
	services.Base
	Callback *equipDataTransferRequestDetail `json:"data"`
}

type equipDataTransferRequestDetail struct {
}

func (r equipDataTransferRequest) GetName() services.Request2ServicesNameType {
	return services.DataTransfer
}

func (r *equipDataTransferRequest) TraceId() string {
	return r.MsgID
}

func (equipDataTransferRequest) IsCallback() bool {
	return true
}

func NewEquipDataTransferRequest() *equipDataTransferRequest {
	req := &equipDataTransferRequest{}
	return req
}

// type equipDataTransferResponse struct {
// 	api.Response
// 	Data *equipDataTransferResponseDetail `json:"data"`
// }

// type equipDataTransferResponseDetail struct {
// }

// func (resp *equipDataTransferResponse) GetStatus() int {
// 	return resp.Status
// }

// func (resp *equipDataTransferResponse) GetMsg() string {
// 	return resp.Msg
// }

func DataTransferRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
