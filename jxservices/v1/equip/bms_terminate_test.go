package equip

import (
	"testing"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
	"github.com/stretchr/testify/assert"
)

func TestBMSTerminate(t *testing.T) {
	assert.Implements(t, new(services.Request), &equipBMSTerminateRequest{})
	assert.Implements(t, new(services.Response), &equipBMSTerminateResponse{})
}
