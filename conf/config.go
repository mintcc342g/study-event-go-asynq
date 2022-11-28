package conf

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	TableDefaultCharset = "DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci"
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
		"db_user":            "root",
		"db_pass":            "",
		"db_host":            "localhost:3306",
		"db_name":            "asynq",
	})
}

func readConfig(defaults map[string]interface{}) *ViperConfig {
	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}

	v.AddConfigPath(".././conf")
	v.AddConfigPath("./conf")

	v.AutomaticEnv()
	v.SetConfigName(".env.dev")

	err := v.ReadInConfig()
	if err != nil {
		zap.S().Errorw("fail to read configs", "err", err)
		return nil
	}

	return &ViperConfig{
		Viper: v,
	}
}
