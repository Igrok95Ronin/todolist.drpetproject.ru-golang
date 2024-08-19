package routes

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Определяем модель заметки для базы данных
type Note struct {
	gorm.Model        // включает встроенные поля ID, CreatedAt, UpdatedAt, DeletedAt
	Title      string `json:"title"` // заголовок заметки
	UserID     uint   // ID пользователя, который создал заметку
}

// Обработчик для создания новой заметки
func (h *handler) CreateNote(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	// Присваиваем заметке ID пользователя
	note.UserID = user.ID

	// Сохраняем заметку в базе данных
	if err := db.Create(&note).Error; err != nil {
		http.Error(w, "Failed to create note", http.StatusInternalServerError)
		h.logger.Errorf("Failed to create note: %s", err)
		return
	}

	w.Write([]byte("Note create successfully"))
}
