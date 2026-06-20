# 0361 — Bomb Enemy

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxKilledEnemies(grid [][]byte) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m*n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #361: Bomb Enemy
// https://leetcode.com/problems/bomb-enemy/
// Difficulty: Medium [Paid]
// Time: O(m*n) | Space: O(n)

import "fmt"

func maxKilledEnemies(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	rowHits := 0
  // Alokasi slice integer
	colHits := make([]int, n)
	maxKill := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// Reset row hit count at start of row or after wall
			if j == 0 || grid[i][j-1] == 'W' {
				rowHits = 0
				for k := j; k < n && grid[i][k] != 'W'; k++ {
					if grid[i][k] == 'E' {
						rowHits++
					}
				}
			}

			// Reset col hit count at start of col or after wall
			if i == 0 || grid[i-1][j] == 'W' {
				colHits[j] = 0
				for k := i; k < m && grid[k][j] != 'W'; k++ {
					if grid[k][j] == 'E' {
						colHits[j]++
					}
				}
			}

			// Place bomb at empty cell
			if grid[i][j] == '0' {
				if rowHits+colHits[j] > maxKill {
					maxKill = rowHits + colHits[j]
				}
			}
		}
	}
	return maxKill
}

func main() {
	// Test case 1
	grid1 := [][]byte{
		{'0', 'E', '0', '0'},
		{'E', '0', 'W', 'E'},
		{'0', 'E', '0', '0'},
	}
	fmt.Println("Test 1:", maxKilledEnemies(grid1))
	// Expected: 3

	// Test case 2
	grid2 := [][]byte{{'W', 'W', 'W'}, {'0', '0', '0'}, {'E', 'E', 'E'}}
	fmt.Println("Test 2:", maxKilledEnemies(grid2))
	// Expected: 1

	// Test case 3: Empty
	grid3 := [][]byte{{'0'}}
	fmt.Println("Test 3:", maxKilledEnemies(grid3))
	// Expected: 0
}
```
