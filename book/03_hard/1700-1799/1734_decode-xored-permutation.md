# 1734 — Decode Xored Permutation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func decode(encoded []int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #1734: Decode XORed Permutation
// https://leetcode.com/problems/decode-xored-permutation/
// Difficulty: Medium (categorized as Hard in this repo)

import "fmt"

func main() {
	encoded1 := []int{3, 1}
	decoded1 := decode(encoded1)
	fmt.Printf("Test 1 - Input: %v\nOutput: %v\nExpected: [1,2,3]\n\n", encoded1, decoded1)

	encoded2 := []int{6, 5, 4, 6}
	decoded2 := decode(encoded2)
	fmt.Printf("Test 2 - Input: %v\nOutput: %v\nExpected: [2,4,1,5,3]\n\n", encoded2, decoded2)

	encoded3 := []int{12, 6, 2}
	decoded3 := decode(encoded3)
	fmt.Printf("Test 3 - Input: %v\nOutput: %v\n", encoded3, decoded3)
}

func decode(encoded []int) []int {
	n := len(encoded) + 1

	// XOR of all numbers from 1 to n
	allXor := 0
	for i := 1; i <= n; i++ {
		allXor ^= i
	}

	// XOR of encoded[1], encoded[3], encoded[5], ... (odd indices)
	// This gives us XOR of perm[1] ^ perm[2] ^ ... ^ perm[n-1]
	oddXor := 0
	for i := 1; i < len(encoded); i += 2 {
		oddXor ^= encoded[i]
	}

	// perm[0] = allXor ^ oddXor
  // Alokasi slice
	perm := make([]int, n)
	perm[0] = allXor ^ oddXor

	// Reconstruct the rest: perm[i] = perm[i-1] ^ encoded[i-1]
	for i := 1; i < n; i++ {
		perm[i] = perm[i-1] ^ encoded[i-1]
	}

	return perm
}
```
