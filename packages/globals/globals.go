package globals

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	PORT          string
	TMDB_API_KEY  string
	TRAKT_API_KEY string
	CLIENT_ORIGIN string
}

type EnvVariables struct {
	EnvVars Config `mapstructure:"env_variables"`
}

func GetConfig() Config {
	var envVariables EnvVariables

	viper.SetConfigFile("app.yaml")
	viper.ReadInConfig()

	err := viper.Unmarshal(&envVariables)
	if err != nil {
		fmt.Println("Error: Unable to unmarshal")
	}

	return envVariables.EnvVars
}
