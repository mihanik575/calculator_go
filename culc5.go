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

	for {
		fmt.Print("Введите выражение в формате 2+5 (или 'exit' для выхода): ")

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка:", err)
		}

		if strings.ToLower(input) == "exit" {
			fmt.Println("Программа завершена")
			break
		}

		result, err := calculation(input)
		if err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			fmt.Println("Результат:", result)
		}

	}
}

func calculation(input string) (int, error) {
	var num1str, num2str string
	var num1, num2 int
	var operator string

	for i, c := range input {
		if c == '+' || c == '-' || c == '*' || c == '/' {
			// operator found
			num1str = strings.TrimSpace(input[:i])
			num2str = strings.TrimSpace(input[i+1:])
			num1, _ = strconv.Atoi(num1str)
			num2, _ = strconv.Atoi(num2str)

			operator = string(c)
			break
		}
	}

	if num1 < 0 || num1 > 11 || num2 < 0 || num2 > 11 {

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
