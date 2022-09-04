package parseConf

import (
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
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
	//log.Println("new config loaded")
	return f, nil
}

func (f *Config) SwitchTokenInput() (*Config, error) {

	switch {
	case f.Security.EnableTokenConfigParse == true && f.Security.TokenBotApi != "":
		fmt.Println("mutually exclusive conditions: select only one way to introduce botApiToken -  through config " +
			"or CLI, deleting info about token in config. Please, enter api token from keyboard:")
		fmt.Scanln(&f.Security.TokenBotApi)
		fmt.Println("thanks o lot")
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

func (f *Config) ConfigNotifier(firstPath, secondPath string) {
	_, parseErr := f.ParseFromTwoDirs("regConfig.yaml", "")
	if parseErr != nil {
		log.Fatal("cannot parse the config:", parseErr)
	}
	_, _ = f.SwitchTokenInput()
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("NewWatcher failed: ", err)
	}
	defer func(watcher *fsnotify.Watcher) {
		errWatcherClose := watcher.Close()
		if errWatcherClose != nil {
			log.Println("cannot defer close watcher:", errWatcherClose)
		}
	}(watcher)

	done := make(chan bool)
	go func() {
		defer close(done)

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				//log.Printf("%s %s\n", event.Name, event.Op)
				if event.Op == fsnotify.Write {
					log.Println("config changed")
					_, notiParseErr := f.ParseFromTwoDirs(firstPath, secondPath)
					if notiParseErr != nil {
						log.Println("cannot parse config after changing it by user:", notiParseErr)
					}
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}

	}()

	err = watcher.Add("./regConfig.yaml")
	if err != nil {
		log.Fatal("Add failed:", err)
	}
	<-done
}
