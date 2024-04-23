package main

import (
	"github.com/rs/zerolog/log"

	"github.com/iliyankg/colab-shield/cli/cmd"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config.json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read config from working directory. Make sure config.json is present.")
	}

	cmd.Execute()
}
