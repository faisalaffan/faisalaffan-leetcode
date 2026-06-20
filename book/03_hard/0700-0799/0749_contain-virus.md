# 0749 — Contain Virus

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func containVirus(isInfected [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer, Stack

**Kompleksitas Waktu:** O(R*C * rounds) in worst case, Space: O(R*C)  
**Kompleksitas Ruang:** O(R*C)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #749: Contain Virus
// https://leetcode.com/problems/contain-virus/
// Difficulty: Hard
//
// A virus spreads in a 2D grid. Each round:
// 1. Identify all connected regions of infected cells (1).
// 2. For each region, count:
//    - threat level = number of distinct uninfected (0) cells it will infect next.
//    - walls needed = number of edges between infected and uninfected cells.
// 3. Build walls around the region with the highest threat level (add walls to total).
//    Mark that region as "contained" (e.g. 2) so it does not spread further.
// 4. Remaining infected cells (1) spread to all adjacent uninfected (0) cells.
// 5. Repeat until no more infection can spread.
// Return: total number of walls built.

// Directions: up, down, left, right.
var dirs = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// region holds info about a connected virus region.
type region struct {
	cells    [][2]int // coordinates of infected cells in this region
	threat   int      // number of distinct uninfected cells adjacent to this region
	walls    int      // number of edges between infected and uninfected cells
}

// containVirus simulates virus containment and returns total walls built.
// Time: O(R*C * rounds) in worst case, Space: O(R*C)
func containVirus(isInfected [][]int) int {
	if len(isInfected) == 0 || len(isInfected[0]) == 0 {
		return 0
	}
	rows, cols := len(isInfected), len(isInfected[0])
	totalWalls := 0

	for {
		// Step 1: Find all virus regions via DFS.
  // Membuat matriks/slice 2D untuk DP
		visited := make([][]bool, rows)
  // Range loop: iterasi dengan indeks + nilai
		for i := range visited {
			visited[i] = make([]bool, cols)
		}
		var regions []region

		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				if isInfected[r][c] == 1 && !visited[r][c] {
					reg := region{}
					// BFS/DFS to find the connected component.
					stack := [][2]int{{r, c}}
					visited[r][c] = true

					// Use a set for threat cells to count distinct ones.
  // Membuat map (HashMap) — pencarian O(1)
					threatSet := make(map[[2]int]bool)

					for len(stack) > 0 {
						cell := stack[len(stack)-1]
						stack = stack[:len(stack)-1]
						reg.cells = append(reg.cells, cell)

						for _, d := range dirs {
							nr, nc := cell[0]+d[0], cell[1]+d[1]
							if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
								continue
							}
							if isInfected[nr][nc] == 0 {
								// Adjacent to uninfected cell.
								reg.walls++
								threatSet[[2]int{nr, nc}] = true
							} else if isInfected[nr][nc] == 1 && !visited[nr][nc] {
								visited[nr][nc] = true
								stack = append(stack, [2]int{nr, nc})
							}
						}
					}
					reg.threat = len(threatSet)
					regions = append(regions, reg)
				}
			}
		}

		if len(regions) == 0 {
			break // no more virus
		}

		// Step 2: Find the region with the highest threat level.
		// If tie, any will do (LeetCode says "most threatening").
		worstIdx := 0
		for i := 1; i < len(regions); i++ {
			if regions[i].threat > regions[worstIdx].threat {
				worstIdx = i
			}
		}

		// Step 3: Contain the worst region.
		totalWalls += regions[worstIdx].walls
		for _, cell := range regions[worstIdx].cells {
			isInfected[cell[0]][cell[1]] = 2 // contained (inactive)
		}

		// Step 4: Spread remaining virus.
		// Collect cells that will become infected (adjacent to any remaining 1).
  // Membuat map (HashMap) — pencarian O(1)
		toInfect := make(map[[2]int]bool)
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				if isInfected[r][c] != 1 {
					continue
				}
				for _, d := range dirs {
					nr, nc := r+d[0], c+d[1]
					if nr >= 0 && nr < rows && nc >= 0 && nc < cols && isInfected[nr][nc] == 0 {
						toInfect[[2]int{nr, nc}] = true
					}
				}
			}
		}
		for cell := range toInfect {
			isInfected[cell[0]][cell[1]] = 1
		}

		// Step 5: If no threat, stop.
		anyThreat := false
		for _, reg := range regions {
			if reg.threat > 0 {
				anyThreat = true
				break
			}
		}
		if !anyThreat {
			break
		}
	}

	return totalWalls
}

// --- Tests -------------------------------------------------------------------

func main() {
	fmt.Println("=== 0749 Contain Virus ===")

	// Test 1: Simple case from LeetCode example.
	grid1 := [][]int{
		{0, 1, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	w1 := containVirus(grid1)
	fmt.Printf("Test 1 - Walls needed = %d (expected 10)\n", w1)

	// Test 2: Single infected cell isolated.
	grid2 := [][]int{
		{1, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	w2 := containVirus(grid2)
	fmt.Printf("Test 2 - Single cell walls = %d (expected 4)\n", w2)

	// Test 3: No infection.
	grid3 := [][]int{
		{0, 0},
		{0, 0},
	}
	w3 := containVirus(grid3)
	fmt.Printf("Test 3 - No virus = %d (expected 0)\n", w3)

	// Test 4: Two regions with different threat levels.
	// Region A (3 cells): (0,0),(0,1),(1,0) threat = 2 (cells (0,2),(2,0))
	// Region B (1 cell):  (0,4) threat = 3 (cells (0,3),(0,5),(1,4))
	// B has higher threat, so walls built around B first.
	grid4 := [][]int{
		{1, 1, 0, 0, 1},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	w4 := containVirus(grid4)
	fmt.Printf("Test 4 - Walls = %d\n", w4)

	// Test 5: Virus that spreads.
	grid5 := [][]int{
		{1, 1, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	w5 := containVirus(grid5)
	fmt.Printf("Test 5 - Spreading virus walls = %d\n", w5)

	// Test 6: Empty grid.
	w6 := containVirus([][]int{})
	fmt.Printf("Test 6 - Empty = %d (expected 0)\n", w6)
}
```
