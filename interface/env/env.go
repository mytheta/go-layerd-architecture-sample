package env

import (
	"log"
	"runtime"

	"github.com/BurntSushi/toml"
)

var conf Config

type (
	// Config is all config
	Config struct {
		App AppConfig
	}

	// AppConfig is app config
	AppConfig struct {
		Port string
	}
)

// Setup is env setup
func Setup(file string) error {
	runtime.GOMAXPROCS(runtime.NumCPU())
	_, err := toml.DecodeFile(file, &conf)
	if err != nil {
		// Error Handling
		log.Panic(err)
		return err
	}

	return nil
}

// GetAppConfig is get app config
func GetAppConfig() *AppConfig {
	return &conf.App
}
