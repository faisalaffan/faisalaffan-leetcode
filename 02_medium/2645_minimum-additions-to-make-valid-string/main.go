package main

// LeetCode #2645: Minimum Additions to Make Valid String
// https://leetcode.com/problems/minimum-additions-to-make-valid-string/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func addMinimum(word string) int {
	n := len(word)
	count := 0
	i := 0

	// Pattern we're building: "abc"
	expected := []byte{'a', 'b', 'c'}
	expectedIdx := 0

	for i < n {
		if word[i] == expected[expectedIdx] {
			i++
		} else {
			count++
		}
		expectedIdx = (expectedIdx + 1) % 3
	}

	// Finish the current "abc" cycle
	for expectedIdx != 0 {
		count++
		expectedIdx = (expectedIdx + 1) % 3
	}

	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", addMinimum("b"))
	// Expected: 2 (add "a" and "c" to make "abc")

	// Test case 2
	fmt.Println("Test 2:", addMinimum("aaa"))
	// Expected: 6 (add "bc" twice)

	// Test case 3
	fmt.Println("Test 3:", addMinimum("abc"))
	// Expected: 0
}
