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

	var num1str, num2str string
	var num1, num2 int
	var operator string

	for {
		fmt.Print("Введите выражение в формате 2+5 (или 'exit' для выхода): ")

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')

		if strings.ToLower(input) == "exit" {
			fmt.Println("Программа завершена")
			break
		}

		for i, c := range input {
			if c == '+' || c == '-' || c == '*' || c == '/' {

				num1str = strings.TrimSpace(input[:i])
				num2str = strings.TrimSpace(input[i+1:])
				num1, _ = strconv.Atoi(num1str)
				num2, _ = strconv.Atoi(num2str)

				operator = string(c)
				break
			}
		}

		result, err := Сalculation(num1, num2, operator)

		if err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			fmt.Println("Результат:", result)
		}
	}
}

func Сalculation(num1, num2 int, operator string) (int, error) {

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
