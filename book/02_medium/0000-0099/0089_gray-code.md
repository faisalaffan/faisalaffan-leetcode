# 0089 — Gray Code

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func grayCode(n int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Bitmask

**Kompleksitas Waktu:** O(2^n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Bitmask** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #89: Gray Code
// https://leetcode.com/problems/gray-code/
// Difficulty: Medium

import "fmt"

func grayCode(n int) []int {
  // Alokasi slice integer
	result := make([]int, 1<<n)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(result); i++ {
		result[i] = i ^ (i >> 1)
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(grayCode(2)) // [0 1 3 2]

	// Test case 2
	fmt.Println(grayCode(1)) // [0 1]

	// Test case 3
	fmt.Println(grayCode(3)) // [0 1 3 2 6 7 5 4]
}

// Time: O(2^n) | Space: O(1)
```
