package jxesam

import "strings"

const (
	Equip     = "device"
	TicketKey = "ServiceInternalTickets"
)

type RequestNameEsamType string

const (
	Access RequestNameEsamType = "accessVerify"
)

func (r RequestNameEsamType) String() string {
	return string(r)
}

func (r RequestNameEsamType) Split() []string {
	for i := 0; i < len(r.String()); i++ {
		str := r.String()[i : i+1]
		if str == strings.ToUpper(str) {
			return []string{r.String()[:i], strings.ToLower(r.String()[i:])}
		}
	}
	return nil
}
