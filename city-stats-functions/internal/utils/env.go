package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	WikipediaApiUrl string
)

func init() {
	fmt.Println("HELLO FROM ENV")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.ReadInConfig()
	fmt.Println(viper.AllSettings())

	WikipediaApiUrl = viper.GetString("WIKIPEDIA_API_URL")
	fmt.Println("URL", WikipediaApiUrl)
}
