package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

// ViperConfig ...
type ViperConfig struct {
	*viper.Viper
}

// Configs ...
var Configs *ViperConfig

func init() {
	Configs = readConfig(map[string]interface{}{
		"debug_route":        false,
		"port":               4569,
		"redis_host":         "localhost:6379",
		"asynqmon_host":      "localhost:8080",
		"use_docker_compose": false,
	})
}

func readConfig(defaults map[string]interface{}) *ViperConfig {
	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}

	v.AddConfigPath("./conf")
	v.AutomaticEnv()
	v.SetConfigName(".env.dev")

	err := v.ReadInConfig()
	if err != nil {
		fmt.Println("conf", "readConfig", "Error", err) // TODO: logger
		return nil
	}

	return &ViperConfig{
		Viper: v,
	}
}
