package main

import (
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/internal/config"
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/internal/routes"
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

func main() {
	// Получаем экземпляр логгера
	logger := logging.GetLogger()

	// Создаем новый HTTP роутер
	router := httprouter.New()

	// Читаем конфигурацию приложения
	cfg := config.GetConfig(logger)

	// Регистрируем обработчик в роутере
	handler := routes.NewHandler(logger)
	handler.Register(router, logger)

	// Запускаем приложение
	start(router, cfg, logger)
}

// Функция start запускает приложение
func start(router *httprouter.Router, cfg *config.Config, logger *logging.Logger) {
	server := &http.Server{
		Handler:      router,
		Addr:         cfg.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	logger.Infof("The server is running on port %s", cfg.Port)
	logger.Fatal(server.ListenAndServe())
}
