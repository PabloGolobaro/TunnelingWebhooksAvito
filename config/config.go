package config

import (
	"fmt"
	"github.com/spf13/viper"
	tele "gopkg.in/telebot.v3"
	"gorm.io/gorm"
	"log"
)

var Config botConfig

type botConfig struct {
	// the data source name (DSN) for connecting to the database. required.
	DSN string `mapstructure:"dsn"`
	//admins id
	Admins []string `mapstructure:"admins"`
	// Where to send webhooks
	Recepient []string `mapstructure:"recepient"`
	// bot token
	Token string `mapstructure:"token"`
	// the shared DB ORM object
	DB *gorm.DB
	//shared telebot object
	Bot *tele.Bot
	// the error thrown be GORM when using DB ORM object
	DBErr error
}

func LoadConfig() error {
	log.Println("Loading config...")
	v := viper.New()
	v.SetConfigName("bot_config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}
	log.Println("Loaded config...")
	return v.Unmarshal(&Config)
}
