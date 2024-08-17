package main

import (
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/internal/config"
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/internal/routes"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

func main() {
	// Читаем конфигурацию приложения
	cfg := config.GetConfig()

	// Создаем новый HTTP роутер
	router := httprouter.New()

	// Регистрируем обработчик в роутере
	routes.NewHandler().Register(router)

	// Запускаем приложение
	start(router, cfg)
}

// Функция start запускает приложение
func start(router *httprouter.Router, cfg *config.Config) {
	server := &http.Server{
		Handler:      router,
		Addr:         cfg.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Println("Start..")
	log.Fatal(server.ListenAndServe())
}
