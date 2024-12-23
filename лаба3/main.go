package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()

	// Маршрут для приветствия
	router.GET("/greet", GreetHandler)

	// Маршруты для калькулятора
	router.GET("/add", func(c *gin.Context) {
		CalculatorHandler(c, "add")
	})
	router.GET("/sub", func(c *gin.Context) {
		CalculatorHandler(c, "sub")
	})
	router.GET("/mul", func(c *gin.Context) {
		CalculatorHandler(c, "mul")
	})
	router.GET("/div", func(c *gin.Context) {
		CalculatorHandler(c, "div")
	})

	// Маршрут для подсчета символов в строке, подключен через charcount.go
	router.POST("/char_count", CharCountHandler)

	// Запускаем сервер на порту 8080
	router.Run(":8080")
}
