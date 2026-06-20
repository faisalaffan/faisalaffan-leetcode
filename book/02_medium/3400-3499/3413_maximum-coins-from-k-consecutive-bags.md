# 3413 — Maximum Coins From K Consecutive Bags

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumWhiteTiles(tiles [][]int, carpetLen int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n log n) Space: O(log n)  
**Kompleksitas Ruang:** O(log n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3413: Maximum Coins From K Consecutive Bags
// https://leetcode.com/problems/maximum-coins-from-k-consecutive-bags/
// Difficulty: Medium
// Time: O(n log n) Space: O(log n)

import (
	"fmt"
	"slices"
)

func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
	cover, left := 0, 0
	ans := 0
	for _, tile := range tiles {
		tl, tr, c := tile[0], tile[1], tile[2] * (tile[1] - tile[0] + 1)
		_ = c
		cover += (tr - tl + 1) * tile[2]
		for tiles[left][1]+carpetLen-1 < tr {
			cover -= (tiles[left][1] - tiles[left][0] + 1) * tiles[left][2]
			left++
		}
		uncover := max((tr-carpetLen+1-tiles[left][0])*tiles[left][2], 0)
		if cover-uncover > ans {
			ans = cover - uncover
		}
	}
	return ans
}

func maximumCoins(coins [][]int, k int) int64 {
	slices.SortFunc(coins, func(a, b []int) int { return a[0] - b[0] })
	ans := maximumWhiteTiles(coins, k)

	// reverse and negate for right-to-left pass
	slices.Reverse(coins)
	for _, t := range coins {
		t[0], t[1] = -t[1], -t[0]
	}
	ans2 := maximumWhiteTiles(coins, k)
	if ans2 > ans {
		ans = ans2
	}
	return int64(ans)
}

func main() {
	fmt.Println(maximumCoins([][]int{{8, 10, 1}, {1, 3, 2}, {5, 6, 4}}, 4)) // 10
	fmt.Println(maximumCoins([][]int{{1, 4, 2}, {5, 8, 1}}, 3)) // 6
}
```
