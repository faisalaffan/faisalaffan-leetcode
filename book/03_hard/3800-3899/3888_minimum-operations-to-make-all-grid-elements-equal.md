# 3888 — Minimum Operations To Make All Grid Elements Equal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minOperations(grid [][]int, k int) int64
```

> **💡 Hint:** Flatten grid, sort. All elements can be made equal to

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3888: Minimum Operations to Make All Grid Elements Equal
// https://leetcode.com/problems/minimum-operations-to-make-all-grid-elements-equal/
// Difficulty: Hard [Paid]
//
// In one operation, you can add or subtract k from any element.
// Find minimum operations to make all grid elements equal.
//
// Approach: Flatten grid, sort. All elements can be made equal to
// any element (the median minimizes total operations). Since all
// elements must be congruent modulo k, check that first. Then
// compute min operations as sum of |val - target| / k.

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(minOperations([][]int{{1, 3}, {5, 7}}, 2))
	// Example 2
	fmt.Println(minOperations([][]int{{2, 4}, {6, 8}}, 2))
	// Edge: single element
	fmt.Println(minOperations([][]int{{5}}, 3))
	// Edge: impossible
	fmt.Println(minOperations([][]int{{1, 2}, {3, 4}}, 2))
}

func minOperations(grid [][]int, k int) int64 {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
  // Alokasi slice integer
	flat := make([]int, 0, m*n)
	rem := grid[0][0] % k
	for _, row := range grid {
		for _, val := range row {
			if val%k != rem {
				return -1
			}
			flat = append(flat, val)
		}
	}

  // Urutkan secara ascending — O(n log n)
	sort.Ints(flat)
	target := flat[len(flat)/2]

	var ops int64
	for _, val := range flat {
		diff := val - target
		if diff < 0 {
			diff = -diff
		}
		ops += int64(diff / k)
	}
	return ops
}
```
