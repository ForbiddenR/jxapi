package equip

import (
	"context"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

type equipUpdatedParaReportRequest struct {
	services.Base
	Data *equipUpdatedParaReportRequestDetail `json:"data"`
}

type equipUpdatedParaReportRequestDetail struct {
	AuthFree           uint8    `json:"authFree"`
	BtFastAuth         uint8    `json:"btFastAuth"`
	OutputType         string   `json:"outputType"`
	OutputLimitCurrent float64  `json:"outputLimitCurrent"`
	Entities           []Entity `json:"entities"`
}

type Entity struct {
	Status      uint8  `json:"status"`
	Cycled      bool   `json:"cycled"`
	StartTime   string `json:"startTime"`
	StopTime    string `json:"stopTime"`
	IntellectId int64  `json:"intellectId"`
}

func (equipUpdatedParaReportRequest) GetName() services.Request2ServicesNameType {
	return services.UpdatedParaReport
}

func (e *equipUpdatedParaReportRequest) TraceId() string {
	return e.MsgID
}

func (equipUpdatedParaReportRequest) IsCallback() bool {
	return false
}

func NewUpdatedParaReportRequest(sn, pod, msgID string, p *services.Protocol) *equipUpdatedParaReportRequest {
	return &equipUpdatedParaReportRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			Category:    services.UpdatedParaReport.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipUpdatedParaReportRequestDetail{
			Entities: make([]Entity, 2),
		},
	}
}

func UpdatedParaReport(ctx context.Context, req *equipUpdatedParaReportRequest) error {
	return services.Transport(ctx, req)
}