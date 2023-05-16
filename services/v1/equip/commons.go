package equip

import (
	ocpp16 "gitee.com/csms/jxeu-ocpp/pkg/ocpp1.6/protocol"
)

type IdTokenStatusTypeEnum int

const (
	Accepted IdTokenStatusTypeEnum = iota
	Blocked
	Expired
	Invalid
	ConcurrentTx
)

func OCPP16AuthorizeRespStatusType(s IdTokenStatusTypeEnum) ocpp16.ResponseJsonIdTagInfoStatus {
	var status ocpp16.ResponseJsonIdTagInfoStatus
	switch s {
	case Accepted:
		status = ocpp16.ResponseJsonIdTagInfoStatusAccepted
	case Blocked:
		status = ocpp16.ResponseJsonIdTagInfoStatusBlocked
	case Expired:
		status = ocpp16.ResponseJsonIdTagInfoStatusExpired
	case Invalid:
		status = ocpp16.ResponseJsonIdTagInfoStatusInvalid
	case ConcurrentTx:
		status = ocpp16.ResponseJsonIdTagInfoStatusConcurrentTx
	}
	return status
}

type Component struct {
	Name string `json:"name"`
	Evse *EVSE  `json:"evse,omitempty"`
}

type EVSE struct {
	Id          string `json:"serial"`
	ConnectorId string `json:"connectorSerial"`
}

type Intellect struct {
	ID        int64  `json:"id"`
	LimitSoC  *int   `json:"limitSoC,omitempty"`
	LimitElec *int   `json:"limitElectricity,omitempty"`
	StopTime  *int64 `json:"stopTime,omitempty"`
}

type IdTokenType struct {
	IdToken string           `json:"idToken"`
	Type    *IdTokenTypeEnum `json:"type,omitempty"`
}

type IdTokenTypeEnum int

type IdTokenInfo struct {
	ExpiryDate    *int64                `json:"expiryDate,omitempty"`
	ParentIdToken *IdTokenType          `json:"parentIdToken,omitempty"`
	GroupIdToken  *IdTokenType          `json:"groupIdToken,omitempty"`
	Status        IdTokenStatusTypeEnum `json:"status"`
}

const (
	IdTokenTypeCentral IdTokenTypeEnum = iota
	IdTokenTypeRFID
	IdTokenTypeBluetooth
	IdTokenTypeNFC
	IdTokenTypeVIN
	IdTokenTypeAPP
	IdTokenTypeHLHT
	IdTokenTypeeMAID
	IdTokenTypeISO14443
	IdTokenTypeISO15693
	IdTokenTypeKeyCode
	IdTokenTypeLocal
	IdTokenTypeMacAddress
	IdTokenTypeNoAuthorization
)

type VariableAttribute struct {
	Value      string     `json:"value"`
	Mutability Mutability `json:"mutability"`
}

type Mutability int

const (
	MutabilityReadOnly Mutability = iota
	MutabilityReadWrite
)
