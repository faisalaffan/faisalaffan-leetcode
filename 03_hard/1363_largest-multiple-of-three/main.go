package main

// LeetCode #1363: Largest Multiple of Three
// https://leetcode.com/problems/largest-multiple-of-three/
// Difficulty: Hard

import (
	"fmt"
)

func largestMultipleOfThree(digits []int) string {
	// Count digits
	count := make([]int, 10)
	sum := 0
	for _, d := range digits {
		count[d]++
		sum += d
	}

	// Remainder map: remainder -> digits to remove (1 or 2 digits)
	rem1 := []int{1, 4, 7}
	rem2 := []int{2, 5, 8}

	rem := sum % 3
	if rem == 1 {
		// Try removing 1 digit with remainder 1
		if !remove(count, rem1) {
			// Remove 2 digits with remainder 2
			remove2(count, rem2)
		}
	} else if rem == 2 {
		// Try removing 1 digit with remainder 2
		if !remove(count, rem2) {
			// Remove 2 digits with remainder 1
			remove2(count, rem1)
		}
	}

	// Build result
	var result []byte
	for d := 9; d >= 0; d-- {
		for i := 0; i < count[d]; i++ {
			result = append(result, byte('0'+d))
		}
	}

	if len(result) == 0 {
		return ""
	}
	if result[0] == '0' {
		return "0"
	}
	return string(result)
}

// remove removes the first digit from candidates that has count > 0
func remove(count []int, candidates []int) bool {
	for _, d := range candidates {
		if count[d] > 0 {
			count[d]--
			return true
		}
	}
	return false
}

// remove2 removes two digits from candidates
func remove2(count []int, candidates []int) bool {
	removed := 0
	for _, d := range candidates {
		for count[d] > 0 && removed < 2 {
			count[d]--
			removed++
		}
		if removed == 2 {
			return true
		}
	}
	return false
}

func main() {
	// Example 1
	fmt.Println(largestMultipleOfThree([]int{8, 1, 9}))
	// Expected: "981"

	// Example 2
	fmt.Println(largestMultipleOfThree([]int{8, 6, 7, 1, 0}))
	// Expected: "8760"

	// Example 3
	fmt.Println(largestMultipleOfThree([]int{1}))
	// Expected: ""

	// Example 4: all zeros
	fmt.Println(largestMultipleOfThree([]int{0, 0, 0}))
	// Expected: "0"
}
