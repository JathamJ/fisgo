package conf

import (
	"github.com/JathamJ/fisgo/log"
	"github.com/JathamJ/fisgo/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// LoadConf read config from config file.
func LoadConf(filePath string, conf interface{}) {
	if !utils.FileExist(filePath) {
		log.Fatalf("config file %s not exist", filePath)
		return
	}
	if yamlFile, err := ioutil.ReadFile(filePath); err != nil {
		log.Fatalf("config file %s read failed, err: %s", filePath, err.Error())
		return
	} else if err = yaml.Unmarshal(yamlFile, conf); err != nil {
		log.Fatalf("config file %s unmarshal failed, err: %s", filePath, err.Error())
		return
	}
}
