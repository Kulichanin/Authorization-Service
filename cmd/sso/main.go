/*
Copyright © 2025 Dmitry mr.gn0m@yandex.ru
*/
package main

import (
	"log"

	"github.com/Kulichanin/Authorization-Service/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra-cli/cmd"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "Моё gRPC приложение",
	Long:  `Полноценное gRPC приложение с CLI интерфейсом`,
}

func main() {
	//  TODO: работа с конфигами через viper

	//  TODO: работа с логами через slog

	//  TODO: инициализация приложения через cobra

	// Инициализация конфигурации
	if err := config.Init(""); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Использование конфигурации
	log.Printf("Starting gRPC server on port: %s", config.AppConfig.Env)
	cmd.Execute()
}
