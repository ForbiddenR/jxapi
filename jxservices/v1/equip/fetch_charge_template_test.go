package equip

import (
	"testing"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
	"github.com/stretchr/testify/assert"
)

func TestFetchChargeTemplateCallback(t *testing.T) {
	assert.Implements(t, new(services.Request), &equipFetchChargeTemplateRequest{})
	assert.Implements(t, new(services.Response), &equipFetchChargeTemplateResponse{})
}