package main

type Config struct {
	VpnSettings       VpnSettings
	GlobalNetSettings GlobalNetSettings
	Security          Security
	Flash             Flash
	Hardware          Hardware
}

type VpnSettings struct {
	PrivateNetwork      string
	PingTimesForVpn     int
	RebootOnPrivateFail bool
	VpnRebootTimeout    int
}

type GlobalNetSettings struct {
	GlobalNetwork       string
	GlobalNetWorkPort   int
	GlobalRebootTimeout int
}
type Security struct {
	EnableTokenConfigParse bool
	TokenBotApi            string
}
type Flash struct {
	PathToDev      string
	MountPointPath string
}

type Hardware struct {
	Ledindication string
}
