package main

// LeetCode #848: Shifting Letters
// https://leetcode.com/problems/shifting-letters/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(ShiftingLetters("abc", []int{3, 5, 9}))
	fmt.Println(ShiftingLetters("aaa", []int{1, 2, 3}))
	fmt.Println(ShiftingLetters("z", []int{52}))
}

// Time: O(n) | Space: O(n)
func ShiftingLetters(s string, shifts []int) string {
	n := len(s)
	// Calculate suffix sum of shifts
	for i := n - 2; i >= 0; i-- {
		shifts[i] = (shifts[i] + shifts[i+1]) % 26
	}

	res := []byte(s)
	for i := 0; i < n; i++ {
		res[i] = byte((int(res[i]-'a')+shifts[i])%26 + 'a')
	}

	return string(res)
}
