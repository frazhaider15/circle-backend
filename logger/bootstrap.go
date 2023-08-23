package logger

import (
	"io"
	"os"

	"github.com/qisst/ms-nadra-verification/conf"
	logredis "github.com/rogierlommers/logrus-redis-hook"
	log "github.com/sirupsen/logrus"
)

var logger_instance = log.New()

var filepath = "./log/logrus.log"

func SetPath(path string) {

	filepath = path
}

func Init() {

	//getting the credentials in the conf file
	conf := conf.GetConfig()

	//setting the format of the logs to be a JSON one
	logger_instance.SetFormatter(&log.JSONFormatter{
		DataKey:     "musdaq",
		PrettyPrint: true,
	})

	//getting the log level set in the configuration file
	logLevel, err := log.ParseLevel(conf.LogLevel)
	//If the log level in conf file can't be parsed, log level should be the default info level
	if err != nil {
		logLevel = log.InfoLevel
	}
	//setting the log level
	logger_instance.SetLevel(logLevel)

	//If we want to throw logs on the network, like redis
	if conf.LogEnvironment == "network" {
		//we dont want to display logs in the output, just on the network
		logger_instance.SetOutput(io.Discard)

		//logurs-redis hook configuration
		hookConfig := logredis.HookConfig{
			Host:     "localhost",
			Key:      "",
			Format:   "v0",
			App:      "Logrus",
			Port:     6379,
			Hostname: "Kibana",
			DB:       2,
			TTL:      3600,
		}

		//Initializing the hook
		hook, err := logredis.NewHook(hookConfig)

		if err == nil {
			//Adding the hook to the logger_instance
			logger_instance.AddHook(hook)
		} else {
			logger_instance.Errorf("logredis error: %q", err)
		}

	} else if conf.LogEnvironment == "local" { //If we want to throw logs into a local file

		logger_instance.SetOutput(os.Stdout)
		//setting it to a file writer
		file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)

		if err == nil {
			logger_instance.Out = file
		} else {
			logger_instance.Info("Failed to log to file, using default stderr")
		}
		logger_instance.Info("Logrus has been initiated")

	}

}

func Instance() *log.Logger {

	return logger_instance
}
