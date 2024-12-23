package main

import (
	"errors"
	"fmt"
)

func hello(name string) string {
	return "Привет, " + name + "!"
}

func printEven(start, end int64) error {
	if start > end {
		return errors.New("левая граница больше правой")
	}

	for i := start; i <= end; i++ {
		if i%2 == 0 {
			fmt.Print(i)
			fmt.Print(" ")
		}
	}

	fmt.Println()

	return nil
}

func apply(a, b float64, operator string) (float64, error) {

	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль")
		}
		return a / b, nil
	default:
		return 0, errors.New("действие не поддерживается")
	}

}

func main() {

	fmt.Println("Задание 2")

	name := "Мир"
	fmt.Println(hello(name))

	name2 := "лаба1"
	fmt.Println(hello(name2))

	fmt.Println()
	fmt.Println("Задание 3")

	if err := printEven(3, 21); err != nil {
		fmt.Println("Ошибка:", err)
	}

	if err := printEven(22, 22); err != nil {
		fmt.Println("Ошибка:", err)
	}

	fmt.Println()
	fmt.Println("Задание 4")

	result, err := apply(3, 5, "+")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}

	result, err = apply(7, 10, "*")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}

	result, err = apply(3, 5, "#")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}
}
