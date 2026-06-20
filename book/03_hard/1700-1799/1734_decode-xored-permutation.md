# 1734 — Decode Xored Permutation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func decode(encoded []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
  // Alokasi slice integer
	perm := make([]int, n)
	perm[0] = allXor ^ oddXor

	// Reconstruct the rest: perm[i] = perm[i-1] ^ encoded[i-1]
	for i := 1; i < n; i++ {
		perm[i] = perm[i-1] ^ encoded[i-1]
	}

	return perm
}
```
