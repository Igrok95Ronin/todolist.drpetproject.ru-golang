package main

import (
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/internal/config"
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/internal/routes"
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"net/http"
	"time"
)

func main() {
	// Инициализация базы данных
	db := routes.InitDB()
	defer db.Close() // закрываем соединение с базой данных при завершении работы

	// Получаем экземпляр логгера
	logger := logging.GetLogger()

	// Создаем новый HTTP роутер
	router := httprouter.New()

	// Читаем конфигурацию приложения
	cfg := config.GetConfig()

	// Создайте обработчик CORS с параметрами по умолчанию.
	corsH := cors.Default().Handler(router)

	// Регистрируем обработчик в роутере
	handler := routes.NewHandler(logger, db, cfg)
	handler.Register(router)

	// Запускаем приложение
	start(corsH, cfg, logger)
}

// Функция start запускает приложение
func start(router http.Handler, cfg *config.Config, logger *logging.Logger) {
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
