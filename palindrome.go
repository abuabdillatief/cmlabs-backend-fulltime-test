package main

import "strings"

func Palindrome(input string) bool {
	input = strings.ToLower(strings.ReplaceAll(input, " ", ""))

	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}

	return true
}
