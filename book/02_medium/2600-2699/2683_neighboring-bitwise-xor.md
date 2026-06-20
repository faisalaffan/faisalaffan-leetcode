# 2683 — Neighboring Bitwise Xor

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func doesValidArrayExist(derived []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
