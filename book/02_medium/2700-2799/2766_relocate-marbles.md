# 2766 — Relocate Marbles

## Deskripsi

**Soal:** [2766. Relocate Marbles](https://leetcode.com/problems/relocate-marbles/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func RelocateMarbles(nums []int, moveFrom []int, moveTo []int) []int`

## Solusi Go

```go
package main

// LeetCode #2766: Relocate Marbles
// https://leetcode.com/problems/relocate-marbles/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func RelocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {
  // Membuat map untuk pencarian O(1): key → value
	positions := make(map[int]bool)
	for _, n := range nums {
		positions[n] = true
	}

  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(moveFrom); i++ {
		delete(positions, moveFrom[i])
		positions[moveTo[i]] = true
	}

  // Membuat slice untuk menyimpan hasil
	result := make([]int, 0, len(positions))
	for p := range positions {
		result = append(result, p)
	}
	sort.Ints(result)
	return result
}

func main() {
	fmt.Println(RelocateMarbles([]int{1, 2, 3}, []int{1}, []int{4}))
	fmt.Println(RelocateMarbles([]int{1, 1, 2, 2}, []int{1, 2}, []int{3, 4}))
}
```
