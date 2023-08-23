package conf

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

var (
	configPath     = "."
	configFileName = "conf.json"
)

type GbeConfig struct {
	Mysql          MysqlConfig      `json:"mysqlConfig"`
	RestServer     RestServerConfig `json:"restServer"`
	LogLevel       string           `json:"logLevel"`
	LogEnvironment string           `json:"logEnvironment"`
	TempDir        TempDir
}

type MysqlConfig struct {
	Host       string `json:"host"`
	Port       string `json:"port"`
	DbName     string `json:"dbName"`
	DbUserName string `json:"dbUserName"`
	DbPassword string `json:"dbPassword"`
}
type TempDir struct {
	Path string
}

type RestServerConfig struct {
	Addr string `json:"addr"`
}

var config GbeConfig
var configOnce sync.Once

func SetConfFilePath(path string) {
	configPath = path
}

func SetConfFileName(name string) {
	configFileName = name
}

func GetConfig() *GbeConfig {
	configOnce.Do(func() {
		bytes, err := os.ReadFile(configPath + "/" + configFileName)
		log.Println(configPath + "/" + configFileName)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &config)
		if err != nil {
			panic(err)
		}
	})
	return &config
}
