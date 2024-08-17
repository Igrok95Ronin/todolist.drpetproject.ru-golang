package routes

import (
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/internal/handlers"
	"github.com/julienschmidt/httprouter"
)

var _ handlers.Handler = &handler{}

type handler struct{}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET("/", h.Home)
}
