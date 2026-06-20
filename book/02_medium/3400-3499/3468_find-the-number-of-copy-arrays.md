# 3468 — Find The Number Of Copy Arrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func countArrays(original []int, bounds [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n) Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3468: Find the Number of Copy Arrays
// https://leetcode.com/problems/find-the-number-of-copy-arrays/
// Difficulty: Medium
// Time: O(n) Space: O(1)

import (
	"fmt"
	"math"
)

func countArrays(original []int, bounds [][]int) int {
	n := len(original)
	lo := math.MinInt64
	hi := math.MaxInt64
	for i := 0; i < n; i++ {
		diff := original[i] - original[0]
		l := bounds[i][0] - diff
		r := bounds[i][1] - diff
		if l > lo {
			lo = l
		}
		if r < hi {
			hi = r
		}
	}
	if hi < lo {
		return 0
	}
	return hi - lo + 1
}

func main() {
	fmt.Println(countArrays([]int{1, 2, 3, 4}, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}})) // 2
	fmt.Println(countArrays([]int{1, 2, 3}, [][]int{{1, 3}, {2, 4}, {3, 5}})) // 3
}
```
