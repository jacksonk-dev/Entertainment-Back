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

func GetConfig() Config {
	var config Config

	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("Error: Unable to unmarshal")
	}

	return config
}
