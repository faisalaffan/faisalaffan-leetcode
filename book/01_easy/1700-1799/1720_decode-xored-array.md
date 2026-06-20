# 1720 — Decode Xored Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func Decode(encoded []int, first int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #1720: Decode XORed Array
// https://leetcode.com/problems/decode-xored-array/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func Decode(encoded []int, first int) []int {
  // Alokasi slice
	result := make([]int, len(encoded)+1)
	result[0] = first
  // Linear scan O(n)
	for i := 0; i < len(encoded); i++ {
		result[i+1] = result[i] ^ encoded[i]
	}
	return result
}

func main() {
	fmt.Println(Decode([]int{1, 2, 3}, 1))
	fmt.Println(Decode([]int{6, 2, 7, 3}, 4))
}
```
