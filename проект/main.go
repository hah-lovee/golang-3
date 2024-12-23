package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	// Путь к изображению
	filePath := `C:\Users\legyx\Desktop\qq.jpg`

	// Открываем файл
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Ошибка открытия файла: %v\n", err)
		return
	}
	defer file.Close()

	// Создаем буфер для данных формы
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Добавляем файл в запрос
	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		fmt.Printf("Ошибка создания файла в форме: %v\n", err)
		return
	}
	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Printf("Ошибка копирования файла в запрос: %v\n", err)
		return
	}

	// Закрываем writer, чтобы завершить формирование запроса
	writer.Close()

	// Создаем HTTP-запрос
	req, err := http.NewRequest("POST", "http://127.0.0.1:8000/analyze", body)
	if err != nil {
		fmt.Printf("Ошибка создания HTTP-запроса: %v\n", err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Отправляем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Ошибка отправки запроса: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Читаем ответ
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Ошибка чтения ответа: %v\n", err)
		return
	}

	// Выводим ответ
	fmt.Printf("Ответ сервера: %s\n", respBody)
}
