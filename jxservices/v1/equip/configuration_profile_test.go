package equip

import (
	"testing"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
	"github.com/stretchr/testify/assert"
)

func TestConfigurationProfile(t *testing.T) {
	assert.Implements(t, new(services.Request), &equipConfigurationProfileRequest{})
	assert.Implements(t, new(services.Response), &equipConfigurationProfileResponse{})
}
