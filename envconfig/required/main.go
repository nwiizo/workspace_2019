package main

import (
	"fmt"
)

type Env struct {
	Moto int `required:"true"` // $HOGEは未定義
}

func main() {
	var goenv Env
	if err := envconfig.Process("", &goenv); err != nil {
		log.Printf("[ERROR] Failed to process env: %s", err)
	}
}
