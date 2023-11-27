package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	var x string
	//var y string
	//var operand string
	for {
		fmt.Print("Введите выражение в формате 2+5 (или 'exit' для выхода): ")
		fmt.Scanln(&input)

		if strings.ToLower(input) == "exit" {
			fmt.Println("Программа завершена")
			break
		}
		x = input[0]
		//y = strconv.Itoa(int(input[2]))
		//operand = strconv.Itoa(int(input[1]))
		fmt.Println(x)
	}
}

//	func input() {
//		var input string
//		for {
//			fmt.Println("Введите выражение (или 'exit' для выхода): ")
//			fmt.Scanln(&input)
//
//			if strings.ToLower(input) == "exit" {
//				fmt.Println("Программа завершена")
//				break
//			}
//
//			//result, err := calculate(input)
//			//if err != nil {
//			//	fmt.Println("Ошибка:", err)
//			//} else {
//			//	fmt.Println("Результат:", result)
//		}
//	}
//func Question() {
//	var input string
//	for {
//		fmt.Print("Введите выражение в формате 2+5 (или 'exit' для выхода): ")
//		fmt.Scanln(&input)
//
//		if strings.ToLower(input) == "exit" {
//			fmt.Println("Программа завершена")
//			break
//		}
//	}
//}

// question()
func Plus(x, y int) int {
	return x + y
}

func Mines(x, y int) int {
	return x - y
}

func Multiply(x, y int) int {
	return x * y
}

func Divide(x, y float64) float64 {
	return x / y
}
