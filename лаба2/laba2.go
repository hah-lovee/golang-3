package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func FormatIP(b [4]byte) string {

	out := ""

	for i := 0; i < len(b); i++ {

		out += strconv.Itoa(int(b[i])) + "."

	}

	return out[:len(out)-1]

}

func listEven(a int, b int) ([]int, error) {
	var slice []int

	if a > b {
		return slice, errors.New("границы перепутаны")
	}

	for i := a; i <= b; i++ {
		if i%2 == 0 {
			slice = append(slice, i)
		}
	}
	return slice, nil
}

func MapTask(str string) map[string]int {

	var symbols map[string]int

	symbols = make(map[string]int)

	for i := 0; i < len(str); i++ {

		char := string(str[i])

		if count, ok := symbols[char]; ok {
			symbols[char] = count + 1
		} else {
			symbols[char] = 1
		}
	}

	return symbols
}

type Point struct {
	X float64
	Y float64
}

type Section struct {
	start Point
	last  Point
}

func (s Section) Length() float64 {
	xDiff := (s.last.X - s.start.X)
	yDiff := (s.last.Y - s.start.Y)

	return math.Sqrt(math.Pow(xDiff, 2) + math.Pow(yDiff, 2))
}

type Triangle struct {
	FirstAngle  Point
	SecondAngle Point
	ThirdAngle  Point
}

func (t Triangle) Area() float64 {

	return 0.5 * math.Abs(t.FirstAngle.X*(t.SecondAngle.Y-t.ThirdAngle.Y)+t.SecondAngle.X*(t.ThirdAngle.Y-t.FirstAngle.Y)+t.ThirdAngle.X*(t.FirstAngle.Y-t.SecondAngle.Y))

}

type Circle struct {
	Centre Point
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

type Shape interface {
	Area() float64
}

func printArea(s Shape) {
	result := s.Area()
	fmt.Printf("Площадь фигуры: %.2f\n", result)
}

func plus1(num int) int {

	return num + 1

}

func Map(ar []int, f func(int) int) []int {

	for i := 0; i < len(ar); i++ {

		ar[i] = f(ar[i])

	}

	return ar

}

func main() {
	// 1
	test1 := [4]byte{127, 0, 0, 1}
	fmt.Println(FormatIP(test1))

	// 2
	fmt.Println(listEven(12, 11))
	fmt.Println(listEven(12, 12))
	fmt.Println(listEven(12, 19))

	//3
	res := MapTask("sassdfdffd")

	for k, v := range res {
		fmt.Println("Символ:"+k, "количесвто:", +v)
	}

	triangle := Triangle{
		FirstAngle:  Point{X: 0, Y: 0},
		SecondAngle: Point{X: 4, Y: 0},
		ThirdAngle:  Point{X: 0, Y: 3},
	}
	printArea(triangle) // 6.00

	circle := Circle{
		Centre: Point{X: 0, Y: 0},
		Radius: 5,
	}
	printArea(circle) // 78.54

	ar := []int{1, 2, 3, 4, 5}

	fmt.Println(Map(ar, plus1))

}
