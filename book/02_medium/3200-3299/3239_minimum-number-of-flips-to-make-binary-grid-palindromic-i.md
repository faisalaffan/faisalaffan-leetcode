# 3239 — Minimum Number Of Flips To Make Binary Grid Palindromic I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func minFlips(grid [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3239: Minimum Number of Flips to Make Binary Grid Palindromic I
// https://leetcode.com/problems/minimum-number-of-flips-to-make-binary-grid-palindromic-i/
// Difficulty: Medium
// Time: O(m * n) | Space: O(1)

import "fmt"

func minFlips(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	rowFlips := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n/2; j++ {
			if grid[i][j] != grid[i][n-1-j] {
				rowFlips++
			}
		}
	}

	colFlips := 0
	for j := 0; j < n; j++ {
		for i := 0; i < m/2; i++ {
			if grid[i][j] != grid[m-1-i][j] {
				colFlips++
			}
		}
	}

	return min(rowFlips, colFlips)
}

func main() {
	fmt.Println(minFlips([][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}})) // Expected: 2
	fmt.Println(minFlips([][]int{{0, 1}, {0, 1}, {0, 0}}))          // Expected: 1
}
```
