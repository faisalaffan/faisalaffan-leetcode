# 2433 — Find The Original Array Of Prefix Xor

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func findArray(pref []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(n)  |  **Ruang:** O(1) (excluding output)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2433: Find The Original Array of Prefix Xor
// https://leetcode.com/problems/find-the-original-array-of-prefix-xor/
// Difficulty: Medium
// Time: O(n) | Space: O(1) (excluding output)
// Given pref[i] = XOR(arr[0..i]), find arr. arr[i] = pref[i] ^ pref[i-1].

import "fmt"

func main() {
	fmt.Println(findArray([]int{5, 2, 0, 3, 1})) // [5, 7, 2, 3, 2]
	fmt.Println(findArray([]int{13}))             // [13]
}

func findArray(pref []int) []int {
	n := len(pref)
  // Alokasi slice
	arr := make([]int, n)
	arr[0] = pref[0]
	for i := 1; i < n; i++ {
		arr[i] = pref[i] ^ pref[i-1]
	}
	return arr
}
```
