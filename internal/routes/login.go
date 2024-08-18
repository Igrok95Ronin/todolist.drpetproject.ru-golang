package routes

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// Определяем секретный ключ для подписи JWT токенов
var mySigningKey = []byte("supersecretkey")

// Структура для входных данных при авторизации пользователя
type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Структура для утверждений (claims) JWT токена
type MyCustomClaims struct {
	Username           string `json:"username"` // имя пользователя
	jwt.StandardClaims        // стандартные утверждения JWT
}

// Обработчик для авторизации пользователя
func (h *handler) login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var input LoginInput
	// Декодируем JSON данные из тела запроса в структуру LoginInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Errorf("Failed authorizations: %s", err)
		return
	}

	// Получаем объект базы данных из контекста запроса
	db := r.Context().Value("db").(*gorm.DB)
	var user User
	// Находим пользователя в базе данных по имени пользователя
	if err := db.Where("username = ?", input.Username).First(&user).Error; err != nil {
		http.Error(w, "Invalid username", http.StatusUnauthorized)
		h.logger.Errorf("Invalid username: %s", err)
		return
	}

	// Проверяем пароль
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		h.logger.Errorf("Invalid password: %s", err)
		return
	}

	// Создаем JWT токен с утверждениями
	claims := MyCustomClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(), // устанавливаем время истечения через 72 часа
			Issuer:    "Todolist",                            // устанавливаем издателя токена (название вашего приложения)
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // создаем новый токен с использованием HMAC SHA256
	tokenString, err := token.SignedString(mySigningKey)       // подписываем токен
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError) // если ошибка, возвращаем статус 500
		h.logger.Errorf("Failed to generate token: %s", err)
		return
	}

	// Сохраняем JWT токен в cookies
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",                          // имя cookies
		Value:    tokenString,                    // значение cookies (JWT токен)
		Expires:  time.Now().Add(72 * time.Hour), // срок действия cookies
		Path:     "/",                            // путь, на который распространяется cookies
		HttpOnly: true,                           // флаг HttpOnly для предотвращения доступа к cookies через JavaScript
	})

	w.Write([]byte("Login successful"))

}
