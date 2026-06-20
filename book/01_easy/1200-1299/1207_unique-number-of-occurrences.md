# 1207 — Unique Number Of Occurrences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func uniqueOccurrences(arr []int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1207: Unique Number of Occurrences
// https://leetcode.com/problems/unique-number-of-occurrences/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3})) // true
	fmt.Println(uniqueOccurrences([]int{1, 2}))             // false
	fmt.Println(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0})) // true
}

// LeetCode submission: uniqueOccurrences
func uniqueOccurrences(arr []int) bool {
  // HashMap: O(1) lookup
	freq := make(map[int]int)
	for _, v := range arr {
		freq[v]++
	}
  // HashMap: O(1) lookup
	seen := make(map[int]bool)
	for _, f := range freq {
		if seen[f] {
			return false
		}
		seen[f] = true
	}
	return true
}
```
