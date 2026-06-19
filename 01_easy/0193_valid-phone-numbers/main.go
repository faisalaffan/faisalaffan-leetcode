package main

// LeetCode #193: Valid Phone Numbers
// https://leetcode.com/problems/valid-phone-numbers/
// Difficulty: Easy

import "fmt"

func ValidPhoneNumbers() string {
	return `grep -E '^(\([0-9]{3}\) [0-9]{3}-[0-9]{4}|[0-9]{3}-[0-9]{3}-[0-9]{4})$' file.txt`
}

func main() {
	fmt.Println(ValidPhoneNumbers())
}
