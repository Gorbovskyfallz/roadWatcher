package Config

import (
	"flag"
	"github.com/fsnotify/fsnotify"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	Vpn       Vpn       `yaml:",inline"`
	GlobalNet GlobalNet `yaml:",inline"`
	Security  Security  `yaml:",inline"`
	Flash     Flash     `yaml:",inline"`
	Hardware  Hardware  `yaml:",inline"`
}

type Vpn struct {
	NetAddress   string `yaml:"privateNetwork"`
	PingQty      int    `yaml:"pingTimesForVpn"`
	RebootEnable bool   `yaml:"rebootWhilePrivateFail"`
	RebootTime   int    `yaml:"vpnRebootTimeout"`
}

type GlobalNet struct {
	NetAddress   string `yaml:"globalNetwork"`
	NetPort      int    `yaml:"globalNetworkPort"`
	RebootTime   int    `yaml:"globalRebootTimeout"`
	RebootEnable bool   `yaml:"rebootWhileGlobalFail"`
}
type Security struct {
	BotToken string `yaml:"tokenBotApi"`
}
type Flash struct {
	DevPath    string `yaml:"pathToDevice"`
	MountPoint string `yaml:"pathToMountPoint"`
}

type Hardware struct {
	LedIndication bool `yaml:"ledIndication"`
}

func ParsePathfromFlag() string {
	argFlag := "conf"
	standartPath := "regConfig.yaml"
	usage := "define path to config"
	confPath := flag.String(argFlag, standartPath, usage)
	flag.Parse()
	log.Printf("using log: %s\n", *confPath)
	return *confPath
}

func (f *Config) ParseFromYaml(configFilePath string) (*Config, error) {
	funcName := "ParseFromYaml"
	yamlFile, yamlParseErr := os.ReadFile(configFilePath)
	if yamlParseErr != nil {
		log.Printf("%s: %v\n", funcName, yamlParseErr)
		return nil, yamlParseErr
	}
	unmarshErr := yaml.Unmarshal(yamlFile, f)
	if unmarshErr != nil {
		log.Printf("%s: %v\n", funcName, unmarshErr)
		return nil, unmarshErr
	}
	return f, nil
}

func (f *Config) AddNotifyWatcher(mainPath string) *fsnotify.Watcher {
	name := "AddNotifyWatcher"
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("%s: create watcher: %v\n", name, err)
	}
	//defer watcher.Close()
	// Start listening for events.
	err = watcher.Add(mainPath)
	if err != nil {
		log.Fatalf("%s: add watcher: %v\n", name, err)
	}
	return watcher

}

func (f *Config) CheckUpdate(watcher *fsnotify.Watcher, mainPath string) {
	name := "CheckUpdate"
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op == fsnotify.Write {

				_, parseErr := f.ParseFromYaml(mainPath)
				if parseErr != nil {
					log.Fatalf("%s: %v\n", name, parseErr)
				}
				log.Printf("%s: config updated\n", name)
			}
		case err, ok := <-watcher.Errors:
			if !ok {

			}
			log.Printf("%s: %v\n", name, err)
		}

	}

}
