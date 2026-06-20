# 1658 — Minimum Operations To Reduce X To Zero

## Deskripsi

**Soal:** [1658. Minimum Operations To Reduce X To Zero](https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1658: Minimum Operations to Reduce X to Zero
// https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinOperations([]int{1, 1, 4, 2, 3}, 5))
	fmt.Println(MinOperations([]int{5, 6, 7, 8, 9}, 4))
	fmt.Println(MinOperations([]int{3, 2, 20, 1, 1, 3}, 10))
}

func MinOperations(nums []int, x int) int {
	// Time: O(N), Space: O(1)
	// Find longest subarray with sum = total - x
	total := 0
	for _, num := range nums {
		total += num
	}

	target := total - x
	if target < 0 {
		return -1
	}
	if target == 0 {
		return len(nums)
	}

	maxLen := -1
	left := 0
	currSum := 0

	for right := 0; right < len(nums); right++ {
		currSum += nums[right]

		for currSum > target {
			currSum -= nums[left]
			left++
		}

		if currSum == target {
			if right-left+1 > maxLen {
				maxLen = right - left + 1
			}
		}
	}

	if maxLen == -1 {
		return -1
	}
	return len(nums) - maxLen
}
```
