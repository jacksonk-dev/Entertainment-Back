package globals

import (
	"fmt"
	"os"

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

func SetConfig() {
	var envVariables EnvVariables

	viper.SetConfigFile("app.yaml")
	viper.ReadInConfig()

	err := viper.Unmarshal(&envVariables)
	if err != nil {
		fmt.Println("Error: Unable to unmarshal")
	}

	os.Setenv("PORT", envVariables.EnvVars.PORT)
	os.Setenv("CLIENT_ORIGIN", envVariables.EnvVars.CLIENT_ORIGIN)
	os.Setenv("TMDB_API_KEY", envVariables.EnvVars.TMDB_API_KEY)
	os.Setenv("TRAKT_API_KEY", envVariables.EnvVars.TRAKT_API_KEY)
}
