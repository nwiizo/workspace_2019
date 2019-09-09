package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	Gopath string // 先頭大文字
}

func main() {
	var goenv Env
	envconfig.Process("", &goenv)
	fmt.Println(goenv) 
}
