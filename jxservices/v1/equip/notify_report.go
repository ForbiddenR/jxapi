package equip

import (
	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type equipNotifyReportRequest struct {
	services.Base
	Data *equipNotifyReportRequestDetail `json:"data"`
}

type equipNotifyReportRequestDetail struct {
	RequestId  int64        `json:"requestId"`
	TBC        bool         `json:"tbc"`
	ReportData []ReportData `json:"reportData"`
}

type ReportData struct {
	Component         Component         `json:"component"`
	Key               string            `json:"key"`
	VariableAttribute VariableAttribute `json:"variableAttribute"`
}

type equipNotifyReportResponse struct {
	api.Response
}

func NewEquipNotifyReportRequest(sn, pod, msgID string, requestId int64, tbc bool, reportDatas ...ReportData) *equipNotifyReportRequest {
	return &equipNotifyReportRequest{
		Base: services.Base{
			EquipmentSn: sn,
			AccessPod:   pod,
			MsgID:       msgID,
			Category:    services.NotifyReport.FirstUpper(),
		},
		Data: &equipNotifyReportRequestDetail{
			RequestId:  requestId,
			TBC:        tbc,
			ReportData: reportDatas,
		},
	}
}
