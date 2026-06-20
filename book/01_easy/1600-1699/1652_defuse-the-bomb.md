# 1652 — Defuse The Bomb

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func Decrypt(code []int, k int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n*  |  **Ruang:** O(n) (or O(1) excluding output)


## 💻 Solusi Go

```go
package main

// LeetCode #1652: Defuse the Bomb
// https://leetcode.com/problems/defuse-the-bomb/
// Difficulty: Easy

import "fmt"

// Time: O(n*|k|), Space: O(n) (or O(1) excluding output)
func Decrypt(code []int, k int) []int {
	n := len(code)
  // Alokasi slice
	result := make([]int, n)
	if k == 0 {
		return result
	}
	for i := 0; i < n; i++ {
		sum := 0
		if k > 0 {
			for j := 1; j <= k; j++ {
				sum += code[(i+j)%n]
			}
		} else {
			for j := 1; j <= -k; j++ {
				sum += code[(i-j+n)%n]
			}
		}
		result[i] = sum
	}
	return result
}

func main() {
	fmt.Println(Decrypt([]int{5, 7, 1, 4}, 3))
	fmt.Println(Decrypt([]int{1, 2, 3, 4}, 0))
	fmt.Println(Decrypt([]int{2, 4, 9, 3}, -2))
}
```
