# 2784 — Check If Array Is Good

## Deskripsi

**Soal:** [2784. Check If Array Is Good](https://leetcode.com/problems/check-if-array-is-good/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2784: Check if Array is Good
// https://leetcode.com/problems/check-if-array-is-good/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(CheckIfArrayIsGood([]int{2, 1, 3}))
	fmt.Println(CheckIfArrayIsGood([]int{3, 4, 4, 1, 2, 1}))
}

func CheckIfArrayIsGood(nums []int) bool {
	n := len(nums) - 1
  // Membuat slice untuk menyimpan hasil
	counts := make([]int, n+1)
	for _, v := range nums {
		if v > n {
			return false
		}
		counts[v]++
	}
	if counts[n] != 2 {
		return false
	}
	for i := 1; i < n; i++ {
		if counts[i] != 1 {
			return false
		}
	}
	return true
}
```
