package main

import (
	"fmt"
	"strings"
)

func palindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return input
}

func main() {
    var n string
	fmt.Scan(&n)
	fmt.Println(palindrome(n))
}