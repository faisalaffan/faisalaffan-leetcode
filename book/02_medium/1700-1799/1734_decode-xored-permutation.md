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

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1734: Decode XORed Permutation
// https://leetcode.com/problems/decode-xored-permutation/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func decode(encoded []int) []int {
	n := len(encoded) + 1

	// XOR of all numbers from 1 to n
	totalXor := 0
	for i := 1; i <= n; i++ {
		totalXor ^= i
	}

	// XOR of encoded[1], encoded[3], encoded[5], ...
	xorOdd := 0
	for i := 1; i < len(encoded); i += 2 {
		xorOdd ^= encoded[i]
	}

	// First element = totalXor ^ xorOdd
  // Alokasi slice integer
	perm := make([]int, n)
	perm[0] = totalXor ^ xorOdd

	// Decode the rest
	for i := 1; i < n; i++ {
		perm[i] = perm[i-1] ^ encoded[i-1]
	}

	return perm
}

func main() {
	fmt.Println(decode([]int{3, 1}))          // Expected: [1, 2, 3]
	fmt.Println(decode([]int{6, 5, 4, 6}))    // Expected: [2, 4, 1, 5, 3]
	fmt.Println(decode([]int{5, 6, 1, 6, 2, 1})) // Expected: [1 4 2 3 5 7 6]
}
```
