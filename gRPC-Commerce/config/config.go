package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port           string `mapstructure:"PORT"`
	AuthSvcUrl     string `mapstructure:"AUTH_SVC_URL"`
	ProductSvcUrl  string `mapstructure:"PRODUCT_SVC_URL"`
	OrderSvcUrl    string `mapstructure:"ORDER_SVC_URL"`
	PaymentSvcUrl  string `mapstructure:"PAYMENT_SVC_URL"`
	TransferSvcUrl string `mapstructure:"TRANSFER_SVC_URL"`
	DBURL          string `mapstructure:"DB_URL"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Println("Error on Read In Config", err)
		return
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		log.Println("Unmarshall Error", err)
		return
	}
	return c, nil
}
