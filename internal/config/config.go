package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

// Определяем структуру для хранения конфигурации
type Config struct {
	Port string `yaml:"port"`
}

// Глобальная переменная для хранения экземпляра конфигурации
var instance *Config

// Синхронизатор для однократного создания экземпляра конфигурации
var once sync.Once

// Функция GetConfig возвращает экземпляр конфигурации
func GetConfig() *Config {
	//logger := logging.GetLogger() // Получаем экземпляр логгера
	once.Do(func() {
		instance = &Config{}

		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Println(help)
			log.Fatal(err)
		}
	})

	return instance
}
