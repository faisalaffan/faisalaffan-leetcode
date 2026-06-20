# 0490 — The Maze

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func TheMaze(maze [][]int, start []int, destination []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(m * n)

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #490: The Maze
// https://leetcode.com/problems/the-maze/
// Difficulty: Medium [Paid]
// Time: O(m * n)
// Space: O(m * n)

import "fmt"

func main() {
	maze := [][]int{
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1},
		{0, 0, 0, 0, 0},
	}
	fmt.Println(TheMaze(maze, []int{0, 4}, []int{4, 4}))
	fmt.Println(TheMaze(maze, []int{0, 4}, []int{3, 2}))
}

func TheMaze(maze [][]int, start []int, destination []int) bool {
	m, n := len(maze), len(maze[0])
  // Membuat matriks/slice 2D untuk DP
	visited := make([][]bool, m)
  // Range loop: iterasi dengan indeks + nilai
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	queue := [][]int{start}
	visited[start[0]][start[1]] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur[0] == destination[0] && cur[1] == destination[1] {
			return true
		}

		for _, d := range dirs {
			r, c := cur[0], cur[1]
			// Roll until hitting a wall
			for r+d[0] >= 0 && r+d[0] < m && c+d[1] >= 0 && c+d[1] < n && maze[r+d[0]][c+d[1]] == 0 {
				r += d[0]
				c += d[1]
			}
			if !visited[r][c] {
				visited[r][c] = true
				queue = append(queue, []int{r, c})
			}
		}
	}

	return false
}
```
