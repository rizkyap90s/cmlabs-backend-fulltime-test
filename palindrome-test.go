package main

import (
	"fmt"
	"strings"
)

func main() {

	words := [7]string{
		"SAIPPUAKIVIKAUPPIAS",
		"Aibohphobia",
		"Anna",
		"Civic",
		"My gym",
		"No lemon, no melon",
		"not palindrome sample",
	}

	for i := 0; i < len(words); i++ {
		if isPalindrome(words[i]) {
			fmt.Println(words[i], "is palindrome")
		} else {
			fmt.Println(words[i], "is not palindrome")
		}
	}

}

func isPalindrome(word string) bool {
	word = strings.ToLower(strings.ReplaceAll(word, " ", ""))
	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-i-1] {
			return false
		}
	}
	return true
}
