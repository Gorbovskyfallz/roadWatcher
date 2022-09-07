package reghandler

import (
	"kek/flash"
	"kek/network"
)

type SysInfo struct {
	Network    network.Network
	FlashUse   flash.FlashUse
	FlashMount flash.StatFlash
}

func NewRegHand(netHost, netPort, vpnHost, proc, service string) SysInfo {
	currentInfo := new(SysInfo)
	currentInfo.Network.NetCheck(netHost, netPort)
	currentInfo.Network.VpnCheck(vpnHost)
	currentInfo.FlashUse.CheckPid(proc)
	currentInfo.FlashUse.CheckService(service)

	return *currentInfo
}
