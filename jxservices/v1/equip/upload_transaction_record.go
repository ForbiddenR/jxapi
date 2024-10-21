package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi/v2"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipUploadTransactionRecordRequest struct {
	services.Base
}

func (r *equipUploadTransactionRecordRequest) GetName() services.Request2ServicesNameType {
	return services.UploadTransactionRecord
}

func (r *equipUploadTransactionRecordRequest) TraceId() string {
	return r.MsgID
}

func (r *equipUploadTransactionRecordRequest) IsCallback() bool {
	return false
}

func NewUploadTransactionRecordRequest(sn, pod, msgId string, p *services.Protocol) *equipUploadTransactionRecordRequest {
	return &equipUploadTransactionRecordRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.UploadTransactionRecord.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
	}
}

type equipUploadTransactionRecordResponse struct {
	api.Response
}

func (r *equipUploadTransactionRecordResponse) GetMsg() string {
	return r.Msg
}

func (r *equipUploadTransactionRecordResponse) GetStatus() int {
	return r.Status
}

func UplaodTransactionRecordRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
