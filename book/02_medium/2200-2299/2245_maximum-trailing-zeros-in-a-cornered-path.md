# 2245 — Maximum Trailing Zeros In A Cornered Path

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxTrailingZeros(grid [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Prefix Sum

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(m * n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2245: Maximum Trailing Zeros in a Cornered Path
// https://leetcode.com/problems/maximum-trailing-zeros-in-a-cornered-path/
// Difficulty: Medium
// Time: O(m * n) | Space: O(m * n)

import "fmt"

func maxTrailingZeros(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// Precompute prefix counts of factors 2 and 5
	type pair struct{ two, five int }

	// Right prefix
  // Membuat matriks/slice 2D untuk DP
	right := make([][]pair, m)
	for i := 0; i < m; i++ {
		right[i] = make([]pair, n+1)
		for j := 0; j < n; j++ {
			two, five := countFactors(grid[i][j])
			right[i][j+1] = pair{right[i][j].two + two, right[i][j].five + five}
		}
	}

	// Down prefix
  // Membuat matriks/slice 2D untuk DP
	down := make([][]pair, m+1)
	for i := 0; i <= m; i++ {
		down[i] = make([]pair, n)
	}
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			two, five := countFactors(grid[i][j])
			down[i+1][j] = pair{down[i][j].two + two, down[i][j].five + five}
		}
	}

	maxZeros := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// From top to (i,j) then right
			t1 := down[i+1][j].two + right[i][n].two - right[i][j+1].two
			f1 := down[i+1][j].five + right[i][n].five - right[i][j+1].five
			if t1 < f1 {
				if t1 > maxZeros {
					maxZeros = t1
				}
			} else {
				if f1 > maxZeros {
					maxZeros = f1
				}
			}

			// From top to (i,j) then left
			t2 := down[i+1][j].two + right[i][j].two
			f2 := down[i+1][j].five + right[i][j].five
			if t2 < f2 {
				if t2 > maxZeros {
					maxZeros = t2
				}
			} else {
				if f2 > maxZeros {
					maxZeros = f2
				}
			}

			// From bottom to (i,j) then right
			t3 := (down[m][j].two - down[i][j].two) + right[i][n].two - right[i][j+1].two
			f3 := (down[m][j].five - down[i][j].five) + right[i][n].five - right[i][j+1].five
			if t3 < f3 {
				if t3 > maxZeros {
					maxZeros = t3
				}
			} else {
				if f3 > maxZeros {
					maxZeros = f3
				}
			}

			// From bottom to (i,j) then left
			t4 := (down[m][j].two - down[i][j].two) + right[i][j].two
			f4 := (down[m][j].five - down[i][j].five) + right[i][j].five
			if t4 < f4 {
				if t4 > maxZeros {
					maxZeros = t4
				}
			} else {
				if f4 > maxZeros {
					maxZeros = f4
				}
			}
		}
	}
	return maxZeros
}

func countFactors(x int) (int, int) {
	two, five := 0, 0
	for x%2 == 0 {
		two++
		x /= 2
	}
	for x%5 == 0 {
		five++
		x /= 5
	}
	return two, five
}

func main() {
	// Test case 1
	fmt.Println(maxTrailingZeros([][]int{{23, 17, 15, 3, 20}, {8, 1, 20, 27, 11}, {9, 4, 6, 2, 21}, {40, 9, 1, 10, 6}, {22, 7, 4, 5, 3}}))
	// Expected: 3

	// Test case 2
	fmt.Println(maxTrailingZeros([][]int{{4, 3, 2}, {7, 6, 1}, {8, 8, 8}}))
	// Expected: 0
}
```
