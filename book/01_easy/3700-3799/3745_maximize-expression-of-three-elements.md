# 3745 — Maximize Expression Of Three Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximizeExpressionOfThreeElements(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3745: Maximize Expression of Three Elements
// https://leetcode.com/problems/maximize-expression-of-three-elements/
// Difficulty: Easy

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(MaximizeExpressionOfThreeElements([]int{1, 4, 2, 5}))
	fmt.Println(MaximizeExpressionOfThreeElements([]int{-2, 0, 5, -2, 4}))
}

// Time: O(n)
// Space: O(1)
func MaximizeExpressionOfThreeElements(nums []int) int {
	max1, max2 := math.MinInt32, math.MinInt32
	min1 := math.MaxInt32

	for _, num := range nums {
		if num > max1 {
			max2 = max1
			max1 = num
		} else if num > max2 {
			max2 = num
		}
		if num < min1 {
			min1 = num
		}
	}

	return max1 + max2 - min1
}
```
