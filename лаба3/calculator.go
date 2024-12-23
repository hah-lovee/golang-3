package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CalculatorHandler обрабатывает арифметические операции
func CalculatorHandler(c *gin.Context, operation string) {
	aStr := c.Query("a")
	bStr := c.Query("b")

	a, err1 := strconv.Atoi(aStr)
	b, err2 := strconv.Atoi(bStr)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var result float64

	switch operation {
	case "add":
		result = float64(a + b)
	case "sub":
		result = float64(a - b)
	case "mul":
		result = float64(a * b)
	case "div":
		if b == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Division by zero"})
			return
		}
		result = float64(a) / float64(b)
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}
