# 0909 — Snakes And Ladders

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func SnakesAndLadders(board [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n^2)

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #909: Snakes and Ladders
// https://leetcode.com/problems/snakes-and-ladders/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(SnakesAndLadders([][]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 35, -1, -1, 13, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 15, -1, -1, -1, -1},
	}))
	fmt.Println(SnakesAndLadders([][]int{{-1, -1}, {-1, 3}}))
}

// Time: O(n^2) | Space: O(n^2)
func SnakesAndLadders(board [][]int) int {
	n := len(board)
	target := n * n

	// Convert board position to (row, col)
	posToCoord := func(pos int) (int, int) {
		row := (pos - 1) / n
		col := (pos - 1) % n
		if row%2 == 1 {
			col = n - 1 - col
		}
		return n - 1 - row, col
	}

  // Alokasi slice integer
	dist := make([]int, target+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dist {
		dist[i] = -1
	}
	dist[1] = 0

	queue := []int{1}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr == target {
			return dist[curr]
		}

		for next := curr + 1; next <= curr+6 && next <= target; next++ {
			r, c := posToCoord(next)
			dest := next
			if board[r][c] != -1 {
				dest = board[r][c]
			}
			if dist[dest] == -1 {
				dist[dest] = dist[curr] + 1
				queue = append(queue, dest)
			}
		}
	}

	return -1
}
```
