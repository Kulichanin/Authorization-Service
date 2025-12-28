package config

import (
	"flag"
	"log"
	"os"

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
	SSLMode  bool   `mapstructure:"sslmode"`
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

func MustLoad() *Config {
	var AppConfig *Config

	path := FetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}
	viper.SetConfigType("yaml")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist:" + path)
	}

	// Добавляем пути поиска конфига
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: config file not found, using defaults and environment variables: %v", err)
	}

	setDefaults()

	// Привязка конфига к структуре
	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		panic("failed to unmarshal config: ")
	}

	log.Printf("Config loaded successfully from: %s", viper.ConfigFileUsed())
	return AppConfig
}

func FetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config")
	flag.Parse()

	if res != "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
