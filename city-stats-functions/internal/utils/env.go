package utils

import (
	"github.com/spf13/viper"
)

var (
	WikipediaApiUrl    string
	WalkscoreBaseUrl   string
	WalkscoreSearchUrl string
	NumbeoUrl          string
)

func init() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("../../.")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	WikipediaApiUrl = viper.GetString("WIKIPEDIA_API_URL")
	WalkscoreBaseUrl = viper.GetString("WALKSCORE_BASE_URL")
	WalkscoreSearchUrl = viper.GetString("WALKSCORE_SEARCH_URL")
	NumbeoUrl = viper.GetString("NUMBEO_URL")
}
