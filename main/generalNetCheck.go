package main

import (
	"log"
	"net"
	"time"
)

// func checking general (not vpn) network with net.DIalTimeout
func generalNetCheck() error {

	host := "google.com"
	port := "80"
	timeout := time.Duration(5 * time.Second)
	_, generalNeterr := net.DialTimeout("tcp", host+":"+port, timeout)

	if generalNeterr != nil {
		log.Print(host+":"+port, " not responding ", generalNeterr.Error())
		return generalNeterr
	}

	return nil

}
