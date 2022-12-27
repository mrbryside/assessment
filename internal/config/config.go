package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
)

func Init() {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./internal/config")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Not found local env - Use only the environment variable")
	}
}

type provider struct {
	DbUrl string
	Port  string
}

func NewProvider() provider {
	return provider{
		viper.GetString("database.url"),
		fmt.Sprintf(":%v", viper.GetInt("port")),
	}
}
