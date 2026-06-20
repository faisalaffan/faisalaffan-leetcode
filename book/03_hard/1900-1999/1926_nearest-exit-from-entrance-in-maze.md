# 1926 — Nearest Exit From Entrance In Maze

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func nearestExit(maze [][]byte, entrance []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1926: Nearest Exit from Entrance in Maze
// https://leetcode.com/problems/nearest-exit-from-entrance-in-maze/
// Difficulty: Hard - BFS

import "fmt"

func nearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	queue := [][2]int{{entrance[0], entrance[1]}}
  // Membuat matriks/slice 2D untuk DP
	dist := make([][]int, m)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}
	dist[entrance[0]][entrance[1]] = 0

	for len(queue) > 0 {
		r, c := queue[0][0], queue[0][1]
		queue = queue[1:]

		// Check if this is an exit (border, not entrance)
		if (r == 0 || r == m-1 || c == 0 || c == n-1) &&
			!(r == entrance[0] && c == entrance[1]) {
			return dist[r][c]
		}

		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < m && nc >= 0 && nc < n &&
				maze[nr][nc] == '.' && dist[nr][nc] == -1 {
				dist[nr][nc] = dist[r][c] + 1
				queue = append(queue, [2]int{nr, nc})
			}
		}
	}
	return -1
}

func main() {
	// Example 1
	maze1 := [][]byte{
		{'+', '+', '.', '+'},
		{'.', '.', '.', '+'},
		{'+', '+', '+', '.'},
	}
	fmt.Println(nearestExit(maze1, []int{1, 2})) // Expected: 1

	// Example 2
	maze2 := [][]byte{
		{'+', '+', '+'},
		{'.', '.', '.'},
		{'+', '+', '+'},
	}
	fmt.Println(nearestExit(maze2, []int{1, 0})) // Expected: 2

	// Example 3
	maze3 := [][]byte{{'.', '+'}}
	fmt.Println(nearestExit(maze3, []int{0, 0})) // Expected: -1
}
```
