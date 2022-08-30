package main

import "testing"

func TestConfig_ParseConfig(t *testing.T) {
	expectedConfig := new(Config)
	expectedConfig.GlobalNetSettings.GlobalNetwork = "8.8.8.8"
	expectedConfig.GlobalNetSettings.GlobalNetWorkPort = 80
	expectedConfig.GlobalNetSettings.GlobalRebootTimeout = 300
	expectedConfig.GlobalNetSettings.RebootIfFail = true

	expectedConfig.VpnSettings.PrivateNetwork = "10.0.0.1"
	expectedConfig.VpnSettings.RebootOnPrivateFail = true
	expectedConfig.VpnSettings.PingTimesForVpn = 3
	expectedConfig.VpnSettings.VpnRebootTimeout = 300

	expectedConfig.Flash.MountPointPath = "dev/passed3"
	expectedConfig.Flash.PathToDev = "dev/sda"

	expectedConfig.Security.EnableTokenConfigParse = true
	expectedConfig.Security.TokenBotApi = "fff"

	expectedConfig.Hardware.Ledindication = true

	receivedConfig := new(Config)
	receivedConfig.ParseConfig()

	if receivedConfig != expectedConfig {
		t.Errorf("config're not pass")
	}

}
