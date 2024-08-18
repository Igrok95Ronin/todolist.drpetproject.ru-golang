package routes

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// Структура для входных данных при регистрации пользователя
type RegisterInput struct {
	Username string `json:"username"` // имя пользователя
	Password string `json:"password"` // пароль пользователя
}

// Обработчик для регистрации нового пользователя
func (h *handler) RegisterUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var input RegisterInput
	// Декодируем JSON данные из тела запроса в структуру RegisterInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Errorf("Failed Register: %s", err)
		return
	}

	// Хешируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		h.logger.Errorf("Failed to hash password : %s", err)
		return
	}

	// Создаем пользователя с хешированным паролем
	user := User{
		Username: input.Username,
		Password: string(hashedPassword),
	}

	// Получаем объект базы данных из контекста запроса
	db := r.Context().Value("db").(*gorm.DB)
	// Сохраняем пользователя в базе данных
	if err = db.Create(&user).Error; err != nil {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		h.logger.Errorf("Username alredy exists : %s", err)
		return
	}

	w.Write([]byte("Registration successful")) // отправляем успешный ответ
}
