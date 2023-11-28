// Program using Replace() to replace strings

//package main
//
//import (
//	"strings"
//	"testing"
//	"unicode"
//)

//func main() {
//
//	text := "car"
//	fmt.Println("Old String:", text)
//
//	// replace r with t
//	replacedText := strings.Replace(text, "r", "t", 1)
//
//	fmt.Println("New String:", replacedText)
//}

//func main() {
//
//	text := "2 + 5"
//	fmt.Println("Old String:", text)
//
//	// replace r with t
//	replacedText := strings.Replace(text, " ", "", 1)
//
//	fmt.Println("New String:", replacedText)
//}

package main_test

import (
	"strings"
	"testing"
	"unicode"
)

func SpaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func SpaceFieldsJoin(str string) string {
	return strings.Join(strings.Fields(str), "")
}

func SpaceStringsBuilder(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

func BenchmarkSpaceMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SpaceMap(data)
	}
}

func BenchmarkSpaceFieldsJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SpaceFieldsJoin(data)
	}
}

func BenchmarkSpaceStringsBuilder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SpaceStringsBuilder(data)
	}
}
