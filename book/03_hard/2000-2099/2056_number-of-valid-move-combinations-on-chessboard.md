# 2056 — Number Of Valid Move Combinations On Chessboard

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countCombinations(pieces []string, positions [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS, Backtracking

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2056: Number of Valid Move Combinations On Chessboard
// https://leetcode.com/problems/number-of-valid-move-combinations-on-chessboard/
// Difficulty: Hard
//
// For each piece on an 8x8 board, enumerate all possible straight-line moves
// (including staying still). Then use backtracking to try all combinations and
// simulate simultaneous movement step-by-step to detect collisions.

import "fmt"

type piece byte

const (
	rook   piece = 'r'
	bishop piece = 'b'
	queen  piece = 'q'
)

type move struct{ dr, dc, steps int }

var rookDirs = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
var bishopDirs = [][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
var queenDirs = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

func main() {
	// Example: single rook at (1,1)
	fmt.Println(countCombinations([]string{"rook"}, [][]int{{1, 1}}))

	// Example: single queen
	fmt.Println(countCombinations([]string{"queen"}, [][]int{{1, 1}}))

	// Two rooks
	fmt.Println(countCombinations([]string{"rook", "rook"}, [][]int{{1, 1}, {8, 8}}))
}

func countCombinations(pieces []string, positions [][]int) int {
	n := len(pieces)
	pts := make([]piece, n)
  // Alokasi slice integer
	start := make([][2]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range pieces {
		pts[i] = piece(pieces[i][0])
		start[i] = [2]int{positions[i][0] - 1, positions[i][1] - 1}
	}

  // Membuat matriks/slice 2D untuk DP
	allMoves := make([][]move, n)
	for i := 0; i < n; i++ {
		allMoves[i] = genMoves(pts[i], start[i])
	}

	ans := 0
	chosen := make([]move, n)
	var dfs func(int)
	dfs = func(idx int) {
		if idx == n {
			if simulate(start, chosen) {
				ans++
			}
			return
		}
		for _, m := range allMoves[idx] {
			chosen[idx] = m
			dfs(idx + 1)
		}
	}
	dfs(0)
	return ans
}

func genMoves(pt piece, pos [2]int) []move {
	var dirs [][2]int
	switch pt {
	case rook:
		dirs = rookDirs
	case bishop:
		dirs = bishopDirs
	case queen:
		dirs = queenDirs
	}

	var moves []move
	// Stay still
	moves = append(moves, move{0, 0, 0})

	for _, d := range dirs {
		maxSteps := 0
		r, c := pos[0]+d[0], pos[1]+d[1]
		for r >= 0 && r < 8 && c >= 0 && c < 8 {
			maxSteps++
			r += d[0]
			c += d[1]
		}
		for s := 1; s <= maxSteps; s++ {
			moves = append(moves, move{d[0], d[1], s})
		}
	}
	return moves
}

func simulate(start [][2]int, chosen []move) bool {
	n := len(start)
  // Alokasi slice integer
	pos := make([][2]int, n)
	copy(pos, start)

	// Find the maximum number of steps among all chosen moves
	maxSteps := 0
	for _, m := range chosen {
		if m.steps > maxSteps {
			maxSteps = m.steps
		}
	}

	for step := 1; step <= maxSteps; step++ {
		for i := 0; i < n; i++ {
			if step <= chosen[i].steps {
				pos[i][0] += chosen[i].dr
				pos[i][1] += chosen[i].dc
			}
		}
		// Check for collisions at this step
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if pos[i] == pos[j] {
					return false
				}
			}
		}
	}
	return true
}
```
