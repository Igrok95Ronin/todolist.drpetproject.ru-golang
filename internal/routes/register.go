package routes

import (
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/internal/handlers"
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

var _ handlers.Handler = &handler{} // Проверяем что интерфейс реализуется

type handler struct {
	logger *logging.Logger
}

// Заполняем структуру
func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger,
	}
}

func (h *handler) Register(router *httprouter.Router, logger *logging.Logger) {
	router.GET("/", h.Home)
	router.POST("/notes", h.CreateNote)
}
