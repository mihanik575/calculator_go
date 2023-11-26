package main

import "fmt"

//	for {
//		fmt.Print("Введите выражение (в формате 2+5) (или 'exit' для выхода): ")
//		fmt.Scanln(&input)
//
//		if strings.ToLower(input) == "exit" {
//			fmt.Println("Программа завершена")
//		break
//	}
func plus(x, y int) int {
	return x + y
}
func mines(x, y int) int {
	return x - y
}

func minus(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y float64) float64 {
	return x / y
}

func main() {
	fmt.Println(plus(42, 13))
	fmt.Println(mines(42, 13))
	fmt.Println(divide(42, 13))
	fmt.Println(multiply(42, 13))
}
