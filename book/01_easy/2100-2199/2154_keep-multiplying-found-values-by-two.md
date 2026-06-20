# 2154 — Keep Multiplying Found Values By Two

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func KeepMultiplyingFoundValuesByTwo(nums []int, original int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2154: Keep Multiplying Found Values by Two
// https://leetcode.com/problems/keep-multiplying-found-values-by-two/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(KeepMultiplyingFoundValuesByTwo([]int{5, 3, 6, 1, 12}, 3))  // 24
	fmt.Println(KeepMultiplyingFoundValuesByTwo([]int{2, 7, 9}, 4))          // 4
}

// Time: O(n), Space: O(n)
func KeepMultiplyingFoundValuesByTwo(nums []int, original int) int {
  // HashMap: O(1) lookup
	set := make(map[int]bool)
	for _, v := range nums {
		set[v] = true
	}

	for set[original] {
		original *= 2
	}
	return original
}
```
