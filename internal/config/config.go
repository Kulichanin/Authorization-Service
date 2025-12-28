package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Env         string `mapstructure:"env"`
	StoragePath string `mapstructure:"storage_path"`
	GRPC        GRPCConfig
	DataBase    DataBaseConfig
}

type GRPCConfig struct {
	Port int    `mapstructure:"port"`
	Host string `masstructure:"host"`
}

type DataBaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

var AppConfig *Config

func Init(pathConfig string) error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Добавляем пути поиска конфига
	viper.AddConfigPath(pathConfig)

	// Настройка дефолтных значений
	setDefaults()

	if err := viper.ReadConfig(); err != nil {
		log.Printf("Warning: config file not found, using defaults and environment variables: %v", err)
	}

	// Привязка конфига к структуре
	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	log.Printf("Config loaded successfully from: %s", viper.ConfigFileUsed())
	return nil
}

func setDefaults() {
	viper.SetDefault("env", "dev")
	viper.SetDefault("storage_path", "./storage/sso.db")

	viper.SetDefault("grpc.port", "44044")
	viper.SetDefault("grpc.timeout", "1h")

	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.sslmode", "disable")
}
