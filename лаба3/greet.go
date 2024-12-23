package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GreetHandler обрабатывает запросы для приветствия
func GreetHandler(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	if name == "" || age == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name and age are required"})
		return
	}

	response := fmt.Sprintf("Меня зовут %s, мне %s лет", name, age)
	c.String(http.StatusOK, response)
}
