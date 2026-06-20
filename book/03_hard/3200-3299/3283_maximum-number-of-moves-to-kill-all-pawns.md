# 3283 — Maximum Number Of Moves To Kill All Pawns

## Deskripsi

**Soal:** [3283. Maximum Number Of Moves To Kill All Pawns](https://leetcode.com/problems/maximum-number-of-moves-to-kill-all-pawns/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP), BFS (Breadth-First Search / pencarian lebar), Bitmask (representasi himpunan dengan bit)

## Solusi Go

```go
package main

// LeetCode #3283: Maximum Number of Moves to Kill All Pawns
// https://leetcode.com/problems/maximum-number-of-moves-to-kill-all-pawns/
// Difficulty: Hard
//
// A knight starts at (kx, ky) on a 50x50 chessboard. There are N pawns
// at given positions. Alice and Bob take turns moving the knight to
// capture a pawn. Alice goes first. Alice wants to maximize the total
// number of moves, Bob wants to minimize it.
//
// This is a minimax DP over subsets (bitmask DP):
//   - Precompute BFS distances between all pairs of positions
//     (starting position + all pawns).
//   - dp[mask][pos] = optimal total moves from this state
//     (mask of captured pawns, current position index).
//   - Alice maximizes, Bob minimizes based on turn parity.
//
// Position indices: 0 = knight start, 1..N = pawns.

import (
	"fmt"
	"math"
)

func main() {
	// Example 1
	fmt.Println(maxMovesToKillAllPawns(0, 0, [][]int{{1, 2}, {2, 4}}))
	// Example 2
	fmt.Println(maxMovesToKillAllPawns(0, 2, [][]int{{1, 1}, {2, 2}, {3, 3}}))
	// Example 3: single pawn
	fmt.Println(maxMovesToKillAllPawns(1, 1, [][]int{{3, 4}}))
	// Example 4: no pawns
	fmt.Println(maxMovesToKillAllPawns(0, 0, [][]int{}))
	// Example 5
	fmt.Println(maxMovesToKillAllPawns(0, 0, [][]int{{0, 1}, {1, 0}, {2, 2}}))
}

var knightMoves = [][2]int{
	{-2, -1}, {-2, 1}, {-1, -2}, {-1, 2},
	{1, -2}, {1, 2}, {2, -1}, {2, 1},
}

const boardSize = 50

func maxMovesToKillAllPawns(kx, ky int, positions [][]int) int {
	n := len(positions)

	// Total points = starting position + N pawns.
	total := n + 1
  // Membuat slice untuk menyimpan hasil
	pts := make([][2]int, total)
	pts[0] = [2]int{kx, ky}
	for i, p := range positions {
		pts[i+1] = [2]int{p[0], p[1]}
	}

	// Precompute BFS distances between every pair of points.
  // Membuat slice 2D untuk DP/tabel
	dist := make([][]int, total)
  // Iterasi seluruh elemen
	for i := range dist {
		dist[i] = make([]int, total)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}

	for i := 0; i < total; i++ {
		// BFS from pts[i] to all other points.
		d := bfs(pts[i][0], pts[i][1], pts)
		for j := 0; j < total; j++ {
			dist[i][j] = d[j]
		}
	}

	// DP[mask][pos] = optimal total moves from this state.
	// mask includes already-captured pawns (bits 0..n-1 for pawns 1..n).
	// pos = current position index (0 = knight start, 1..n = pawns).
	// For Alice's turn (even popcount of captured = Alice's turn to move):
	//   maximize over next pawn p of (dist[pos][p] + solve(mask|(1<<p), p))
	// For Bob's turn (odd popcount):
	//   minimize over next pawn p of (dist[pos][p] + solve(mask|(1<<p), p))

  // Edge case: input kosong
	if n == 0 {
		return 0
	}

  // Membuat slice 2D untuk DP/tabel
	memo := make([][]int, 1<<n)
  // Iterasi seluruh elemen
	for i := range memo {
		memo[i] = make([]int, total)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var solve func(mask, pos int) int
	solve = func(mask, pos int) int {
		if mask == (1<<n)-1 {
			return 0
		}
		if memo[mask][pos] != -1 {
			return memo[mask][pos]
		}

		// Alice's turn if popcount(mask) is even (0, 2, 4, ...).
		isAlice := (bitsOn(mask) % 2) == 0

		best := -1
		if !isAlice {
			best = math.MaxInt32
		}

		for p := 0; p < n; p++ {
			if mask&(1<<p) != 0 {
				continue
			}
			// Map pawn index p (0-based in positions) to point index p+1.
			pawnIdx := p + 1
			d := dist[pos][pawnIdx]
			sub := solve(mask|(1<<p), pawnIdx)
			candidate := d + sub

			if isAlice {
				if candidate > best {
					best = candidate
				}
			} else {
				if candidate < best {
					best = candidate
				}
			}
		}

		memo[mask][pos] = best
		return best
	}

	return solve(0, 0)
}

func bitsOn(mask int) int {
	c := 0
	for mask != 0 {
		c++
		mask &= mask - 1
	}
	return c
}

func bfs(sx, sy int, targets [][2]int) []int {
	n := len(targets)
  // Membuat slice 2D untuk DP/tabel
	dist := make([][]int, boardSize)
  // Iterasi seluruh elemen
	for i := range dist {
		dist[i] = make([]int, boardSize)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}

	queue := [][2]int{{sx, sy}}
	dist[sx][sy] = 0

	for len(queue) > 0 {
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]
		for _, m := range knightMoves {
			nx, ny := x+m[0], y+m[1]
			if nx >= 0 && nx < boardSize && ny >= 0 && ny < boardSize && dist[nx][ny] == -1 {
				dist[nx][ny] = dist[x][y] + 1
				queue = append(queue, [2]int{nx, ny})
			}
		}
	}

  // Membuat slice untuk menyimpan hasil
	result := make([]int, n)
	for i, t := range targets {
		result[i] = dist[t[0]][t[1]]
	}
	return result
}
```
