package main

import (
	"log"
	"net"
	"time"
)

type RegHand struct {
	NetworkStatus Network
	FlashUse      FlashUse
	FlashMount    bool
}

type FlashUse struct {
	ServiceWork bool
	ProcessWork bool
}

type Network struct {
	ModemNetWorkStatus bool
	VPNNetworkStatus   bool
}

func (n *Network) GeneralNetCheck(host string, port string) bool {

	//host := "google.com"
	//port := "80"
	timeout := 5 * time.Second
	_, generalNeterr := net.DialTimeout("tcp", host+":"+port, timeout)

	if generalNeterr != nil {
		log.Print(host+":"+port, " not responding ", generalNeterr.Error())
		n.ModemNetWorkStatus = false
	} else {
		n.ModemNetWorkStatus = true
	}

	return n.ModemNetWorkStatus

}
