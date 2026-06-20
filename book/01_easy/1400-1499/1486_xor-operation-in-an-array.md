# 1486 — Xor Operation In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func xorOperation(n int, start int) int

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1486: XOR Operation in an Array
// https://leetcode.com/problems/xor-operation-in-an-array/
// Difficulty: Easy
//
// LeetCode submission: func xorOperation(n int, start int) int

import "fmt"

func main() {
	fmt.Println(XorOperationInAnArray(5, 0)) // 8
	fmt.Println(XorOperationInAnArray(4, 3)) // 8
	fmt.Println(XorOperationInAnArray(1, 7)) // 7
}

// Time: O(n), Space: O(1)
func XorOperationInAnArray(n int, start int) int {
	res := 0
	for i := 0; i < n; i++ {
		res ^= start + 2*i
	}
	return res
}
```
