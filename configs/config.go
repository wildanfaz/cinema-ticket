package configs

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	AppAuthor    string
	AppPort      string
	JWTSecretKey []byte
	JWTDuration  time.Duration
	DatabaseURL  string
}

func InitConfig() *Config {
	config := viper.New()
	config.AutomaticEnv()
	config.SetConfigFile("config.json")
	config.AddConfigPath(".")

	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	return &Config{
		AppAuthor:    config.GetString("app.author"),
		AppPort:      config.GetString("app.port"),
		JWTSecretKey: []byte(config.GetString("jwt.secret_key")),
		JWTDuration:  config.GetDuration("jwt.duration"),
		DatabaseURL:  config.GetString("database.postgresql"),
	}
}
