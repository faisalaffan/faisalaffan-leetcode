# 0265 — Paint House Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minCostII(costs [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #265: Paint House II
// https://leetcode.com/problems/paint-house-ii/
// Difficulty: Hard [Paid]

import (
	"fmt"
	"math"
)

func minCostII(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}

	n := len(costs)
	k := len(costs[0])

	// Find min and second min for the first house
	prevMin1, prevMin2 := math.MaxInt32, math.MaxInt32
	prevMinIdx := -1

	for j := 0; j < k; j++ {
		cost := costs[0][j]
		if cost < prevMin1 {
			prevMin2 = prevMin1
			prevMin1 = cost
			prevMinIdx = j
		} else if cost < prevMin2 {
			prevMin2 = cost
		}
	}

	for i := 1; i < n; i++ {
		curMin1, curMin2 := math.MaxInt32, math.MaxInt32
		curMinIdx := -1

		for j := 0; j < k; j++ {
			var cost int
			if j == prevMinIdx {
				cost = costs[i][j] + prevMin2
			} else {
				cost = costs[i][j] + prevMin1
			}

			if cost < curMin1 {
				curMin2 = curMin1
				curMin1 = cost
				curMinIdx = j
			} else if cost < curMin2 {
				curMin2 = cost
			}
		}

		prevMin1, prevMin2 = curMin1, curMin2
		prevMinIdx = curMinIdx
	}

	return prevMin1
}

func main() {
	fmt.Println(minCostII([][]int{{1, 5, 3}, {2, 9, 4}}))
}
```
