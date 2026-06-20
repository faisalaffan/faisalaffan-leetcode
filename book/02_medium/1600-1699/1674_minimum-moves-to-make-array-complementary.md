# 1674 — Minimum Moves To Make Array Complementary

## Deskripsi

**Soal:** [1674. Minimum Moves To Make Array Complementary](https://leetcode.com/problems/minimum-moves-to-make-array-complementary/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n + limit), Space: O(limit)  
**Kompleksitas Ruang:** O(limit)

**Algoritma:** —

**Fungsi Solusi:** `func minMoves(nums []int, limit int) int`

## Solusi Go

```go
package main

// LeetCode #1674: Minimum Moves to Make Array Complementary
// https://leetcode.com/problems/minimum-moves-to-make-array-complementary/
// Difficulty: Medium
// Time: O(n + limit), Space: O(limit)

import "fmt"

func minMoves(nums []int, limit int) int {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	delta := make([]int, 2*limit+2) // difference array

	for i := 0; i < n/2; i++ {
		a, b := nums[i], nums[n-1-i]
		// Pair sum
		pairSum := a + b
		// min achievable sum: min(a,b) + 1
		// max achievable sum: max(a,b) + limit

		// 0 changes: [pairSum, pairSum] is already achievable
		// 1 change: [min(a,b)+1, max(a,b)+limit]
		// 2 changes: everything else [2, 2*limit]

		delta[2] += 2
		delta[min(a, b)+1]-- // Reduce to 1 change at this range start
		delta[pairSum]--     // Make it 0 changes at pairSum
		delta[pairSum+1]++   // Back to 1 change after pairSum
		delta[max(a, b)+limit+1]++ // Back to 2 changes
	}

	ans := n
	cur := 0
	for target := 2; target <= 2*limit; target++ {
		cur += delta[target]
		if cur < ans {
			ans = cur
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minMoves([]int{1, 2, 4, 3}, 4)) // Expected: 1

	// Test case 2
	fmt.Println("Test 2:", minMoves([]int{1, 2, 2, 1}, 2)) // Expected: 2

	// Test case 3
	fmt.Println("Test 3:", minMoves([]int{1, 1, 1, 1}, 3)) // Expected: 0 (all pairs already sum to 2)
}
```
