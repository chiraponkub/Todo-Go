package environment

import (
	"github.com/Netflix/go-env"
	"log"
)

type Flavor string
type URL string

const (
	Develop    Flavor = "DEVELOP"
	Devspace   Flavor = "DEVSPACE"
	Production Flavor = "PRODUCTION"
)

type Properties struct {
	Flavor Flavor `env:"FLAVOR,default=DEVELOP"`

	//GormHost string `env:"GORM_HOST,default=rdbms"`
	//GormHost string `env:"GORM_HOST,default=localhost"`
	GormHost   string `env:"GORM_HOST,default=todolist-rdbms"`
	GormPort   string `env:"GORM_PORT,default=5432"`
	GormNameDB string `env:"GORM_NAME,default=postgres_db"`
	GormUserDB string `env:"GORM_USER,default=postgres"`
	GormPassDB string `env:"GORM_PASS,default=pgpassword"`
}

func Build() *Properties {
	var prop Properties
	if _, err := env.UnmarshalFromEnviron(&prop); err != nil {
		log.Panic(err)
	}
	return &prop
}
