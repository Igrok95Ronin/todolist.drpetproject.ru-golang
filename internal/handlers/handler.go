package handlers

import (
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

type Handler interface {
	Register(router *httprouter.Router, logger *logging.Logger)
}
