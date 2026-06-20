# 3046 — Split The Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func SplitTheArray(nums []int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3046: Split the Array
// https://leetcode.com/problems/split-the-array/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: isPossibleToSplit
	fmt.Println(SplitTheArray([]int{1, 1, 2, 2, 3, 4})) // true
	fmt.Println(SplitTheArray([]int{1, 1, 1, 1}))       // false
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: isPossibleToSplit
// Each number can appear at most twice (once in each half of the split)
func SplitTheArray(nums []int) bool {
  // HashMap: O(1) lookup
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
		if freq[v] > 2 {
			return false
		}
	}
	return true
}
```
