package main

import (
	"fmt"
	"log"
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (

	// TCP is the network we are scanning
	TCP string = "tcp"
)

type PortScanner struct {
	// The Host we're scanning

	ADDR string
}

// NewScanner creates and returns a PortScanner type
func NewScanner(addr string) *PortScanner {
	return &PortScanner{
		ADDR: addr,
	}
}

// BuildAdress returns the full adress string
func (p *PortScanner) BuildAdress(port string) string {
	// join the hostname with port
	return p.ADDR + ":" + port
}

func (p *PortScanner) ScanPort(port string) {
	// initialize the tcp connection

	conn, err := net.DialTimeout(TCP, p.BuildAdress(port), 1*time.Second)
	//  can't connect -> port is closed
	if err != nil {
		fmt.Printf("PORT %s : closed \n", port)
		return
	}
	defer conn.Close()

	// can connect -> port is open
	fmt.Printf("PORT %s : open \n", port)
}

func (p *PortScanner) ScanRange(ip_range string) {
	re := regexp.MustCompile(`^\d+-\d+$`)
	is_valid := re.MatchString(ip_range)
	if !is_valid {
		log.Fatal("invalid ip range")
	}
	port_range := strings.Split(ip_range, "-")
	start, err := strconv.Atoi(port_range[0])
	if err != nil {
		panic(err)
	}
	end, err := strconv.Atoi(port_range[1])
	if err != nil {
		panic(err)
	}
	for i := start; i <= end; i++ {
		p.ScanPort(fmt.Sprint(i))
	}
}
