# 2017 — Grid Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func gridGame(grid [][]int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum, Bitmask

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2017: Grid Game
// https://leetcode.com/problems/grid-game/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func gridGame(grid [][]int) int64 {
	n := len(grid[0])

  // Alokasi slice integer
	topPrefix := make([]int64, n+1)
  // Alokasi slice integer
	bottomPrefix := make([]int64, n+1)
	for i := 0; i < n; i++ {
		topPrefix[i+1] = topPrefix[i] + int64(grid[0][i])
		bottomPrefix[i+1] = bottomPrefix[i] + int64(grid[1][i])
	}

	result := int64(1<<63 - 1)

	for k := 0; k < n; k++ {
		// Robot 1 goes down at column k
		topRemaining := topPrefix[n] - topPrefix[k+1]
		bottomRemaining := bottomPrefix[k]
		score := topRemaining
		if bottomRemaining > score {
			score = bottomRemaining
		}
		if score < result {
			result = score
		}
	}

	return result
}

func main() {
	// Test case 1
	grid1 := [][]int{{2, 5, 4}, {1, 5, 1}}
	fmt.Println("Test 1:", gridGame(grid1))
	// Expected: 4

	// Test case 2
	grid2 := [][]int{{3, 3, 1}, {8, 5, 2}}
	fmt.Println("Test 2:", gridGame(grid2))
	// Expected: 4

	// Test case 3
	grid3 := [][]int{{1, 3, 1, 15}, {1, 3, 3, 1}}
	fmt.Println("Test 3:", gridGame(grid3))
	// Expected: 7
}
```
