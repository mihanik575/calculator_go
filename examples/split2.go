//package main
//
//import (
//	"fmt"
//	"strconv"
//)

//import (
//	"fmt"
//	"strconv"
//)
//
//func convertStringToOperator(input string) (int, int, string) {
//	num1, _ := strconv.Atoi(string(input[0]))
//	operator := string(input[1])
//	num2, _ := strconv.Atoi(string(input[2]))
//
//	return num1, num2, operator
//}
//
//func main() {
//	num1, num2, operator := convertStringToOperator("2+5")
//	fmt.Printf("num1: %d, num2: %d, operator: %s\n", num1, num2, operator)
//}

//import (
//	"fmt"
//	"strconv"
//	"strings"
//)
//
//func convertStringToOperators(input string) (int, int, string) {
//	// Splitting the input string to extract numbers and the operator
//	parts := strings.Split(input, "+")
//	num1, err1 := strconv.Atoi(parts[0])
//	num2, err2 := strconv.Atoi(parts[1])
//	if err1 != nil || err2 != nil {
//		fmt.Println("Error converting string to int")
//		return 0, 0, ""
//	}
//	return num1, num2, "+"
//}
//
//func main() {
//	input := "12+5"
//	num1, num2, operator := convertStringToOperators(input)
//	fmt.Println("Number 1:", num1)
//	fmt.Println("Number 2:", num2)
//	fmt.Println("Operator:", operator)
//}

package main

import (
	"fmt"
	"strconv"
)

func parseString(s string) (int, int, string) {
	var num1, num2 int
	var operator string

	for i, c := range s {
		if c == '+' || c == '-' || c == '*' || c == '/' {
			// operator found
			num1, _ = strconv.Atoi(s[:i])
			num2, _ = strconv.Atoi(s[i+1:])
			operator = string(c)
			break
		}
	}

	return num1, num2, operator
}

func main() {
	num1, num2, operator := parseString("12*5")
	fmt.Println("Num1:", num1)
	fmt.Println("Num2:", num2)
	fmt.Println("Operator:", operator)
}
