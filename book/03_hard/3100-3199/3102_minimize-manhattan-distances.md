# 3102 — Minimize Manhattan Distances

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumDistance(points [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3102: Minimize Manhattan Distances
// https://leetcode.com/problems/minimize-manhattan-distances/
// Difficulty: Hard
//
// Manhattan distance = |x1-x2| + |y1-y2| = max(u1-u2, v1-v2) where u=x+y, v=x-y.
// Max Manhattan distance among points = max(max_u - min_u, max_v - min_v).
// To minimize after removing one point, try removing each of the 4 extreme points.

import (
	"fmt"
	"math"
)

func minimumDistance(points [][]int) int {
	n := len(points)
	if n <= 2 {
		return 0
	}

	// Track top-2 max and min for u = x+y and v = x-y
	max1U, max2U := math.MinInt32, math.MinInt32
	min1U, min2U := math.MaxInt32, math.MaxInt32
	max1V, max2V := math.MinInt32, math.MinInt32
	min1V, min2V := math.MaxInt32, math.MaxInt32
	idxMax1U, idxMin1U := -1, -1
	idxMax1V, idxMin1V := -1, -1

	for i, p := range points {
		u := p[0] + p[1]
		v := p[0] - p[1]

		if u > max1U {
			max2U = max1U
			max1U = u
			idxMax1U = i
		} else if u > max2U {
			max2U = u
		}
		if u < min1U {
			min2U = min1U
			min1U = u
			idxMin1U = i
		} else if u < min2U {
			min2U = u
		}

		if v > max1V {
			max2V = max1V
			max1V = v
			idxMax1V = i
		} else if v > max2V {
			max2V = v
		}
		if v < min1V {
			min2V = min1V
			min1V = v
			idxMin1V = i
		} else if v < min2V {
			min2V = v
		}
	}

	// Try removing each extreme point candidate
	candidates := map[int]bool{
		idxMax1U: true,
		idxMin1U: true,
		idxMax1V: true,
		idxMin1V: true,
	}

	result := math.MaxInt32
	for idx := range candidates {
		maxU := max1U
		if idx == idxMax1U {
			maxU = max2U
		}
		minU := min1U
		if idx == idxMin1U {
			minU = min2U
		}
		maxV := max1V
		if idx == idxMax1V {
			maxV = max2V
		}
		minV := min1V
		if idx == idxMin1V {
			minV = min2V
		}
		dist := max(maxU-minU, maxV-minV)
		if dist < result {
			result = dist
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minimumDistance([][]int{{3, 10}, {5, 15}, {1, 5}, {2, 2}, {4, 4}}))
	// Expected: 9

	// Test case 2
	fmt.Println("Test 2:", minimumDistance([][]int{{3, 10}, {5, 15}, {10, 2}, {4, 4}}))
	// Expected: 12

	// Test case 3
	fmt.Println("Test 3:", minimumDistance([][]int{{1, 1}, {1, 1}, {1, 1}}))
	// Expected: 0

	// Test case 4: two points
	fmt.Println("Test 4:", minimumDistance([][]int{{0, 0}, {3, 4}}))
	// Expected: 0 (n <= 2 → 0)

	// Test case 5: LeetCode example
	fmt.Println("Test 5:", minimumDistance([][]int{{1, 2}, {3, 4}, {5, 6}}))
	// Expected: 4
}
```
