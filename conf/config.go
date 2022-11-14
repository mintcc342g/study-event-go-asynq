package conf

import (
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
		"debug_route":   false,
		"debug_db":      true,
		"port":          4569,
		"redis_host":    "localhost:6379",
		"asynqmon_host": "localhost:8080",
		// "broker_host": "redis://localhost:6379",
		// "db":          "mysql",
		// "db_host":     "localhost",
		// "db_port":     3306,
		// "db_name":     "asynq",
		// "db_user":     "root",
		// "db_pass":     "",
	})
}

func readConfig(defaults map[string]interface{}) *ViperConfig {
	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}

	return &ViperConfig{
		Viper: v,
	}
}
