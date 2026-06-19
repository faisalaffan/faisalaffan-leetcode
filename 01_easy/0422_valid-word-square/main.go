package main

// LeetCode #422: Valid Word Square
// https://leetcode.com/problems/valid-word-square/
// Difficulty: Easy [Paid]

import "fmt"

// Time: O(n*m), Space: O(1)
func ValidWordSquare(words []string) bool {
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			if j >= len(words) || i >= len(words[j]) || words[i][j] != words[j][i] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(ValidWordSquare([]string{"abcd", "bnrt", "crmy", "dtyx"}))
	fmt.Println(ValidWordSquare([]string{"abcd", "bnrt", "crm", "dt"}))
	fmt.Println(ValidWordSquare([]string{"ball", "area", "lead", "lady"}))
}
