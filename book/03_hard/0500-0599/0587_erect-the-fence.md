# 0587 — Erect The Fence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func outerTrees(points [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #587: Erect the Fence
// https://leetcode.com/problems/erect-the-fence/
// Difficulty: Hard

import (
	"fmt"
	"sort"
)

func main() {
	points := [][]int{{1, 1}, {2, 2}, {2, 0}, {2, 4}, {3, 3}, {4, 2}}
	result := outerTrees(points)
	fmt.Println(result)
	// Expected: [[1,1],[2,0],[3,3],[2,4],[4,2]] (order may vary)
}

func outerTrees(points [][]int) [][]int {
	n := len(points)
	if n <= 1 {
		return points
	}

	// Sort by x, then by y
  // Custom sort dengan comparator
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] != points[j][0] {
			return points[i][0] < points[j][0]
		}
		return points[i][1] < points[j][1]
	})

	// Cross product: (b - a) x (c - a)
	cross := func(a, b, c []int) int {
		return (b[0]-a[0])*(c[1]-a[1]) - (b[1]-a[1])*(c[0]-a[0])
	}

	lower := [][]int{}
	for _, p := range points {
		for len(lower) >= 2 && cross(lower[len(lower)-2], lower[len(lower)-1], p) < 0 {
			lower = lower[:len(lower)-1]
		}
		lower = append(lower, p)
	}

	upper := [][]int{}
	for i := n - 1; i >= 0; i-- {
		p := points[i]
		for len(upper) >= 2 && cross(upper[len(upper)-2], upper[len(upper)-1], p) < 0 {
			upper = upper[:len(upper)-1]
		}
		upper = append(upper, p)
	}

	// Remove last point of each half (it's repeated as first of the other half)
	upper = upper[:len(upper)-1]
	lower = lower[:len(lower)-1]

	// Combine and deduplicate
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[[2]int]bool)
	hull := [][]int{}
	for _, p := range append(lower, upper...) {
		key := [2]int{p[0], p[1]}
		if !seen[key] {
			seen[key] = true
			hull = append(hull, p)
		}
	}
	return hull
}
```
