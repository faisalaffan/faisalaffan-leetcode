# 0594 — Longest Harmonious Subsequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func LongestHarmoniousSubsequence(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #594: Longest Harmonious Subsequence
// https://leetcode.com/problems/longest-harmonious-subsequence/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func LongestHarmoniousSubsequence(nums []int) int {
  // HashMap: O(1) lookup
	count := make(map[int]int)
	for _, v := range nums {
		count[v]++
	}
	maxLen := 0
	for v, c := range count {
		if c2, ok := count[v+1]; ok {
			if c+c2 > maxLen {
				maxLen = c + c2
			}
		}
	}
	return maxLen
}

func main() {
	fmt.Println(LongestHarmoniousSubsequence([]int{1, 3, 2, 2, 5, 2, 3, 7}))
	fmt.Println(LongestHarmoniousSubsequence([]int{1, 2, 3, 4}))
	fmt.Println(LongestHarmoniousSubsequence([]int{1, 1, 1, 1}))
}
```
