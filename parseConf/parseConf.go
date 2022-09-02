package parseConf

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	VpnSettings       VpnSettings       `yaml:",inline"`
	GlobalNetSettings GlobalNetSettings `yaml:",inline"`
	Security          Security          `yaml:",inline"`
	Flash             Flash             `yaml:",inline"`
	Hardware          Hardware          `yaml:",inline"`
}

type VpnSettings struct {
	PrivateNetwork      string `yaml:"privateNetwork"`
	PingTimesForVpn     int    `yaml:"pingTimesForVpn"`
	RebootOnPrivateFail bool   `yaml:"rebootWhilePrivateFail"`
	VpnRebootTimeout    int    `yaml:"vpnRebootTimeout"`
}

type GlobalNetSettings struct {
	GlobalNetwork       string `yaml:"globalNetwork"`
	GlobalNetWorkPort   int    `yaml:"globalNetworkPort"`
	GlobalRebootTimeout int    `yaml:"globalRebootTimeout"`
	RebootIfFail        bool   `yaml:"rebootWhileGlobalFail"`
}
type Security struct {
	EnableTokenConfigParse bool   `yaml:"enableTokenConfigParse"`
	TokenBotApi            string `yaml:"tokenBotApi"`
}
type Flash struct {
	PathToDev      string `yaml:"pathToDevice"`
	MountPointPath string `yaml:"pathToMountPoint"`
}

type Hardware struct {
	LedIndication bool `yaml:"ledIndication"`
}

func (f *Config) ParseConfig(configFilePath string) (*Config, error) {
	yamlFile, yamlParseErr := os.ReadFile(configFilePath)
	if yamlParseErr != nil {
		log.Println("ParseСonfig (parseConf package): yamlParse:", yamlParseErr)
		return nil, yamlParseErr
	}
	unmarshErr := yaml.Unmarshal(yamlFile, f)
	if unmarshErr != nil {
		log.Println("ParseConfig (parseConf package): unmarshParse:", unmarshErr)
		return nil, unmarshErr
	}
	return f, nil
}

func (f *Config) ParseFromTwoDirs(firstPath, SecondPath string) (*Config, error) {
	_, homeDirErr := f.ParseConfig(firstPath)
	if homeDirErr != nil {
		if errors.Unwrap(homeDirErr).Error() == "no such file or directory" {
			_, etcDirConfigErr := f.ParseConfig(SecondPath)
			if etcDirConfigErr != nil {
				log.Fatalln("ParseFromTwoDirs (parseConf package): no config in /etc/ or home dirs, terminate")
				return nil, etcDirConfigErr
			}

		}
	}
	return f, nil
}

func (f *Config) SwitchTokenInput() (*Config, error) {

	switch {
	case f.Security.EnableTokenConfigParse == true && f.Security.TokenBotApi != "":
		log.Println("mutually exclusive conditions: select only one way to introduce botApiToken: through config" +
			"or CLI")
	case f.Security.EnableTokenConfigParse == true && f.Security.TokenBotApi == "":
		fmt.Println("you selected parsing bot api from CLI, please enter your token in string format:")
		fmt.Scanln(&f.Security.TokenBotApi)
	case f.Security.EnableTokenConfigParse == false && f.Security.TokenBotApi != "":
		log.Println("selected config variant introduce api's token")
	case f.Security.TokenBotApi == "" && f.Security.EnableTokenConfigParse == false:
		log.Println("no way to parse telegram api's token, select on of the methods (CLI or config)")
	}

	return f, nil
}
