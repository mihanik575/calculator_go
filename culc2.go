package main

import (
	"fmt"
	"strings"
)

//type num1 int
//type num2 int
//type operator string

func main() {
	var input string

	for {
		fmt.Println("Введите выражение в формате 2+5 (или 'exit' для выхода): ")
		fmt.Scanln(input)

		fmt.Println(input)
		evaluateExpression(input)

		//input = strings.ReplaceAll(input, " ", "")
		//result := strings.Join(strings.Fields(input), "")
		//x := (int(input[1]))
		//y = strconv.Itoa(int(input[2]))
		//operand = strconv.Itoa(int(input[1]))
		//fmt.Println(result)

		if strings.ToLower(input) == "exit" {
			fmt.Println("Программа завершена")
			break

		}

	}

}

func evaluateExpression(input string) (int, error) {
	var num1 int
	var num2 int
	var operator string

	if len(input) == 3 {
		num1 = int(input[0])
		num2 = int(input[2])
		operator = string(input[1])
		switch operator {
		case "+":
			return num1 + num2, nil
		case "-":
			return num1 - num2, nil
		case "*":
			return num1 * num2, nil
		case "/":
			if num2 == 0 {
				return 0, fmt.Errorf("деление на ноль")
			}
			return num1 / num2, nil
		default:
			return 0, fmt.Errorf("неверная операция")
		}
	} else {
		fmt.Println("неправильное значение, введите цифры до 10 в формате 2+3")
	}
	//if err != nil {
	//	return 0, err
	//}

}

//func Plus(x, y int) int {
//	return x + y
//}
//
//func Mines(x, y int) int {
//	return x - y
//}
//
//func Multiply(x, y int) int {
//	return x * y
//}
//
//func Divide(x, y float64) float64 {
//	return x / y
//}
