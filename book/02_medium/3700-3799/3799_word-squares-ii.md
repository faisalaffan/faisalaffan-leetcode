# 3799 — Word Squares Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func WordSquaresIi(words []string) [][]string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Sorting

**Waktu:** O(N^4 * L)  |  **Ruang:** O(N)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3799: Word Squares II
// https://leetcode.com/problems/word-squares-ii/
// Difficulty: Medium
// Time: O(N^4 * L) | Space: O(N)
// Approach: Generate all valid 4-word squares [top, left, right, bottom]
// satisfying corner constraints: top[0]==left[0], top[3]==right[0],
// bottom[0]==left[3], bottom[3]==right[3]. All 4 words must be distinct.

import (
	"fmt"
	"sort"
)

func WordSquaresIi(words []string) [][]string {
	n := len(words)
	result := [][]string{}

	for ti := 0; ti < n; ti++ {
		top := words[ti]
		for li := 0; li < n; li++ {
			if li == ti {
				continue
			}
			left := words[li]
			if left[0] != top[0] {
				continue
			}
			for ri := 0; ri < n; ri++ {
				if ri == ti || ri == li {
					continue
				}
				right := words[ri]
				if right[0] != top[3] {
					continue
				}
				for bi := 0; bi < n; bi++ {
					if bi == ti || bi == li || bi == ri {
						continue
					}
					bottom := words[bi]
					if bottom[0] != left[3] || bottom[3] != right[3] {
						continue
					}
					square := []string{top, left, right, bottom}
					result = append(result, square)
				}
			}
		}
	}

	// Sort by (top, left, right, bottom) lexicographically
  // Custom sort
	sort.Slice(result, func(i, j int) bool {
		for k := 0; k < 4; k++ {
			if result[i][k] != result[j][k] {
				return result[i][k] < result[j][k]
			}
		}
		return false
	})

	return result
}

func main() {
	// Example 1
	words1 := []string{"able", "area", "echo", "also"}
	result1 := WordSquaresIi(words1)
	fmt.Println("Result 1:")
	for _, sq := range result1 {
		fmt.Printf("  %v\n", sq)
	}

	// Example 2
	words2 := []string{"ball", "area", "lead", "lady"}
	result2 := WordSquaresIi(words2)
	fmt.Println("Result 2:")
	for _, sq := range result2 {
		fmt.Printf("  %v\n", sq)
	}
}
```
