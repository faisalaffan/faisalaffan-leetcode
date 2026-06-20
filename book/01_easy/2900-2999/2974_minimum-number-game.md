# 2974 — Minimum Number Game

## Deskripsi

**Soal:** [2974. Minimum Number Game](https://leetcode.com/problems/minimum-number-game/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2974: Minimum Number Game
// https://leetcode.com/problems/minimum-number-game/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	// LeetCode name: numberGame
	fmt.Println(MinimumNumberGame([]int{5, 4, 2, 3})) // [3, 2, 5, 4]
	fmt.Println(MinimumNumberGame([]int{2, 5}))       // [5, 2]
}

// Time: O(n log n) | Space: O(n)
// LeetCode submission name: numberGame
func MinimumNumberGame(nums []int) []int {
	sort.Ints(nums)
  // Membuat slice untuk menyimpan hasil
	result := make([]int, len(nums))
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(nums); i += 2 {
		result[i] = nums[i+1]
		result[i+1] = nums[i]
	}
	return result
}
```
