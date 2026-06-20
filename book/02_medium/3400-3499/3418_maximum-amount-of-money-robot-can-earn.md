# 3418 — Maximum Amount Of Money Robot Can Earn

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumAmount(coins [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m*n) Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3418: Maximum Amount of Money Robot Can Earn
// https://leetcode.com/problems/maximum-amount-of-money-robot-can-earn/
// Difficulty: Medium
// Time: O(m*n) Space: O(n)

import (
	"fmt"
	"math"
)

func maximumAmount(coins [][]int) int {
	n := len(coins[0])
  // Alokasi slice integer
	f := make([][3]int, n+1)
	for j := range f {
		f[j] = [3]int{math.MinInt / 2, math.MinInt / 2, math.MinInt / 2}
	}
	f[1] = [3]int{}

	for _, row := range coins {
		for j, x := range row {
			f[j+1][2] = max(f[j][2]+x, f[j+1][2]+x,
				max(f[j][1], f[j+1][1]))
			f[j+1][1] = max(f[j][1]+x, f[j+1][1]+x,
				max(f[j][0], f[j+1][0]))
			f[j+1][0] = max(f[j][0], f[j+1][0]) + x
		}
	}
	return f[n][2]
}

func main() {
	fmt.Println(maximumAmount([][]int{{0, 1, -1}, {1, -2, 3}, {2, -3, 4}})) // 8
	fmt.Println(maximumAmount([][]int{{10, 10, 10}, {10, 10, 10}}))        // 40
}
```
