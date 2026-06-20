# 3184 — Count Pairs That Form A Complete Day I

## Deskripsi

**Soal:** [3184. Count Pairs That Form A Complete Day I](https://leetcode.com/problems/count-pairs-that-form-a-complete-day-i/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n). Space: O(24).  
**Kompleksitas Ruang:** O(24).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3184: Count Pairs That Form a Complete Day I
// https://leetcode.com/problems/count-pairs-that-form-a-complete-day-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountPairsThatFormACompleteDayI([]int{12, 12, 30, 24, 24}))
	fmt.Println(CountPairsThatFormACompleteDayI([]int{72, 48, 24, 3}))
}

// CountPairsThatFormACompleteDayI counts pairs (i, j) where i < j and hours[i] + hours[j] is divisible by 24.
// Time: O(n). Space: O(24).
func CountPairsThatFormACompleteDayI(hours []int) int {
	count := 0
  // Membuat slice untuk menyimpan hasil
	rem := make([]int, 24)
	for _, h := range hours {
		r := h % 24
		need := (24 - r) % 24
		count += rem[need]
		rem[r]++
	}
	return count
}
```
