package config

import (
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

// Определяем структуру для хранения конфигурации
type Config struct {
	Port         string `yaml:"port"`
	MySigningKey string `yaml:"mySigningKey"`
}

// Глобальная переменная для хранения экземпляра конфигурации
var instance *Config

// Синхронизатор для однократного создания экземпляра конфигурации
var once sync.Once

// Функция GetConfig возвращает экземпляр конфигурации
func GetConfig() *Config {
	logger := logging.GetLogger()
	once.Do(func() {
		instance = &Config{}

		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Error(err)
		}
	})

	return instance
}
