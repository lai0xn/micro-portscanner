package main

import "net"

type PortScanner struct {
	ADDR string
	PORT string
}

func NewScanner(addr string, port string) *PortScanner {
	return &PortScanner{
		ADDR: addr,
		PORT: port,
	}
}
