# 3078 — Match Alphanumerical Pattern In Matrix I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func matchAlphanumericalPattern(board [][]int, pattern [][]byte) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(R*C*P*Q)  |  **Ruang:** O(min(26, 10))

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3078: Match Alphanumerical Pattern in Matrix I (PAID)
// https://leetcode.com/problems/match-alphanumerical-pattern-in-matrix-i/
// Difficulty: Medium [Paid]
// Time: O(R*C*P*Q) | Space: O(min(26, 10))

// Given a numeric board and a pattern of characters (digits or lowercase
// letters), find the top-left [row, col] of the first submatrix that matches
// the pattern under a bijective mapping: same letter -> same digit,
// different letters -> different digits. Digits in the pattern must match
// exactly. Return [-1, -1] if no match.

import "fmt"

func main() {
	// Test 1: No valid mapping (conflicting bijection)
	board := [][]int{{1, 2, 3}, {3, 1, 2}, {2, 3, 1}}
	pattern := [][]byte{{'a', 'b'}, {'b', 'a'}}
	fmt.Println(matchAlphanumericalPattern(board, pattern)) // [-1 -1]

	// Test 2: Simple match
	board2 := [][]int{{1, 2}, {2, 1}}
	pattern2 := [][]byte{{'a', 'b'}, {'b', 'a'}}
	fmt.Println(matchAlphanumericalPattern(board2, pattern2)) // [0 0]

	// Test 3: Letter pattern with digits in pattern
	board3 := [][]int{{5, 1, 3}, {2, 5, 4}}
	pattern3 := [][]byte{{'5', 'a'}, {'b', 'c'}}
	fmt.Println(matchAlphanumericalPattern(board3, pattern3)) // [0 0] (a=1,b=2,c=5)

	// Test 4: Pattern larger than board
	board4 := [][]int{{1}}
	pattern4 := [][]byte{{'a', 'b'}}
	fmt.Println(matchAlphanumericalPattern(board4, pattern4)) // [-1 -1]
}

func matchAlphanumericalPattern(board [][]int, pattern [][]byte) []int {
	R, C := len(board), len(board[0])
	P, Q := len(pattern), len(pattern[0])

	if R < P || C < Q {
		return []int{-1, -1}
	}

	for r := 0; r <= R-P; r++ {
		for c := 0; c <= C-Q; c++ {
			if matches(board, pattern, r, c) {
				return []int{r, c}
			}
		}
	}
	return []int{-1, -1}
}

func matches(board [][]int, pattern [][]byte, r, c int) bool {
  // HashMap: O(1) lookup
	charToDigit := make(map[byte]int)
  // HashMap: O(1) lookup
	digitToChar := make(map[int]byte)

  // Linear scan O(n)
	for i := 0; i < len(pattern); i++ {
		for j := 0; j < len(pattern[0]); j++ {
			ch := pattern[i][j]
			digit := board[r+i][c+j]

			if ch >= '0' && ch <= '9' {
				// Digit in pattern must match board digit exactly
				if int(ch-'0') != digit {
					return false
				}
			} else {
				// Letter in pattern must follow the bijective mapping
				if mapped, ok := charToDigit[ch]; ok {
					if mapped != digit {
						return false
					}
				} else if mappedChar, ok := digitToChar[digit]; ok {
					if mappedChar != ch {
						return false
					}
				} else {
					charToDigit[ch] = digit
					digitToChar[digit] = ch
				}
			}
		}
	}
	return true
}
```
