# 1582 — Special Positions In A Binary Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func numSpecial(mat [][]int) int

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m * n), Space: O(m + n)  
**Kompleksitas Ruang:** O(m + n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1582: Special Positions in a Binary Matrix
// https://leetcode.com/problems/special-positions-in-a-binary-matrix/
// Difficulty: Easy
//
// LeetCode submission: func numSpecial(mat [][]int) int

import "fmt"

func main() {
	mat1 := [][]int{
		{1, 0, 0},
		{0, 0, 1},
		{1, 0, 0},
	}
	fmt.Println(SpecialPositionsInABinaryMatrix(mat1)) // 1

	mat2 := [][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	fmt.Println(SpecialPositionsInABinaryMatrix(mat2)) // 3
}

// Time: O(m * n), Space: O(m + n)
func SpecialPositionsInABinaryMatrix(mat [][]int) int {
	m, n := len(mat), len(mat[0])
  // Alokasi slice integer
	rows := make([]int, m)
  // Alokasi slice integer
	cols := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				rows[i]++
				cols[j]++
			}
		}
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 && rows[i] == 1 && cols[j] == 1 {
				count++
			}
		}
	}
	return count
}
```
