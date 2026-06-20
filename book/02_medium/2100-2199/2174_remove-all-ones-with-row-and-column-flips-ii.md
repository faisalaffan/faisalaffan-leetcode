# 2174 — Remove All Ones With Row And Column Flips Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func removeOnes(grid [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS, Bitmask

**Kompleksitas Waktu:** O(2^(m*n) * m * n)  
**Kompleksitas Ruang:** O(2^(m*n))

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2174: Remove All Ones With Row and Column Flips II
// https://leetcode.com/problems/remove-all-ones-with-row-and-column-flips-ii/
// Difficulty: Medium [Paid]
// Time: O(2^(m*n) * m * n) | Space: O(2^(m*n))

import "fmt"

func removeOnes(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	total := m * n

	// Compress grid to bitmask
	start := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				start |= 1 << (i*n + j)
			}
		}
	}

	if start == 0 {
		return 0
	}

  // Alokasi slice integer
	dist := make([]int, 1<<total)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dist {
		dist[i] = -1
	}
	dist[start] = 0
	queue := []int{start}

	for len(queue) > 0 {
		mask := queue[0]
		queue = queue[1:]

		// Only try cells that were 1 in the ORIGINAL grid
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] == 0 {
					continue
				}
				next := mask
				// Clear entire row i
				for c := 0; c < n; c++ {
					next &^= 1 << (i*n + c)
				}
				// Clear entire column j
				for r := 0; r < m; r++ {
					next &^= 1 << (r*n + j)
				}

				if dist[next] == -1 {
					dist[next] = dist[mask] + 1
					if next == 0 {
						return dist[next]
					}
					queue = append(queue, next)
				}
			}
		}
	}

	return -1
}

func main() {
	// Test case 1 (Example 2 from problem)
	fmt.Println("Test 1:", removeOnes([][]int{{0, 1, 0}, {1, 0, 1}, {0, 1, 0}}))
	// Expected: 2

	// Test case 2 (Example 1 from problem)
	fmt.Println("Test 2:", removeOnes([][]int{{1, 1, 1}, {1, 1, 1}, {0, 1, 0}}))
	// Expected: 2

	// Test case 3 (all zeros)
	fmt.Println("Test 3:", removeOnes([][]int{{0, 0}, {0, 0}}))
	// Expected: 0
}
```
