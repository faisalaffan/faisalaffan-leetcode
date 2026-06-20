# 0305 — Number Of Islands Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func NewUnionFind(n int) *UnionFind
```

> **💡 Hint:** Union-Find (Disjoint Set Union) with path compression and union by size.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Union-Find (DSU)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Union-Find (DSU)** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #305: Number of Islands II
// https://leetcode.com/problems/number-of-islands-ii/
// Difficulty: Hard [Paid]
//
// Approach: Union-Find (Disjoint Set Union) with path compression and union by size.
//   - Start with an empty grid of size m x n.
//   - For each land addition (r, c), mark it as land and check 4-directional neighbors.
//   - If a neighbor is also land, union the two cells.
//   - Track the current number of islands after each operation.
//   - Map 2D coordinates to 1D index: id = r * n + c.

import "fmt"

func main() {
	// Example: m=3, n=3, positions=[[0,0],[0,1],[1,2],[2,1]]
	positions := [][]int{{0, 0}, {0, 1}, {1, 2}, {2, 1}}
	result := numIslands2(3, 3, positions)
	fmt.Println("Number of islands after each addition:", result) // [1, 1, 2, 3]

	// Example: all adjacent
	positions2 := [][]int{{0, 0}, {0, 1}, {1, 1}, {1, 0}}
	result2 := numIslands2(2, 2, positions2)
	fmt.Println("All adjacent:", result2) // [1, 1, 1, 1]

	// Example: isolated islands
	positions3 := [][]int{{0, 0}, {0, 2}, {2, 0}, {2, 2}}
	result3 := numIslands2(3, 3, positions3)
	fmt.Println("Isolated:", result3) // [1, 2, 3, 4]

	// Example: duplicate positions.
	positions4 := [][]int{{0, 0}, {0, 0}}
	result4 := numIslands2(1, 1, positions4)
	fmt.Println("Duplicate:", result4) // [1, 1]
}

// UnionFind implements DSU with path compression and union by size.
type UnionFind struct {
	parent []int
	size   []int
}

// NewUnionFind creates a new UnionFind for n elements.
func NewUnionFind(n int) *UnionFind {
  // Alokasi slice integer
	parent := make([]int, n)
  // Alokasi slice integer
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFind{parent: parent, size: size}
}

// Find finds the root of x (with path compression).
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

// Union unions two elements. Returns true if they were actually merged.
func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return false
	}

	// Union by size.
	if uf.size[rootX] < uf.size[rootY] {
		rootX, rootY = rootY, rootX
	}

	uf.parent[rootY] = rootX
	uf.size[rootX] += uf.size[rootY]
	return true
}

// numIslands2 returns the number of islands after each land addition.
func numIslands2(m int, n int, positions [][]int) []int {
	if m <= 0 || n <= 0 || len(positions) == 0 {
		return []int{}
	}

  // Membuat matriks/slice 2D untuk DP
	grid := make([][]bool, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]bool, n)
	}

	uf := NewUnionFind(m * n)
  // Alokasi slice integer
	result := make([]int, 0, len(positions))
	islands := 0

	// 4-directional neighbors.
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for _, pos := range positions {
		r, c := pos[0], pos[1]

		if grid[r][c] {
			// Duplicate position.
			result = append(result, islands)
			continue
		}

		grid[r][c] = true
		islands++

		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] {
				id1 := r*n + c
				id2 := nr*n + nc
				if uf.Union(id1, id2) {
					islands--
				}
			}
		}

		result = append(result, islands)
	}

	return result
}

// Stub compatibility.
func NumberOfIslandsIi() any {
	positions := [][]int{{0, 0}, {0, 1}, {1, 2}, {2, 1}}
	return numIslands2(3, 3, positions)
}
```
