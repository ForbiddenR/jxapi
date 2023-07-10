package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type BMSLimitRequestInterface interface {
	services.Request
	Construct(BMSLimitRequestConfig)
}

type equipBMSLimitRequest struct {
	services.Base
	Data *equipBMSLimitRequestDetail `json:"data"`
}

type equipBMSLimitRequestDetail struct {
	EvseId           *string `json:"evseSerial"`
	ConnectorId      string  `json:"connectorSerial"`
	MaxTemp          uint64  `json:"maxTemp"`
	MaxCurrent       float64 `json:"maxCurrent"`
	MaxVoltage       float64 `json:"maxVoltage"`
	MaxOutputVoltage float64 `json:"maxOutputVoltage"`
	MaxOutputCurrent float64 `json:"maxOutputCurrent"`
	Soc              float64 `json:"soc"`
	Capacity         float64 `json:"capacity"`
	Vin              string  `json:"vin"`
	Prepare          bool    `json:"prepare"`
}

func (equipBMSLimitRequest) GetName() string {
	return services.BMSLimit.String()
}

type BMSLimitRequestConfig struct {
	services.ReusedConfig
	ConnectorId      string
	MaxTemp          uint64
	MaxCurrent       float64
	MaxVoltage       float64
	MaxOutputVoltag  float64
	MaxOutputCurrent float64
	Soc              float64
	Capacity         float64
	Vin              string
	Prepare          bool
}

func NewEquipBMSLimitRequestWithConfig(config *BMSLimitRequestConfig) *equipBMSLimitRequest {
	return &equipBMSLimitRequest{
		Base: services.Base{
			EquipmentSn: config.Sn,
			Protocol:    config.Protocol,
			Category:    services.BMSLimit.FirstUpper(),
			AccessPod:   config.Pod,
			MsgID:       config.MsgID,
		},
		Data: &equipBMSLimitRequestDetail{
			ConnectorId:      config.ConnectorId,
			MaxTemp:          config.MaxTemp,
			MaxCurrent:       config.MaxCurrent,
			MaxVoltage:       config.MaxVoltage,
			MaxOutputVoltage: config.MaxOutputVoltag,
			MaxOutputCurrent: config.MaxOutputCurren,
			Soc:              config.Soc,
			Capacity:         config.Capacity,
			Vin:              config.Vin,
			Prepare:          config.Prepare,
		},
	}
}

func NewEquipBMSLimitRequest(sn string, protocol *services.Protocol, pod, msgID string, connecorId string, maxTemp uint64,
	maxCurrent, maxVoltage, maxOutputVoltage, maxOutputCurrent, soc, capacity float64, vin string, prepare bool) *equipBMSLimitRequest {
	req := &equipBMSLimitRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    protocol,
			Category:    services.BMSLimit.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipBMSLimitRequestDetail{
			ConnectorId:      connecorId,
			MaxTemp:          maxTemp,
			MaxCurrent:       maxCurrent,
			MaxVoltage:       maxVoltage,
			MaxOutputVoltage: maxOutputVoltage,
			MaxOutputCurrent: maxOutputCurrent,
			Soc:              soc,
			Capacity:         capacity,
			Vin:              vin,
			Prepare:          prepare,
		},
	}
	return req
}

var _ services.Response = &equipBMSLimitResponse{}

type equipBMSLimitResponse struct {
	api.Response
}

func (resp *equipBMSLimitResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipBMSLimitResponse) GetMsg() string {
	return resp.Msg
}

func BMSLimitRequest(ctx context.Context, req services.Request) error {
	header := services.GetSimpleHeaderValue(services.BMSLimit)

	url := services.GetSimpleURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipBMSLimitResponse{})
}
