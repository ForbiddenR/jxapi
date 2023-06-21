package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type BatteryType uint8

const (
	BatteryTypeLeadAcidBattery BatteryType = iota
	BatteryTypeNickelCadmiumBattery
	BatteryTypePhosphorusSodiumBattery
	BatteryTypePhosphorusIronBattery
	BatteryTypeGalliumArsenideBattery
	BatteryTypeTripleMaterialBattery
	BatteryTypePolymerLiIonBattery
	BatteryTypeTitaniumArsenideBattery
	BatteryTypeOtherBattery
)

type equipBMSInfoRequest struct {
	services.Base
	Data *equipBMSInfoRequestDetail `json:"data"`
}

type equipBMSInfoRequestDetail struct {
	EvseId       *string `json:"evseSerial"`
	ConnectorId  string  `json:"connectorSerial"`
	LimitVoltage float64 `json:"limitVoltage"`
	Version      string  `json:"version"`
	Type         uint8   `json:"type"`
	TotalVoltage float64 `json:"totalVoltage"`
	TotalEnergy  float64 `json:"totalEnergy"`
	Vin          string  `json:"vin"`
}

func (equipBMSInfoRequest) GetName() string {
	return services.BMSInfo.String()
}

func NewEquipBMSInfoRequest(sn string, protocol *services.Protocol, pod, msgID string, connectorId string, limitVoltage float64, version string, ty uint8, totalVoltage, totalEnergy float64, vin string) *equipBMSInfoRequest {
	req := &equipBMSInfoRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    protocol,
			Category:    services.BMSInfo.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipBMSInfoRequestDetail{
			ConnectorId:  connectorId,
			LimitVoltage: limitVoltage,
			Version:      version,
			Type:         ty,
			TotalVoltage: totalVoltage,
			TotalEnergy:  totalEnergy,
			Vin:          vin,
		},
	}
	return req
}

var _ services.Response = &equipBMSInfoResponse{}

type equipBMSInfoResponse struct {
	api.Response
}

func (resp *equipBMSInfoResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipBMSInfoResponse) GetMsg() string {
	return resp.Msg
}

func BMSInfoRequest(ctx context.Context, req services.Request) error {
	header := services.GetSimpleHeaderValue(services.BMSInfo)

	url := services.GetSimpleURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipBMSInfoResponse{})
}
