package reghandler

import (
	"kek/flash"
	"kek/network"
)

type RegHand struct {
	NetworkStatus network.Network
	FlashUse      flash.FlashUse
	FlashMount    flash.Flash
}

func NewRegHand() (completeRegHand RegHand) {
	completeRegHand.NetworkStatus.ModemNetCheck("google.com", "80")
	completeRegHand.NetworkStatus.VpnNetCheck("10.0.0.1")
	completeRegHand.FlashUse.CheckPid("ffmpeg")
	completeRegHand.FlashUse.CheckService("rtsp-simple-service.service")

	return completeRegHand
}
