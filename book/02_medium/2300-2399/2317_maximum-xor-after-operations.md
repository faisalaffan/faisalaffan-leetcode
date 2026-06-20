# 2317 — Maximum Xor After Operations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumXOR(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2317: Maximum XOR After Operations
// https://leetcode.com/problems/maximum-xor-after-operations/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maximumXOR(nums []int) int {
	result := 0
	for _, v := range nums {
		result |= v
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(maximumXOR([]int{3, 2, 4, 6}))
	// Expected: 7

	// Test case 2
	fmt.Println(maximumXOR([]int{1, 2, 3, 4, 5, 6, 7}))
	// Expected: 7

	// Test case 3
	fmt.Println(maximumXOR([]int{0}))
	// Expected: 0
}
```
