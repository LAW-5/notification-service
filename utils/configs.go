package utils

import (
	"log"

	"github.com/spf13/viper"
)

var ApiConfig Config

type Config struct {
	AMQPServerURL	string	`mapstructure:"AMQP_SERVER_URL"`
	DBHost			string	`mapstructure:"DB_HOST"`
	DBPort			string	`mapstructure:"DB_PORT"`
	DBUser			string	`mapstructure:"DB_USER"`
	DBPassword		string	`mapstructure:"DB_PASSWORD"`
	DBName			string	`mapstructure:"DB_NAME"`
}

func LoadConfig() {
	viper.AddConfigPath("../")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = viper.Unmarshal(&ApiConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("load environment variable")
}