package main

// LeetCode #17: Letter Combinations of a Phone Number
// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
// Difficulty: Medium

import "fmt"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	phone := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	result := []string{""}
	for i := 0; i < len(digits); i++ {
		letters := phone[digits[i]]
		var temp []string
		for _, prefix := range result {
			for j := 0; j < len(letters); j++ {
				temp = append(temp, prefix+string(letters[j]))
			}
		}
		result = temp
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(letterCombinations("23")) // ["ad","ae","af","bd","be","bf","cd","ce","cf"]

	// Test case 2
	fmt.Println(letterCombinations("")) // []

	// Test case 3
	fmt.Println(letterCombinations("2")) // ["a","b","c"]
}

// Time: O(4^n) | Space: O(4^n)
