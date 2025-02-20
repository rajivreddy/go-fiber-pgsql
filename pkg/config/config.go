package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config is a struct that represents the configuration of the application

type Db struct {
	Host     string `yaml:"host" env:"DB_HOST" env-default:"localhost"`
	Port     string `yaml:"port" env:"DB_PORT" env-default:"5432"`
	Username string `yaml:"username" env:"DB_USERNAME" env-required:"true"`
	Password string `yaml:"password" env:"DB_PASSWORD" env-required:"true"`
	DbName   string `yaml:"dbname" env:"DB_NAME" env-required:"true"`
	SslMode  string `yaml:"sslmode" env:"DB_SSLMODE" env-default:"disable"`
}
type Server struct {
	Port string `yaml:"port" env:"PORT" env-default:"8080"`
	Host string `yaml:"host" env:"HOST" env-default:"localhost"`
}
type Config struct {
	Db     Db     `yaml:"db"`
	Server Server `yaml:"server"`
}

func LoadConfig() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("config", "", "path to the config file")
		flag.Parse()
		configPath = *flags

		if configPath == "" {
			log.Fatal("config file is required")
		}
	}
	//Check file existed or not

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file not found: %v", configPath)
	}

	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		log.Fatalf("Failed to read config file: %v", err.Error())
	}

	return &cfg
}
