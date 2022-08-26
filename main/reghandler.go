package main

type RegHand struct {
	NetworkStatus Network
	FlashUse      FlashUse
	FlashMount    FlashMount
}

func NewRegHand() (completeRegHand RegHand) {
	completeRegHand.NetworkStatus.ModemNetCheck("google.com", "80")
	completeRegHand.NetworkStatus.VpnNetCheck("10.0.0.1")
	completeRegHand.FlashUse.CheckPid("ffmpeg")
	completeRegHand.FlashUse.CheckService("rtsp-simple-service.service")

	return completeRegHand
}
