package routes

import (
	"context"
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/internal/config"
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
	cfg    *config.Config
}

// Заполняем структуру
func NewHandler(logger *logging.Logger, db *gorm.DB, cfg *config.Config) handlers.Handler {
	return &handler{
		logger,
		db,
		cfg,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	// Middleware для добавления базы данных в контекст
	dbMiddleware := func(next httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			ctx := context.WithValue(r.Context(), "db", h.db)
			next(w, r.WithContext(ctx), ps)
		}
	}

	router.GET("/", h.Home)
	router.POST("/register", dbMiddleware(h.RegisterUser))                  // Маршрут для регистрации пользователя
	router.POST("/login", dbMiddleware(h.Login))                            // Маршрут для авторизации пользователя
	router.POST("/notes", authMiddleware(dbMiddleware(h.CreateNote)))       // Защищенный маршрут для создания заметки
	router.GET("/notes", authMiddleware(dbMiddleware(h.GetNotes)))          // Защищенный маршрут для получения всех заметок пользователя
	router.PUT("/notes/:id", authMiddleware(dbMiddleware(h.UpdateNote)))    // Защищенный маршрут для обновления заметки
	router.DELETE("/notes/:id", authMiddleware(dbMiddleware(h.DeleteNote))) // Защищенный маршрут для удаления заметки
}
