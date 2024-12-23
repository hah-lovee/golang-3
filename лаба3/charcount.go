package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CharCountHandler - обработчик для подсчета количества символов
func CharCountHandler(c *gin.Context) {

	// Структура для хранения входящего JSON
	var requestBody struct {
		Text string `json:"text"`
	}

	// Парсим JSON из тела запроса
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Подсчитываем количество вхождений каждого символа
	charCount := make(map[string]int)
	for _, char := range requestBody.Text {
		charCount[string(char)]++
	}

	// Возвращаем результат в формате JSON
	c.JSON(http.StatusOK, charCount)
}
