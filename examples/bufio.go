package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Введите выражение в формате 2+5 (или 'exit' для выхода): ")
	//fmt.Scanln(&input)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
	//
	//fmt.Println("Enter text: ")
	//text2 := ""
	//fmt.Scanln(text2)
	//fmt.Println(text2)
	//
	//ln := ""
	//fmt.Sscanln("%v", ln)
	//fmt.Println(ln)
}
