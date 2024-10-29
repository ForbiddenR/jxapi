package equip

import (
	"testing"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
	"github.com/stretchr/testify/assert"
)

func TestEquipClearOfflineCardsCallback(t *testing.T) {
	assert.Implements(t, new(services.Request), &equipClearOfflineCardsCallbackRequest{})
}
