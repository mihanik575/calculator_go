package main

import (
	"fmt"
	"strings"
)

func main() {
	x := "2+5"
	y := "2 + 5"
	fmt.Printf("%c\n", x[0])
	fmt.Printf("%c\n", x[1])
	fmt.Printf("%c\n", x[2])
	fmt.Printf("%c\n", y[1])
	fmt.Printf("%q\n", strings.Split(x, " "))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(y, " "))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ""))
	fmt.Printf("%q\n", strings.SplitAfter("2+5", " "))
	fmt.Printf("%q\n", strings.SplitAfter("2 + 5", ","))
}
