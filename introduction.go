package main

import (
	"fmt"
)

func PrintAlphabet() {
	for i := 'a'; i <= 'z'; i++ {
		fmt.Printf("%c", i)
	}
	fmt.Println()
}

func PrintReverseaAlphabet() {
	for i := 'z'; i >= 'a'; i-- {
		fmt.Printf("%c", i)
	}
	fmt.Println()
}
func PrintDigits() {
	for i := 0; i <= 9; i++ {
		fmt.Println(i)
	}
}
func IsNegative(n int) string {
	if n < 0 {
		return "T"

	}
	return "F"
}

func main() {
	PrintAlphabet()
	PrintReverseaAlphabet()
	PrintDigits()
	IsNegative(1)
	IsNegative(0)
	IsNegative(-1)
}

