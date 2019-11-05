package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Nmap: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}

	binary, lookErr := exec.LookPath("/usr/bin/nmap")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"nmap", "-v", "-A", os.Args[1]}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
