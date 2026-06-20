# 3219 — Minimum Cost For Cutting Cake Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumCostForCuttingCakeIi(h int, w int, horizontalCut []int, verticalCut []int) int64
```

> **💡 Hint:** Greedy — always cut with the highest cost first.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3219: Minimum Cost for Cutting Cake II
// https://leetcode.com/problems/minimum-cost-for-cutting-cake-ii/
// Difficulty: Hard
//
// You have a cake of size h x w. You must make all horizontal and vertical
// cuts to divide it into unit squares. horizontalCut[i] is the cost of making
// the i-th horizontal cut. verticalCut[j] is the cost of the j-th vertical
// cut. Each cut's cost is multiplied by the number of pieces it goes through.
// Find the minimum total cost.
//
// Approach: Greedy — always cut with the highest cost first.
//   hPieces = number of horizontal strips (initially 1)
//   vPieces = number of vertical strips (initially 1)
//   When making a horizontal cut, cost *= vPieces, then hPieces++
//   When making a vertical cut,   cost *= hPieces, then vPieces++

import (
	"fmt"
	"sort"
)

func minimumCostForCuttingCakeIi(h int, w int, horizontalCut []int, verticalCut []int) int64 {
  // Custom sort dengan comparator
	sort.Slice(horizontalCut, func(i, j int) bool { return horizontalCut[i] > horizontalCut[j] })
  // Custom sort dengan comparator
	sort.Slice(verticalCut, func(i, j int) bool { return verticalCut[i] > verticalCut[j] })

	var total int64 = 0
	hPieces, vPieces := 1, 1
	i, j := 0, 0

	for i < len(horizontalCut) || j < len(verticalCut) {
		if j >= len(verticalCut) || (i < len(horizontalCut) && horizontalCut[i] > verticalCut[j]) {
			total += int64(horizontalCut[i]) * int64(vPieces)
			hPieces++
			i++
		} else {
			total += int64(verticalCut[j]) * int64(hPieces)
			vPieces++
			j++
		}
	}
	return total
}

func main() {
	// h=3, w=2, horizontal=[1,3], vertical=[5]
	fmt.Println(minimumCostForCuttingCakeIi(3, 2, []int{1, 3}, []int{5})) // expect: 3*1 + 3*2 + 5*3 = 3+6+15=24
}
```
