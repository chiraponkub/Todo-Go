package control

import (
	"github.com/chiraponkub/Todo-Go/access"
	"sync"
)

var (
	apiControl APIControl
	apiOnce    sync.Once
)

type APIControl struct {
	access *access.Access
}

func APICreate(access *access.Access) *APIControl {
	apiOnce.Do(func() {
		apiControl = APIControl{access: access}
	})
	return &apiControl
}
