# 1562 — Find Latest Group Of Size M

## Deskripsi

**Soal:** [1562. Find Latest Group Of Size M](https://leetcode.com/problems/find-latest-group-of-size-m/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N), Space: O(N)  
**Kompleksitas Ruang:** O(N)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1562: Find Latest Group of Size M
// https://leetcode.com/problems/find-latest-group-of-size-m/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FindLatestStep([]int{3, 5, 1, 2, 4}, 1))
	fmt.Println(FindLatestStep([]int{3, 1, 5, 4, 2}, 2))
	fmt.Println(FindLatestStep([]int{1}, 1))
}

func FindLatestStep(arr []int, m int) int {
	// Time: O(N), Space: O(N)
	// Use union-find-like approach tracking group lengths
	n := len(arr)
	if n == m {
		return n
	}

  // Membuat slice untuk menyimpan hasil
	length := make([]int, n+2) // length of group at position i
  // Membuat slice untuk menyimpan hasil
	count := make([]int, n+2)  // count of groups of length i
	result := -1

	for step := 0; step < n; step++ {
		pos := arr[step]
		leftLen := length[pos-1]
		rightLen := length[pos+1]
		total := leftLen + rightLen + 1

		// Decrement counts for the merging groups
		count[leftLen]--
		count[rightLen]--
		// Increment count for the new merged group
		count[total]++

		// Update lengths at boundaries
		length[pos-leftLen] = total
		length[pos+rightLen] = total

		// Check if we have exactly m groups of size m
		if count[m] > 0 {
			// This step is valid (but since we want the latest step before m disappears,
			// we'll track the result after the step)
		}

		// The result is the latest step where count[m] > 0
		// But we want the latest step where a group of size m *exists*
		// We need to check AFTER the current operation
	}

	// Re-simulate to find latest step with count[m] > 0
	length = make([]int, n+2)
	count = make([]int, n+2)

	for step := 0; step < n; step++ {
		pos := arr[step]
		leftLen := length[pos-1]
		rightLen := length[pos+1]
		total := leftLen + rightLen + 1

		count[leftLen]--
		count[rightLen]--
		count[total]++

		length[pos-leftLen] = total
		length[pos+rightLen] = total

		if count[m] > 0 {
			result = step + 1 // 1-indexed step
		}
	}

	return result
}
```
