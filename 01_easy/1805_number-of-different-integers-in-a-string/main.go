package main

// LeetCode #1805: Number of Different Integers in a String
// https://leetcode.com/problems/number-of-different-integers-in-a-string/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func NumDifferentIntegers(word string) int {
	seen := make(map[string]bool)
	i := 0
	for i < len(word) {
		if word[i] >= '0' && word[i] <= '9' {
			j := i
			for j < len(word) && word[j] >= '0' && word[j] <= '9' {
				j++
			}
			for i < j && word[i] == '0' {
				i++
			}
			num := word[i:j]
			if !seen[num] {
				seen[num] = true
			}
			i = j
		} else {
			i++
		}
	}
	return len(seen)
}

func main() {
	fmt.Println(NumDifferentIntegers("a123bc34d8ef34"))
	fmt.Println(NumDifferentIntegers("leet1234code234"))
	fmt.Println(NumDifferentIntegers("a1b01c001"))
}
