package equip

import (
	"testing"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
	"github.com/stretchr/testify/assert"
)

func TestSetChargeTemplateCallback(t *testing.T) {
	assert.Implements(t, new(services.Request), &equipSetChargeTemplateCallbackRequest{})
	assert.Implements(t, new(services.Response), &equipSetChargeTemplateCallbackResponse{})
}
