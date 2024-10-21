package equip

import (
	"testing"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
	"github.com/stretchr/testify/assert"
)

func TestChargerTerminate(t *testing.T) {
	assert.Implements(t, new(services.Request), &equipChargerTerminateRequest{})
	assert.Implements(t, new(services.Response), &equipChargerTerminateResponse{})
}
