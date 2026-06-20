# 1497 — Check If Array Pairs Are Divisible By K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func CanArrange(arr []int, k int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(N), Space: O(K)  |  **Ruang:** O(K)


## 💻 Solusi Go

```go
package main

// LeetCode #1497: Check If Array Pairs Are Divisible by k
// https://leetcode.com/problems/check-if-array-pairs-are-divisible-by-k/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(CanArrange([]int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}, 5))
	fmt.Println(CanArrange([]int{1, 2, 3, 4, 5, 6}, 7))
	fmt.Println(CanArrange([]int{1, 2, 3, 4, 5, 6}, 10))
}

func CanArrange(arr []int, k int) bool {
	// Time: O(N), Space: O(K)
  // Alokasi slice
	remainder := make([]int, k)
	for _, num := range arr {
		r := ((num % k) + k) % k
		remainder[r]++
	}

	// Numbers divisible by k must pair among themselves
	if remainder[0]%2 != 0 {
		return false
	}

	// For i and k-i, their counts must match
	for i := 1; i < k; i++ {
		if remainder[i] != remainder[k-i] {
			return false
		}
	}

	return true
}
```
