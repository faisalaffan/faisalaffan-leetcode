# 2829 — Determine The Minimum Sum Of A K Avoiding Array

## Deskripsi

**Soal:** [2829. Determine The Minimum Sum Of A K Avoiding Array](https://leetcode.com/problems/determine-the-minimum-sum-of-a-k-avoiding-array/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func DetermineTheMinimumSumOfAKAvoidingArray(n int, k int) int`

## Solusi Go

```go
package main

// LeetCode #2829: Determine the Minimum Sum of a k-avoiding Array
// https://leetcode.com/problems/determine-the-minimum-sum-of-a-k-avoiding-array/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func DetermineTheMinimumSumOfAKAvoidingArray(n int, k int) int {
  // Membuat map untuk pencarian O(1): key → value
	used := make(map[int]bool)
	sum := 0
	for i := 1; len(used) < n; i++ {
		if used[k-i] {
			continue
		}
		used[i] = true
		sum += i
	}
	return sum
}

func main() {
	fmt.Println(DetermineTheMinimumSumOfAKAvoidingArray(5, 4))
	fmt.Println(DetermineTheMinimumSumOfAKAvoidingArray(3, 5))
}
```
