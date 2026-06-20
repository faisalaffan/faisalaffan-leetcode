# 1536 — Minimum Swaps To Arrange A Binary Grid

## Deskripsi

**Soal:** [1536. Minimum Swaps To Arrange A Binary Grid](https://leetcode.com/problems/minimum-swaps-to-arrange-a-binary-grid/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N^2), Space: O(N)  
**Kompleksitas Ruang:** O(N)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1536: Minimum Swaps to Arrange a Binary Grid
// https://leetcode.com/problems/minimum-swaps-to-arrange-a-binary-grid/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinSwaps([][]int{{0, 0, 1}, {1, 1, 0}, {1, 0, 0}}))
	fmt.Println(MinSwaps([][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {0, 1, 1, 0}, {0, 1, 1, 0}}))
	fmt.Println(MinSwaps([][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 1}}))
}

func MinSwaps(grid [][]int) int {
	// Time: O(N^2), Space: O(N)
	n := len(grid)

	// trailingZeros[i] = number of trailing zeros in row i
  // Membuat slice untuk menyimpan hasil
	trailingZeros := make([]int, n)
	for i := 0; i < n; i++ {
		count := 0
		for j := n - 1; j >= 0 && grid[i][j] == 0; j-- {
			count++
		}
		trailingZeros[i] = count
	}

	swaps := 0

	for i := 0; i < n; i++ {
		// Row i needs at least n-i-1 trailing zeros
		needed := n - i - 1
		found := -1

		for j := i; j < n; j++ {
			if trailingZeros[j] >= needed {
				found = j
				break
			}
		}

		if found == -1 {
			return -1
		}

		// Bubble the found row up to position i
		for j := found; j > i; j-- {
			trailingZeros[j], trailingZeros[j-1] = trailingZeros[j-1], trailingZeros[j]
			swaps++
		}
	}

	return swaps
}
```
