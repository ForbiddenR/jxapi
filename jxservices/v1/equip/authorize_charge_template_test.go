package equip

import (
	"testing"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
	"github.com/stretchr/testify/assert"
)

func TestAuthorizeChargeTemplate(t *testing.T) {
	assert.Implements(t, new(services.Request), &equipAuthorizeChargeTemplateRequest{})
	assert.Implements(t, new(services.Response), &equipAuthorizeChargeTemplateResponse{})
}
