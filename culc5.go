package main

// работает
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//var input string

	for {
		//fmt.Println("Введите выражение в формате 2+5 (или 'exit' для выхода): ")
		//fmt.Scanln(&input)
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Введите выражение в формате 2+5 (или 'exit' для выхода): ")
		input, err := reader.ReadString('\n')
		fmt.Println(input)
		if err != nil {
			fmt.Println("Ошибка:", err)
		}
		if strings.ToLower(input) == "exit" {
			fmt.Println("Программа завершена")
			break
		}

		result, err := calculation2(input)
		if err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			fmt.Println("Результат:", result)
		}

	}
}

func calculation2(input string) (int, error) {
	var num1, num2 int
	var operator string
	fmt.Println(input)
	for i, c := range input {
		if c == '+' || c == '-' || c == '*' || c == '/' {
			num1, _ = strconv.Atoi(input[:i])
			num2, _ = strconv.Atoi(input[i+1:])
			operator = string(c)
			break
		}
	}
	fmt.Println(num1, operator, num2)
	if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
		return 0, fmt.Errorf("Неправильный ввод введите число от 1 до 10")
	}
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
}
