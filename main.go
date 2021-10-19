package main

import (
	"ProjectEcho/access"
	"ProjectEcho/control"
	"ProjectEcho/environment"
	"ProjectEcho/present"
	"log"
	"os"
	"os/signal"
)

func main() {

	// -- Build Environment
	prop := environment.Build()
	if prop == nil {
		log.Panic("environment not exist")
	}

	// Init Access
	acc := access.Initial(prop)

	sign := make(chan os.Signal, 1)
	signal.Notify(sign, os.Interrupt, os.Kill)

	present.APICreate(control.APICreate(acc))

	<-sign
}