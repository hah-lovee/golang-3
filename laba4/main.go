package main

import (
	"encoding/json"
	"fmt"
	"laba4/database"
	"laba4/models"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Инициализация базы данных
	database.InitDB() // Не используем переменную db напрямую
	defer database.CloseDB()

	// Настройка маршрутов
	router := http.NewServeMux()
	router.HandleFunc("/customers/register", registerCustomer)
	router.HandleFunc("/customers/login", loginCustomer)

	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Обработчик регистрации пользователя
func registerCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Декодирование JSON тела запроса
	var customer models.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, "Невалидный JSON", http.StatusBadRequest)
		return
	}

	// Проверка на существование email в базе данных
	exists, err := database.GetCustomerByEmail(customer.Email)
	if err == nil && exists.Email != "" {
		http.Error(w, "Email уже зарегистрирован", http.StatusConflict)
		return
	}

	// Хеширование пароля
	hash, err := bcrypt.GenerateFromPassword([]byte(customer.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Ошибка хеширования пароля", http.StatusInternalServerError)
		return
	}
	customer.PasswordHash = string(hash)

	// Сохранение пользователя в базе данных
	if err := database.SaveCustomer(customer); err != nil {
		http.Error(w, "Ошибка сохранения в базу данных", http.StatusInternalServerError)
		return
	}

	// Успешная регистрация
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Пользователь зарегистрирован: %s", customer.Name)
}

// Обработчик логина пользователя
func loginCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Декодирование JSON тела запроса
	var customer models.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, "Невалидный JSON", http.StatusBadRequest)
		return
	}

	// Получаем покупателя по email
	storedCustomer, err := database.GetCustomerByEmail(customer.Email)
	if err != nil || storedCustomer.Email == "" {
		http.Error(w, "Неверный email или пароль", http.StatusUnauthorized)
		return
	}

	// Проверка пароля
	if !database.CheckPasswordHash(customer.PasswordHash, storedCustomer.PasswordHash) {
		http.Error(w, "Неверный email или пароль", http.StatusUnauthorized)
		return
	}

	// Генерация JWT токена
	token, err := database.GenerateJWT(fmt.Sprintf("%d", storedCustomer.ID))
	if err != nil {
		http.Error(w, "Ошибка генерации токена", http.StatusInternalServerError)
		return
	}

	// Возвращаем JWT токен клиенту
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
