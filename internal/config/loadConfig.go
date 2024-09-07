package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var (
	Config = new(config)
)

func Load() error {
	// Загружаем env файл
	if err := godotenv.Load("./configs/.env"); err != nil {
		return err
	}
	// Получаем путь к кофигурационному файлу
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		return errors.New("variable CONFIG_PATH is empty")
	}
	// Проверяем существует ли кофигурационный файл
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		strErr := fmt.Sprintf("the configuration file is on the path: %s - not found", configPath)
		return errors.New(strErr)
	}
	// Загружаем конфиг
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	// Парсим конфиг в структуру
	if err := viper.Unmarshal(Config); err != nil {
		return err
	}

	return nil
}
