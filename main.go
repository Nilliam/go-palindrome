package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter the palindrome")

	var palindrome string
	fmt.Scanf("%s", &palindrome)

	if isValidPalindrome(palindrome) {
		fmt.Println("It is a palindrome")
	} else {
		fmt.Println("It is not a palindrome")
	}
}

func isValidPalindrome(palindrome string) bool {
	length := len(palindrome)

	for i := 0; i < (length / 2); i++ {
		if palindrome[i] != palindrome[length-1-i] {
			return false
		}
	}

	return true
}

