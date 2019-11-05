package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Using: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	target := os.Args[1]
	var wg sync.WaitGroup

	for port := 0; port <= 65535; port++ {
		wg.Add(1)
		go testTCPConnection(target, port, &wg)
	}
	wg.Wait()
}

func testTCPConnection(ip string, port int, wg *sync.WaitGroup) {
	_, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(port), time.Second*15)
	if err == nil {
		fmt.Printf("Port %d: Open\n", port)
	}
	wg.Done()
}
