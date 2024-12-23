package database

import (
	"database/sql"
	"laba4/models"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

// InitDB Инициализация базы данных
func InitDB() *sql.DB {
	var err error
	connStr := "host=localhost port=5432 user=postgres password=1122 dbname=customers sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// CloseDB Закрытие соединения с базой данных
func CloseDB() {
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}

// SaveCustomer Сохранение нового пользователя
func SaveCustomer(customer models.Customer) error {
	query := `INSERT INTO customers(name, email, phone, password_hash) VALUES($1, $2, $3, $4)`
	_, err := db.Exec(query, customer.Name, customer.Email, customer.Phone, customer.PasswordHash)
	return err
}

// GetCustomerByEmail Получить пользователя по email
func GetCustomerByEmail(email string) (models.Customer, error) {
	var customer models.Customer
	query := `SELECT id, name, email, phone, password_hash FROM customers WHERE email=$1`
	row := db.QueryRow(query, email)
	err := row.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Phone, &customer.PasswordHash)
	if err != nil {
		return customer, err
	}
	return customer, nil
}

// CheckPasswordHash Проверка пароля
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateJWT Генерация JWT токена
func GenerateJWT(userID string) (string, error) {
	// Срок действия токена — 1 день
	expirationTime := time.Now().Add(24 * time.Hour)

	// Создаем полезную нагрузку
	claims := jwt.MapClaims{
		"sub": userID,                // ID пользователя
		"exp": expirationTime.Unix(), // Время истечения токена
	}

	// Генерируем токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен секретным ключом
	secretKey := []byte("your-256-bit-secret")
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
