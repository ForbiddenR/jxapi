package equip

import (
	"fmt"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

func UnspportedCallbackRequest(name, msgID string) error {
	resp, ok := services.FetchFC(name)
	if !ok {
		fmt.Errorf("this feature hasn't implemented. name: %s, id: %s", name, msgID)
	}
	
}
