package main

import (
	"flag"
)

type portlist []string

func (i *portlist) String() string {
	return "representation of ports list"
}

func main() {
	// Command line arguments

	// if the hostname is not provided use localhost

	address := flag.String("h", "127.0.0.1", "target hostname")

	port_range := flag.String("r", "1-1024", "port range")

	port := flag.String("p", "", "port list")

	flag.Parse()

	scanner := NewScanner(*address)

	if *port != "" {
		scanner.ScanPort(*port)
	} else {
		scanner.ScanRange(*port_range)
	}
}
