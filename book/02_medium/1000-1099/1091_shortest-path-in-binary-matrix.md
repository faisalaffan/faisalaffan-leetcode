# 1091 — Shortest Path In Binary Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func shortestPathBinaryMatrix(grid [][]int) int
```

> **💡 Hint:** BFS from (0,0) to (n-1,n-1) with 8-directional moves

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n^2)

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1091: Shortest Path in Binary Matrix
// https://leetcode.com/problems/shortest-path-in-binary-matrix/
// Difficulty: Medium
//
// Approach: BFS from (0,0) to (n-1,n-1) with 8-directional moves
// Time: O(n^2)
// Space: O(n^2)

import "fmt"

func main() {
	fmt.Println(shortestPathBinaryMatrix([][]int{{0, 1}, {1, 0}})) // 2
	fmt.Println(shortestPathBinaryMatrix([][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}})) // 4
}

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}
	if n == 1 {
		return 1
	}

	dirs := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
  // Alokasi slice integer
	queue := make([][2]int, 0, n*n)
	queue = append(queue, [2]int{0, 0})
	grid[0][0] = 1
	distance := 1

	for len(queue) > 0 {
		size := len(queue)
		distance++
		for k := 0; k < size; k++ {
			cur := queue[k]
			for _, d := range dirs {
				ni, nj := cur[0]+d[0], cur[1]+d[1]
				if ni >= 0 && ni < n && nj >= 0 && nj < n && grid[ni][nj] == 0 {
					if ni == n-1 && nj == n-1 {
						return distance
					}
					grid[ni][nj] = 1
					queue = append(queue, [2]int{ni, nj})
				}
			}
		}
		queue = queue[size:]
	}

	return -1
}
```
