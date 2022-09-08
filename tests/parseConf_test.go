package tests

import (
	"fmt"
	"kek/Config"
	"testing"
)

func TestConfig_ParseConfig(t *testing.T) {
	expectedConfig := new(Config.Config)
	expectedConfig.GlobalNet.NetAddress = "8.8.8.8"
	expectedConfig.GlobalNet.NetPort = 80
	expectedConfig.GlobalNet.RebootTime = 300
	expectedConfig.GlobalNet.RebootEnable = true

	expectedConfig.Vpn.NetAddress = "10.0.0.1"
	expectedConfig.Vpn.RebootEnable = true
	expectedConfig.Vpn.PingQty = 3
	expectedConfig.Vpn.RebootTime = 300

	expectedConfig.Flash.MountPoint = "/media/passed3"
	expectedConfig.Flash.DevPath = "/dev/sda"

	expectedConfig.Security.CliTokenParse = true
	expectedConfig.Security.BotToken = "fff"

	expectedConfig.Hardware.LedIndication = true

	receivedConfig := new(Config.Config)
	receivedConfig.ParseFromYaml("/home/passed3/GolandProjects/rpi-registartor/main/regConfig.yaml")

	if expectedConfig.Flash != receivedConfig.Flash {
		t.Errorf("flash Section does not match")
	}
	if expectedConfig.Vpn != receivedConfig.Vpn {
		t.Errorf("vpn setting section does not match")
	}
	if expectedConfig.GlobalNet != receivedConfig.GlobalNet {
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

func TestConfig_ParseFromTwoDirs(t *testing.T) {
	testConfig := new(Config.Config)
	_, err := testConfig.ParseTwoDirs("regConfig.yaml", "/etc/")
	if err != nil {
		t.Errorf("expected nil, but received \"%v\"", err)
	}
}
func TestSwitchTokenInput(t *testing.T) {

	testConf := new(Config.Config)
	// input by cli
	testConf.Security.CliTokenParse = true
	testConf.Security.BotToken = ""
	testConf.SwitchTokenInput()
	if testConf.Security.BotToken != "qwerty" {
		t.Errorf("expected \"qwerty\", but received \"%s\"", testConf.Security.BotToken)
	}

}
