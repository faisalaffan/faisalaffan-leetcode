# 3449 — Maximize The Minimum Game Score

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxScore(points []int, m int) int64
```

> **💡 Hint:** Binary search on the minimum score. For a candidate

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3449: Maximize the Minimum Game Score
// https://leetcode.com/problems/maximize-the-minimum-game-score/
// Difficulty: Hard
//
// Given points array and m total moves. Each move: select index i
// and add points[i] to score. After selecting index i, next move
// must be at i-1 or i+1 (can't stay). Maximize the minimum total
// score after all m moves.
//
// Approach: Binary search on the minimum score. For a candidate
// min score, check if we can achieve it with m moves using a
// greedy strategy.

import "fmt"

func main() {
	// Example 1
	fmt.Println(maxScore([]int{2, 1, 3}, 4))
	// Example 2
	fmt.Println(maxScore([]int{1, 1, 1}, 3))
	// Edge: m = 0
	fmt.Println(maxScore([]int{5, 4, 3}, 0))
}

func maxScore(points []int, m int) int64 {
	n := len(points)
	if m < n {
		return 0
	}

	// Check if we can achieve at least target score at each position
	can := func(target int64) bool {
  // Alokasi slice integer
		needed := make([]int64, n)
		for i, p := range points {
			// Moves needed at position i to reach target
			needed[i] = (target + int64(p) - 1) / int64(p)
		}

		totalMoves := int64(0)
		extra := int64(0) // moves carried from previous position

		for i := 0; i < n; i++ {
			if i == n-1 {
				if needed[i] > extra+1 {
					totalMoves += (needed[i] - extra - 1) * 2
				}
				break
			}

			if extra >= needed[i] {
				extra = 0
				totalMoves++
				continue
			}

			need := needed[i] - extra
			// We need need-1 moves between i and i+1
			// Each such move costs 2 (go to i, back to i+1, or vice versa)
			movesHere := need*2 - 1
			totalMoves += movesHere
			extra = need - 1
		}

		return totalMoves <= int64(m)
	}

	var lo, hi int64 = 0, int64(1e18)
	for lo < hi {
		mid := lo + (hi-lo+1)/2
		if can(mid) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return lo
}
```
