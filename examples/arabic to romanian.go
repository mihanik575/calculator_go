package main

import "fmt"

func arabicToRoman(num int) string {
	arabic := {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	roman := {"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	result := ""


	if num < 1 || num > 10 {
		return "Invalid input, number should be between 1 and 10"
	}



	for i := 4; i >= 0; i-- {
		for num >= arabic {
			result += roman
			num -= arabic
		}
	}

	return result
}

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d - %s\n", i, arabicToRoman(i))
	}
}
