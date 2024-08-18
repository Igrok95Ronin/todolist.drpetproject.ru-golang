package routes

import (
	"context"
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/internal/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Middleware для проверки авторизации пользователя
func authMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		cfg := config.GetConfig()

		// Получаем JWT токен из cookies
		cookie, err := r.Cookie("jwt")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized) // если ошибка, возвращаем статус 401
			return
		}

		// Проверяем и декодируем токен
		token, err := jwt.ParseWithClaims(cookie.Value, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.MySigningKey), nil // возвращаем секретный ключ для проверки подписи токена
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized) // если ошибка, возвращаем статус 401
			return
		}

		claims, ok := token.Claims.(*MyCustomClaims)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized) // если ошибка, возвращаем статус 401
			return
		}

		// Сохраняем информацию о пользователе в контексте запроса
		ctx := context.WithValue(r.Context(), "username", claims.Username)
		next(w, r.WithContext(ctx), ps) // передаем управление следующему обработчику
	}
}
