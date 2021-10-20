package access

import (
	"github.com/chiraponkub/Todo-Go/environment"
)

type Access struct {
	ENV *environment.Properties
}

func Initial(properties *environment.Properties) *Access {
	return &Access{
		ENV: properties,
	}
}
