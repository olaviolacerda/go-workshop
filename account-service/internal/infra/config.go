package infra

import (
	"log"

	"github.com/spf13/viper"
)

var (
	cfg      *viper.Viper
	isLoaded bool = false
)

func _loadConfigFile() {
	log.Println("[INFO] Loading configuration file.")
	cfg = viper.New()
	cfg.AddConfigPath("./config")
	cfg.SetConfigName("server")
	cfg.SetConfigType("json")
}

// GetConfig read a key configuration in 'server.json' file
func GetConfig(key string) string {
	if isLoaded {
		return cfg.Get(key).(string)
	}

	_loadConfigFile()

	if err := cfg.ReadInConfig(); err != nil {
		log.Printf("[ERROR] Fail to read configuration file: %s \n", err.Error())
		return ""
	}

	isLoaded = true
	return cfg.Get(key).(string)
}
