package main

import (
	"fmt"
	"testing"
)

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

	expectedConfig.Flash.MountPointPath = "/media/passed3"
	expectedConfig.Flash.PathToDev = "/dev/sda"

	expectedConfig.Security.EnableTokenConfigParse = true
	expectedConfig.Security.TokenBotApi = "fff"

	expectedConfig.Hardware.LedIndication = true

	receivedConfig := new(Config)
	receivedConfig.ParseConfig("/home/passed3/GolandProjects/rpi-registartor/main/regConfig.yaml")

	if expectedConfig.Flash != receivedConfig.Flash {
		t.Errorf("flash Section does not match")
	}
	if expectedConfig.VpnSettings != receivedConfig.VpnSettings {
		t.Errorf("vpn setting section does not match")
	}
	if expectedConfig.GlobalNetSettings != receivedConfig.GlobalNetSettings {
		t.Errorf("global network section does not match")
	}
	if expectedConfig.Security != receivedConfig.Security {
		t.Errorf("security section does not match")
	}
	if expectedConfig.Hardware != receivedConfig.Hardware {
		t.Error("Hardware section does not match")
	}

	fmt.Println(*expectedConfig)
	fmt.Println(*receivedConfig)

}
