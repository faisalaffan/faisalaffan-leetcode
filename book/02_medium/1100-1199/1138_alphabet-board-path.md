# 1138 — Alphabet Board Path

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func alphabetBoardPath(target string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"strings"
)

// LeetCode #1138: Alphabet Board Path
// https://leetcode.com/problems/alphabet-board-path/

// On an alphabet board (5x6, 'a' at (0,0), 'z' at (5,0)),
// we start at position 'a'. We can move U/D/L/R and append '!' to pick.
// Return the shortest sequence to spell the target string.

// The only tricky part is 'z' which is alone in the last row.
// When moving from or to 'z', we must avoid out-of-bounds moves.
// Strategy: always move U first, then L/R, then D (to handle 'z' adjacency).

// Time complexity: O(n * 5) = O(n) where n = len(target)
// Space complexity: O(n) for the result

func alphabetBoardPath(target string) string {
	var result strings.Builder
	curR, curC := 0, 0

	for _, ch := range target {
		idx := int(ch - 'a')
		targetR := idx / 5
		targetC := idx % 5

		// For 'z', we need to move up first then left/right
		// because there's no cell below 'z' and no cell to the right.
		// General case: move U first (never out of bounds), then L/R, then D
		for curR > targetR {
			result.WriteByte('U')
			curR--
		}
		for curC > targetC {
			result.WriteByte('L')
			curC--
		}
		for curC < targetC {
			result.WriteByte('R')
			curC++
		}
		for curR < targetR {
			result.WriteByte('D')
			curR++
		}

		result.WriteByte('!')
	}

	return result.String()
}

func main() {
	// Test case 1
	fmt.Printf("alphabetBoardPath(%q) = %q (expected: %q)\n",
		"leet", alphabetBoardPath("leet"), "DDR!UURRR!!DDD!")

	// Test case 2
	fmt.Printf("alphabetBoardPath(%q) = %q (expected: %q)\n",
		"code", alphabetBoardPath("code"), "RR!DDR!UUL!R!")

	// Test case 3: Single character
	fmt.Printf("alphabetBoardPath(%q) = %q (expected: %q)\n",
		"a", alphabetBoardPath("a"), "!")
}
```
