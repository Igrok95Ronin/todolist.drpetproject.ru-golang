package routes

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Обработчик для получения всех заметок пользователя
func (h *handler) GetNotes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Получаем имя пользователя из контекста
	username := r.Context().Value("username").(string)
	// Получаем объект базы данных из контекста запроса
	db := r.Context().Value("db").(*gorm.DB)
	var user User
	// Находим пользователя в базе данных по имени пользователя
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		h.logger.Errorf("User not found: %s", err)
		return
	}

	var notes []Note
	// Находим все заметки пользователя
	if err := db.Where("user_id = ?", user.ID).Find(&notes).Error; err != nil {
		http.Error(w, "Failed to get notes", http.StatusInternalServerError)
		h.logger.Errorf("Failed to get notes: %s", err)
		return
	}

	// Отправляем заметки в ответе
	json.NewEncoder(w).Encode(notes)
}
