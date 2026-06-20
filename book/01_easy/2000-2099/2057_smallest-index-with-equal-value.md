# 2057 — Smallest Index With Equal Value

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func SmallestIndexWithEqualValue(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2057: Smallest Index With Equal Value
// https://leetcode.com/problems/smallest-index-with-equal-value/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SmallestIndexWithEqualValue([]int{0, 1, 2}))          // 0
	fmt.Println(SmallestIndexWithEqualValue([]int{4, 3, 2, 1}))       // 2
	fmt.Println(SmallestIndexWithEqualValue([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})) // -1
}

// Time: O(n), Space: O(1)
func SmallestIndexWithEqualValue(nums []int) int {
	for i, v := range nums {
		if i%10 == v {
			return i
		}
	}
	return -1
}
```
