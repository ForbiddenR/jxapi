package equip

import (
	"testing"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
	"github.com/stretchr/testify/assert"
)

func TestChargingHandshake(t *testing.T) {
	assert.Implements(t, new(services.Request), &equipChargingHandshakeRequest{})
	assert.Implements(t, new(services.Response), &equipChargingHandshakeResponse{})
}
