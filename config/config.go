package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/joeshaw/envdecode"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port int `env:"PORT,default=8080"`
}

type DatabaseConfig struct {
	Username string `env:"DB_USERNAME,default=ecom"`
	Password string `env:"DB_PASSWORD,default="`
	Host     string `env:"DB_HOST,default=localhost"`
	Port     int    `env:"DB_PORT,default=5432"`
}

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

func LoadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AllowEmptyEnv(true)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config File not found")
		} else {
			log.Fatal(err)
		}
	}

	var config Config

	if err := envdecode.Decode(&config); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("Unable to decode config into struct, %v\n", err)
	}

	fmt.Printf("config: %+v\n", config)
}