# 3718 — Smallest Missing Multiple Of K

## Deskripsi

**Soal:** [3718. Smallest Missing Multiple Of K](https://leetcode.com/problems/smallest-missing-multiple-of-k/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n + max_missing/k)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3718: Smallest Missing Multiple of K
// https://leetcode.com/problems/smallest-missing-multiple-of-k/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SmallestMissingMultipleOfK([]int{8, 2, 3, 4, 6}, 2))
	fmt.Println(SmallestMissingMultipleOfK([]int{1, 4, 7, 10, 15}, 5))
}

// Time: O(n + max_missing/k)
// Space: O(n)
func SmallestMissingMultipleOfK(nums []int, k int) int {
  // Membuat map untuk pencarian O(1): key → value
	has := make(map[int]bool)
	for _, v := range nums {
		has[v] = true
	}

	for x := k; ; x += k {
		if !has[x] {
			return x
		}
	}
}
```
