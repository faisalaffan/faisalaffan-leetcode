# 2959 — Number Of Possible Sets Of Closing Branches

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func numberOfSets(n int, maxDistance int, roads [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Floyd-Warshall

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Floyd-Warshall** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2959: Number of Possible Sets of Closing Branches
// https://leetcode.com/problems/number-of-possible-sets-of-closing-branches/
// Difficulty: Hard
//
// Given n nodes (0..n-1), some branches may be closed (removed).
// For each subset of remaining branches, run Floyd-Warshall to compute
// all-pairs shortest paths. A subset is valid if all pairwise distances
// among remaining nodes are <= maxDistance.
// Count valid subsets (including empty set? Typically yes, with n=0 it's trivially valid).

import (
	"fmt"
)

func numberOfSets(n int, maxDistance int, roads [][]int) int {
	const inf = 1 << 29

	// Build adjacency matrix
  // Membuat matriks/slice 2D untuk DP
	g := make([][]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = inf
		}
		g[i][i] = 0
	}
	for _, r := range roads {
		u, v, w := r[0], r[1], r[2]
		if w < g[u][v] {
			g[u][v] = w
			g[v][u] = w
		}
	}

	ans := 0

	// Try all subsets of open branches
	for mask := 0; mask < (1 << n); mask++ {
		// Copy distances for this subset
  // Membuat matriks/slice 2D untuk DP
		dist := make([][]int, n)
  // Range loop: iterasi dengan indeks + nilai
		for i := range dist {
			dist[i] = make([]int, n)
			copy(dist[i], g[i])
		}

		// Floyd-Warshall only for open branches
		for k := 0; k < n; k++ {
			if mask>>k&1 == 0 {
				continue
			}
			for i := 0; i < n; i++ {
				if mask>>i&1 == 0 || dist[i][k] == inf {
					continue
				}
				for j := 0; j < n; j++ {
					if mask>>j&1 == 0 {
						continue
					}
					if nd := dist[i][k] + dist[k][j]; nd < dist[i][j] {
						dist[i][j] = nd
					}
				}
			}
		}

		// Validate all pairwise distances for open branches
		ok := true
		for i := 0; i < n && ok; i++ {
			if mask>>i&1 == 0 {
				continue
			}
			for j := i + 1; j < n; j++ {
				if mask>>j&1 == 0 {
					continue
				}
				if dist[i][j] > maxDistance {
					ok = false
					break
				}
			}
		}
		if ok {
			ans++
		}
	}
	return ans
}

func main() {
	// Example
	fmt.Println(numberOfSets(3, 5, [][]int{{0, 1, 2}, {1, 2, 10}, {0, 2, 10}}))
	fmt.Println(numberOfSets(3, 5, [][]int{{0, 1, 20}, {0, 2, 5}, {1, 2, 2}}))

	// Edge cases
	fmt.Println(numberOfSets(1, 0, [][]int{}))
	fmt.Println(numberOfSets(2, 1, [][]int{{0, 1, 2}}))
}
```
