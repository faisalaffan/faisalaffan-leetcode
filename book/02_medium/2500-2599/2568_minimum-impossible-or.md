# 2568 — Minimum Impossible Or

## Deskripsi

**Soal:** [2568. Minimum Impossible Or](https://leetcode.com/problems/minimum-impossible-or/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func minImpossibleOR(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #2568: Minimum Impossible OR
// https://leetcode.com/problems/minimum-impossible-or/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func minImpossibleOR(nums []int) int {
  // Membuat map untuk pencarian O(1): key → value
	seen := make(map[int]bool)
	for _, v := range nums {
		seen[v] = true
	}

	// Check powers of 2: 1, 2, 4, 8, ...
	// If any is missing, that's the answer (since it can't be formed by OR of smaller numbers)
	pow2 := 1
	for {
		if !seen[pow2] {
			return pow2
		}
		pow2 <<= 1
	}
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minImpossibleOR([]int{2, 1}))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", minImpossibleOR([]int{5, 3, 2}))
	// Expected: 1

	// Test case 3
	fmt.Println("Test 3:", minImpossibleOR([]int{1, 2, 4, 8}))
	// Expected: 16
}
```
