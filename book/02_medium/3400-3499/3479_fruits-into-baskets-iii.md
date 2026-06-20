# 3479 — Fruits Into Baskets Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FruitsIntoBasketsIii(fruits []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Two Pointer

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3479: Fruits Into Baskets III
// https://leetcode.com/problems/fruits-into-baskets-iii/
// Difficulty: Medium
// Complexity: O(n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", FruitsIntoBasketsIii([]int{1, 2, 1}))
	// Test case 2
	fmt.Println("Test 2:", FruitsIntoBasketsIii([]int{0, 1, 2, 2}))
	// Test case 3
	fmt.Println("Test 3:", FruitsIntoBasketsIii([]int{1, 2, 3, 2, 2}))
}

func FruitsIntoBasketsIii(fruits []int) int {
  // HashMap: O(1) lookup
	freq := make(map[int]int)
	left := 0
	maxLen := 0
	for right := 0; right < len(fruits); right++ {
		freq[fruits[right]]++
		for len(freq) > 2 {
			freq[fruits[left]]--
			if freq[fruits[left]] == 0 {
				delete(freq, fruits[left])
			}
			left++
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}
```
