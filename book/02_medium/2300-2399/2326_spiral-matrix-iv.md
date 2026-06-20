# 2326 — Spiral Matrix Iv

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func spiralMatrix(m int, n int, head *ListNode) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2326: Spiral Matrix IV
// https://leetcode.com/problems/spiral-matrix-iv/
// Difficulty: Medium
// Time: O(m * n) | Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
  // Membuat matriks/slice 2D untuk DP
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = -1
		}
	}

	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dir := 0
	r, c := 0, 0

	for head != nil {
		result[r][c] = head.Val
		head = head.Next

		nr, nc := r+dirs[dir][0], c+dirs[dir][1]
		if nr < 0 || nr >= m || nc < 0 || nc >= n || result[nr][nc] != -1 {
			dir = (dir + 1) % 4
			nr, nc = r+dirs[dir][0], c+dirs[dir][1]
		}
		r, c = nr, nc
	}
	return result
}

func main() {
	// Test case 1: head = [3,0,2,6,8,1,7,9,4,2,5,5,0], m = 3, n = 5
	head1 := &ListNode{3, &ListNode{0, &ListNode{2, &ListNode{6, &ListNode{8, &ListNode{1, &ListNode{7, &ListNode{9, &ListNode{4, &ListNode{2, &ListNode{5, &ListNode{5, &ListNode{0, nil}}}}}}}}}}}}}
	fmt.Println(spiralMatrix(3, 5, head1))

	// Test case 2: head = [0,1,2], m = 1, n = 4
	head2 := &ListNode{0, &ListNode{1, &ListNode{2, nil}}}
	fmt.Println(spiralMatrix(1, 4, head2))
}
```
