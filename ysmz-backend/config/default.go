package config

import (
	"log"

	"github.com/spf13/viper"
)

type ConfigInterface interface {
	Config() *ConfigStruct
}

type ConfigStruct struct {
	ClientOrigin            string `mapstructure:"CLIENT_ORIGIN"`
	OAuthGoogleClientID     string `mapstructure:"OAUTH_GOOGLE_CLIENT_ID"`
	OAuthGoogleClientSecret string `mapstructure:"OAUTH_GOOGLE_CLIENT_SECRET"`
	OAuthGoogleRedirectURL  string `mapstructure:"OAUTH_GOOGLE_REDIRECT_URL"`
}

var config *ConfigStruct // singleton

func Config() *ConfigStruct {
	if config == nil {
		config := new(ConfigStruct)

		viper.AddConfigPath(".")
		viper.SetConfigType("env")
		viper.SetConfigName("app")
		viper.AutomaticEnv()

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("Error while reading config file %s", err)
		}

		err = viper.Unmarshal(config)
		return config
	}

	return config
}
