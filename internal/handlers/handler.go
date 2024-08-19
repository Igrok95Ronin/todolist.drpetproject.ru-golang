package handlers

import (
	"github.com/julienschmidt/httprouter"
)

type Handler interface {
	Routes(router *httprouter.Router)
}
