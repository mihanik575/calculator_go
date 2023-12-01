package main

import "fmt"

func convertRomanToArabic(romanNum string) int {
	romanMap := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	return romanMap[romanNum]
}

func main() {
	fmt.Println(convertRomanToArabic("V"))  // Output: 5
	fmt.Println(convertRomanToArabic("IX")) // Output: 9
	fmt.Println(convertRomanToArabic("X"))  // Output: 10
}
