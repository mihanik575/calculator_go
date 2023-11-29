package main

import (
	"fmt"
	"strconv"
	"strings"
)

//type num1 int
//type num2 int
//type operator string

func main() {
	var input string

	for {
		fmt.Print("Введите выражение в формате 2+5 (или 'exit' для выхода): ")

		fmt.Scanln(&input)

		//fmt.Println(input)

		result, err := calculation(input)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Result:", result)
		}

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

func calculation(input string) (int, error) {
	var num1 int
	var num2 int
	var operator string

	if len(input) == 3 {
		num1, _ = strconv.Atoi(string(input[0]))
		operator = string(input[1])
		num2, _ = strconv.Atoi(string(input[2]))
		//input = strings.Split(input)
		//num1, _ = strconv.Atoi()
		//num2 = int(input[2])
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
		return 0, fmt.Errorf("неверный формат укажите числа до 9")
	}

	//if error != nil {
	//	return 0, error
	//}

}

//func evaluateExpression(input string) (int, error) {
//	var num1, num2 int
//	var operator string
//	_, err := fmt.Sscanf(input, "%d %s %d", &num1, &operator, &num2)
//	if err != nil {
//		return 0, err
//	}
//
//	switch operator {
//	case "+":
//		return num1 + num2, nil
//	case "-":
//		return num1 - num2, nil
//	case "*":
//		return num1 * num2, nil
//	case "/":
//		if num2 == 0 {
//			return 0, fmt.Errorf("деление на ноль")
//		}
//		return num1 / num2, nil
//	default:
//		return 0, fmt.Errorf("неверная операция")
//	}
//}

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
