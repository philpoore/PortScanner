package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

const maxPort = 65535
const requestDelay = time.Microsecond * 100
const requestTimeout = time.Second * 3

func usage() {
	fmt.Printf("PortScanner go v0.0.1\n\n")
	fmt.Printf("Usage: portscanner <host|ip>\n\n")
	fmt.Printf("Examples: portscanner google.com\n")
	fmt.Printf("          portscanner 127.0.0.1\n")
}

func checkPort(done chan<- bool, host string, port int) {
	address := string(host + ":" + strconv.Itoa(port))
	conn, err := net.DialTimeout("tcp", address, requestTimeout)

	if err != nil {
		done <- false
		return
	}

	fmt.Printf("[+] Port %d open\n", port)
	conn.Close()
	done <- true
}

func main() {
	if len(os.Args) == 1 {
		usage()
		os.Exit(1)
	}

	host := os.Args[1]
	done := make(chan bool)
	activeRequests := 0
	stats := statsCounter{}

	fmt.Printf("Starting port scan of %v\n", host)

	for port := 1; port <= maxPort; port++ {
		activeRequests++
		go checkPort(done, host, port)
		time.Sleep(requestDelay)
	}
	for {
		res := <-done
		stats.update(res)
		activeRequests--
		if activeRequests == 0 {
			break
		}
	}

	stats.display()
}
