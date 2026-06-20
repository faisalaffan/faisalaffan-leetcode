# 3834 — Merge Adjacent Equal Elements

## Deskripsi

**Soal:** [3834. Merge Adjacent Equal Elements](https://leetcode.com/problems/merge-adjacent-equal-elements/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(N)

**Algoritma:** Stack (tumpukan LIFO)

**Fungsi Solusi:** `func MergeAdjacentEqualElements(nums []int) []int64`

> **Ide Kunci:** Use a stack to repeatedly merge leftmost adjacent equal pairs.

## Solusi Go

```go
package main

// LeetCode #3834: Merge Adjacent Equal Elements
// https://leetcode.com/problems/merge-adjacent-equal-elements/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Use a stack to repeatedly merge leftmost adjacent equal pairs.

import "fmt"

func MergeAdjacentEqualElements(nums []int) []int64 {
  // Membuat slice untuk menyimpan hasil
	stack := make([]int64, 0)

	for _, v := range nums {
		cur := int64(v)
		// While stack top equals cur, merge (pop + double)
		for len(stack) > 0 && stack[len(stack)-1] == cur {
			cur += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, cur)
	}

	return stack
}

func main() {
	// Example 1
	fmt.Println(MergeAdjacentEqualElements([]int{3, 1, 1, 2})) // Expected: [3 4]

	// Example 2
	fmt.Println(MergeAdjacentEqualElements([]int{2, 2, 4})) // Expected: [8]

	// Example 3
	fmt.Println(MergeAdjacentEqualElements([]int{3, 7, 5})) // Expected: [3 7 5]
}
```
