# 2611 — Mice And Cheese

## Deskripsi

**Soal:** [2611. Mice And Cheese](https://leetcode.com/problems/mice-and-cheese/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func miceAndCheese(reward1 []int, reward2 []int, k int) int`

## Solusi Go

```go
package main

// LeetCode #2611: Mice and Cheese
// https://leetcode.com/problems/mice-and-cheese/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	n := len(reward1)
  // Membuat slice untuk menyimpan hasil
	diff := make([]int, n)
	for i := 0; i < n; i++ {
		diff[i] = reward1[i] - reward2[i]
	}

	// We need to pick k indices for mouse 1 to maximize total
	// total = sum(reward2) + sum of top k diffs
	total := 0
	for _, v := range reward2 {
		total += v
	}

	sort.Slice(diff, func(i, j int) bool {
		return diff[i] > diff[j]
	})

	for i := 0; i < k; i++ {
		total += diff[i]
	}
	return total
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", miceAndCheese([]int{1, 1, 3, 4}, []int{4, 4, 1, 1}, 2))
	// Expected: 15

	// Test case 2
	fmt.Println("Test 2:", miceAndCheese([]int{1, 1}, []int{1, 1}, 2))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", miceAndCheese([]int{2, 3, 5}, []int{5, 2, 1}, 1))
	// Expected: 12
}
```
