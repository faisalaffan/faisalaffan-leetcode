# 0119 — Pascals Triangle Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func GetRow(rowIndex int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(rowIndex^2)  
**Kompleksitas Ruang:** O(rowIndex)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #119: Pascal's Triangle II
// https://leetcode.com/problems/pascals-triangle-ii/
// Difficulty: Easy

import "fmt"

// Time: O(rowIndex^2) | Space: O(rowIndex)
func GetRow(rowIndex int) []int {
  // Alokasi slice integer
	res := make([]int, rowIndex+1)
	res[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			res[j] += res[j-1]
		}
	}
	return res
}

func main() {
	fmt.Println(GetRow(3))
	fmt.Println(GetRow(0))
	fmt.Println(GetRow(4))
}
```
