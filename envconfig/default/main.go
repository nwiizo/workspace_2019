package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Specification struct {
	Port       int `envconfig:"PORT" default:"80"`
}


func main() {
	var AppPort Specification
	envconfig.Process("nwiizo", &AppPort)
	fmt.Println(AppPort.Port) // 80
}
