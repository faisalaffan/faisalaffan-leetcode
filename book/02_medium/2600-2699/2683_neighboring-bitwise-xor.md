# 2683 — Neighboring Bitwise Xor

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func doesValidArrayExist(derived []int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2683: Neighboring Bitwise XOR
// https://leetcode.com/problems/neighboring-bitwise-xor/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func doesValidArrayExist(derived []int) bool {
	xor := 0
	for _, v := range derived {
		xor ^= v
	}
	return xor == 0
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", doesValidArrayExist([]int{1, 1, 0}))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", doesValidArrayExist([]int{1, 1}))
	// Expected: true

	// Test case 3
	fmt.Println("Test 3:", doesValidArrayExist([]int{1, 0}))
	// Expected: false
}
```
