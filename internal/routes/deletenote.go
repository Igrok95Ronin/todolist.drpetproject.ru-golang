package routes

import (
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Обработчик для удаления заметки
func (h *handler) DeleteNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Получаем имя пользователя из контекста
	username := r.Context().Value("username")
	// Получаем объект базы данных из контекста запроса
	db := r.Context().Value("db").(*gorm.DB)
	var user User
	// Находим пользователя в базе данных по имени пользователя
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		h.logger.Errorf("User not found: %s", err)
		return
	}

	// Проверяем, что заметка принадлежит пользователю
	var note Note
	if err := db.Where("id = ? AND user_id = ?", ps.ByName("id"), user.ID).First(&note).Error; err != nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		h.logger.Errorf("Note not found: %s", err)
		return
	}

	// Удаляем заметку
	if err := db.Delete(&note).Error; err != nil {
		http.Error(w, "Failed to delete note", http.StatusInternalServerError)
		h.logger.Errorf("Failed to delete note", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Note deleted successfully"))
}
