package main

// LeetCode #2024: Maximize the Confusion of an Exam
// https://leetcode.com/problems/maximize-the-confusion-of-an-exam/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maxConsecutiveAnswers(answerKey string, k int) int {
	return max(
		longestChar(answerKey, k, 'T'),
		longestChar(answerKey, k, 'F'),
	)
}

func longestChar(s string, k int, target byte) int {
	left := 0
	flips := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		if s[right] != target {
			flips++
		}
		for flips > k {
			if s[left] != target {
				flips--
			}
			left++
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxConsecutiveAnswers("TTFF", 2))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", maxConsecutiveAnswers("TFFT", 1))
	// Expected: 3

	// Test case 3
	fmt.Println("Test 3:", maxConsecutiveAnswers("TTFTTFTT", 1))
	// Expected: 5
}
