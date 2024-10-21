package equip

import (
	"testing"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
	"github.com/stretchr/testify/assert"
)

func TestErrorReport(t *testing.T) {
	assert.Implements(t, new(services.Request), &equipErrorReportRequest{})
	assert.Implements(t, new(services.Response), &equipErrorReportResponse{})
}
