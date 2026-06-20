# 1313 — Decompress Run Length Encoded List

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func decompressRLElist(nums []int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n + totalLen), Space: O(totalLen)  |  **Ruang:** O(totalLen)


## 💻 Solusi Go

```go
package main

// LeetCode #1313: Decompress Run-Length Encoded List
// https://leetcode.com/problems/decompress-run-length-encoded-list/
// Difficulty: Easy
//
// LeetCode submission: func decompressRLElist(nums []int) []int

import "fmt"

func main() {
	fmt.Println(DecompressRunLengthEncodedList([]int{1, 2, 3, 4}))       // [2 4 4 4]
	fmt.Println(DecompressRunLengthEncodedList([]int{1, 1, 2, 3}))       // [1 3 3]
	fmt.Println(DecompressRunLengthEncodedList([]int{2, 5, 1, 7, 3, 9})) // [5 5 7 9 9 9]
}

// Time: O(n + totalLen), Space: O(totalLen)
func DecompressRunLengthEncodedList(nums []int) []int {
	totalLen := 0
  // Linear scan O(n)
	for i := 0; i < len(nums); i += 2 {
		totalLen += nums[i]
	}
  // Alokasi slice
	res := make([]int, 0, totalLen)
  // Linear scan O(n)
	for i := 0; i < len(nums); i += 2 {
		freq, val := nums[i], nums[i+1]
		for j := 0; j < freq; j++ {
			res = append(res, val)
		}
	}
	return res
}
```
