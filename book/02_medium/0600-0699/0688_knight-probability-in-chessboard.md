# 0688 — Knight Probability In Chessboard

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func knightProbability(n int, k int, row int, column int) float64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(K * N^2)  
**Kompleksitas Ruang:** O(N^2)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #688: Knight Probability in Chessboard
// https://leetcode.com/problems/knight-probability-in-chessboard/
// Difficulty: Medium
// Time: O(K * N^2)
// Space: O(N^2)

import "fmt"

func main() {
	fmt.Println(knightProbability(3, 2, 0, 0))
	fmt.Println(knightProbability(1, 0, 0, 0))
}

func knightProbability(n int, k int, row int, column int) float64 {
	dirs := [][2]int{{-2, -1}, {-2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {2, -1}, {2, 1}}
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]float64, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]float64, n)
	}
	dp[row][column] = 1.0

	for step := 0; step < k; step++ {
  // Membuat matriks/slice 2D untuk DP
		next := make([][]float64, n)
  // Range loop: iterasi dengan indeks + nilai
		for i := range next {
			next[i] = make([]float64, n)
		}
		for r := 0; r < n; r++ {
			for c := 0; c < n; c++ {
				if dp[r][c] == 0 {
					continue
				}
				for _, d := range dirs {
					nr, nc := r+d[0], c+d[1]
					if nr >= 0 && nr < n && nc >= 0 && nc < n {
						next[nr][nc] += dp[r][c] / 8.0
					}
				}
			}
		}
		dp = next
	}

	result := 0.0
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			result += dp[r][c]
		}
	}
	return result
}
```
