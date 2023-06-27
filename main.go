package main

import "fmt"
func main() {
	FizzBuzz()

	words := []string{
		"SAIPPUAKIVIKAUPPIAS",
		"Aibohphobia",
		"Anna",
		"Civic",
		"My gym",
		"No lemon, no melon",
		"bad",
		"wrong",
	}

	for _, v := range words {
		fmt.Println(v,":", Palindrome((v)))
	}
}
