# 1765 — Map Of Highest Peak

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func highestPeak(isWater [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** O(m * n), Space: O(m * n)  
**Kompleksitas Ruang:** O(m * n)

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1765: Map of Highest Peak
// https://leetcode.com/problems/map-of-highest-peak/
// Difficulty: Medium
// Time: O(m * n), Space: O(m * n)

import "fmt"

func highestPeak(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
  // Membuat matriks/slice 2D untuk DP
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = -1
		}
	}

  // Alokasi slice integer
	queue := make([][2]int, 0, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				result[i][j] = 0
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for len(queue) > 0 {
		size := len(queue)
		for k := 0; k < size; k++ {
			r, c := queue[k][0], queue[k][1]
			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < m && nc >= 0 && nc < n && result[nr][nc] == -1 {
					result[nr][nc] = result[r][c] + 1
					queue = append(queue, [2]int{nr, nc})
				}
			}
		}
		queue = queue[size:]
	}
	return result
}

func main() {
	fmt.Println(highestPeak([][]int{{0, 1}, {0, 0}})) // Expected: [[1,0],[2,1]]
	fmt.Println(highestPeak([][]int{{0, 0, 1}, {1, 0, 0}, {0, 0, 0}})) // Expected: [[1,1,0],[0,1,1],[1,2,2]]
}
```
