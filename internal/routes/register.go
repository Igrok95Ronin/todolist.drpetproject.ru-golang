package routes

import (
	"context"
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/internal/handlers"
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/pkg/logging"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var _ handlers.Handler = &handler{} // Проверяем что интерфейс реализуется

type handler struct {
	logger *logging.Logger
	db     *gorm.DB
}

// Заполняем структуру
func NewHandler(logger *logging.Logger, db *gorm.DB) handlers.Handler {
	return &handler{
		logger,
		db,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	// Middleware для добавления базы данных в контекст
	dbMiddleware := func(next httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			ctx := context.WithValue(r.Context(), "db", h.db)
			next(w, r.WithContext(ctx), ps)
			//h.logger.Info("DB added to context")
		}
	}

	router.GET("/", h.Home)
	router.POST("/register", dbMiddleware(h.RegisterUser)) // Маршрут для регистрации пользователя
	router.POST("/login", dbMiddleware(h.login))           // Маршрут для авторизации пользователя
	router.POST("/notes", h.CreateNote)

}
