package parseConf

import (
	"errors"
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

func (f *Config) ParseConfig(configFilePath string) (*Config, error) {
	funcName := "ParseConfig"
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

func (f *Config) ParseTwoDirs(firstPath, SecondPath string) (*Config, error) {
	_, homeDirErr := f.ParseConfig(firstPath)
	if homeDirErr != nil {
		if errors.Unwrap(homeDirErr).Error() == "no such file or directory" {
			_, etcDirConfigErr := f.ParseConfig(SecondPath)
			if etcDirConfigErr != nil {
				log.Fatalf("ParseTwoDirs: %v\n", etcDirConfigErr)
				return nil, etcDirConfigErr
			}

		}
	}
	//log.Printf("config loaded\n")
	return f, nil
}

//тут должна быть функция нотифаера!!!

func (f *Config) ConfWatcher(mainPath, secPath string) {
	name := "ConfWatcher"
	_, parseErr := f.ParseTwoDirs(mainPath, secPath)
	if parseErr != nil {
		log.Fatalf("%s: %v\n", name, parseErr)
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	// Start listening for events.
	err = watcher.Add(mainPath)
	if err != nil {
		log.Fatal(err)
	}
	go func() {

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op == fsnotify.Write {

					_, parseErr := f.ParseTwoDirs(mainPath, secPath)
					if parseErr != nil {
						log.Fatalf("%s: %v\n", name, parseErr)
					}
					log.Printf("%s: config updated\n", name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Printf("%s: %v\n", name, err)
			}

		}

	}()

	// Add a path.

	// Block main goroutine forever.
	//<-make(chan struct{})

}
