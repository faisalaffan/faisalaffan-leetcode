package main

// LeetCode #1395: Count Number of Teams
// https://leetcode.com/problems/count-number-of-teams/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(numTeams([]int{2, 5, 3, 4, 1})) // 3

	// Test case 2
	fmt.Println(numTeams([]int{2, 1, 3})) // 0

	// Test case 3
	fmt.Println(numTeams([]int{1, 2, 3, 4})) // 4
}

// Time: O(n^2) using middle element approach
// Space: O(1)
func numTeams(rating []int) int {
	n := len(rating)
	count := 0

	for j := 1; j < n-1; j++ {
		// Count elements smaller/larger on left
		leftSmaller, leftLarger := 0, 0
		for i := 0; i < j; i++ {
			if rating[i] < rating[j] {
				leftSmaller++
			} else if rating[i] > rating[j] {
				leftLarger++
			}
		}

		// Count elements smaller/larger on right
		rightSmaller, rightLarger := 0, 0
		for k := j + 1; k < n; k++ {
			if rating[k] < rating[j] {
				rightSmaller++
			} else if rating[k] > rating[j] {
				rightLarger++
			}
		}

		// Increasing: leftSmaller * rightLarger
		// Decreasing: leftLarger * rightSmaller
		count += leftSmaller*rightLarger + leftLarger*rightSmaller
	}

	return count
}
