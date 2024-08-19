package routes

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Обработчик для обновления заметки
func (h *handler) UpdateNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var note Note
	// Декодируем JSON данные из тела запроса в структуру Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Error(err)
		return
	}

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

	// Проверяем, что заметка принадлежит пользователю
	var existingNote Note
	if err := db.Where("id = ? AND user_id = ?", ps.ByName("id"), user.ID).First(&existingNote).Error; err != nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		h.logger.Errorf("Note not found: %s", err)
		return
	}

	// Обновляем заметку
	existingNote.Title = note.Title
	if err := db.Save(&existingNote).Error; err != nil {
		http.Error(w, "Failed to update note", http.StatusInternalServerError)
		h.logger.Errorf("Failed to update note: %s", err)
		return
	}

	w.Write([]byte("Note updated successfully"))
}
