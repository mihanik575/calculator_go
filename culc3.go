package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var input string

	for {
		fmt.Print("Введите выражение в формате 2+5 (или 'exit' для выхода): ")

		fmt.Scanln(&input)

		result, err := calculation(input)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Result:", result)
		}

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
}
