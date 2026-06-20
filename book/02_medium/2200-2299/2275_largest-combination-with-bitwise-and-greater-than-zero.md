# 2275 — Largest Combination With Bitwise And Greater Than Zero

## Deskripsi

**Soal:** [2275. Largest Combination With Bitwise And Greater Than Zero](https://leetcode.com/problems/largest-combination-with-bitwise-and-greater-than-zero/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * 24)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func largestCombination(candidates []int) int`

## Solusi Go

```go
package main

// LeetCode #2275: Largest Combination With Bitwise AND Greater Than Zero
// https://leetcode.com/problems/largest-combination-with-bitwise-and-greater-than-zero/
// Difficulty: Medium
// Time: O(n * 24) | Space: O(1)

import "fmt"

func largestCombination(candidates []int) int {
	maxCount := 0
	for bit := 0; bit < 24; bit++ {
		count := 0
		for _, c := range candidates {
			if c&(1<<bit) != 0 {
				count++
			}
		}
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}

func main() {
	// Test case 1
	fmt.Println(largestCombination([]int{16, 17, 71, 62, 12, 24, 14}))
	// Expected: 4

	// Test case 2
	fmt.Println(largestCombination([]int{8, 8}))
	// Expected: 2

	// Test case 3
	fmt.Println(largestCombination([]int{1, 2, 4, 8}))
	// Expected: 1
}
```
