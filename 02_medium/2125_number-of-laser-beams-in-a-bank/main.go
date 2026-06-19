package main

// LeetCode #2125: Number of Laser Beams in a Bank
// https://leetcode.com/problems/number-of-laser-beams-in-a-bank/
// Difficulty: Medium
// Time: O(m*n) | Space: O(1)

import "fmt"

func numberOfBeams(bank []string) int {
	prevCount := 0
	result := 0

	for _, row := range bank {
		count := 0
		for _, c := range row {
			if c == '1' {
				count++
			}
		}
		if count > 0 {
			result += prevCount * count
			prevCount = count
		}
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", numberOfBeams([]string{"011001", "000000", "010100", "001000"}))
	// Expected: 8

	// Test case 2
	fmt.Println("Test 2:", numberOfBeams([]string{"000", "111", "000"}))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", numberOfBeams([]string{"101", "010", "101"}))
	// Expected: 4
}
