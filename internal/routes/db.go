package routes

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // используем SQLite в качестве базы данных
)

// Определяем модель пользователя для базы данных
type User struct {
	gorm.Model        // включает встроенные поля ID, CreatedAt, UpdatedAt, DeletedAt
	Username   string `gorm:"unique"` // уникальное имя пользователя
	Password   string // хешированный пароль пользователя
}

// Функция для инициализации базы данных
func InitDB() *gorm.DB {
	// Открываем соединение с базой данных SQLite
	db, err := gorm.Open("sqlite3", "todolist.db")
	if err != nil {
		panic("Failed to connect database") // если не удается подключиться к базе данных, выходим с ошибкой
	}

	db.AutoMigrate(&User{})
	return db
}
