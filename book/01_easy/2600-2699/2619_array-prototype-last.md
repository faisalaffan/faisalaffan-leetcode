# 2619 — Array Prototype Last

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func arrayPrototypeLast(arr []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2619: Array Prototype Last
// https://leetcode.com/problems/array-prototype-last/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Returns last element of a slice.

import "fmt"

func main() {
	fmt.Println(arrayPrototypeLast([]int{1, 2, 3}))
	fmt.Println(arrayPrototypeLast([]int{}))
}

func arrayPrototypeLast(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	return arr[len(arr)-1]
}
```
