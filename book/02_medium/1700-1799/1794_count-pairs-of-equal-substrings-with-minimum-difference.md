# 1794 — Count Pairs Of Equal Substrings With Minimum Difference

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func countQuadruples(firstString string, secondString string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1794: Count Pairs of Equal Substrings With Minimum Difference
// https://leetcode.com/problems/count-pairs-of-equal-substrings-with-minimum-difference/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1)

import "fmt"

func countQuadruples(firstString string, secondString string) int {
  // Alokasi slice
	firstPos := make([]int, 26)
  // Alokasi slice
	lastPos := make([]int, 26)
  // Range loop
	for i := range firstPos {
		firstPos[i] = -1
		lastPos[i] = -1
	}

	for i, ch := range firstString {
		idx := ch - 'a'
		if firstPos[idx] == -1 {
			firstPos[idx] = i
		}
	}
	for i, ch := range secondString {
		idx := ch - 'a'
		lastPos[idx] = i
	}

	minDiff := 1 << 30
	count := 0

	for i := 0; i < 26; i++ {
		if firstPos[i] != -1 && lastPos[i] != -1 {
			diff := firstPos[i] - lastPos[i]
			if diff < minDiff {
				minDiff = diff
				count = 1
			} else if diff == minDiff {
				count++
			}
		}
	}
	return count
}

func main() {
	fmt.Println(countQuadruples("abcd", "bcd")) // test 1
	fmt.Println(countQuadruples("abc", "abc")) // test 2
	fmt.Println(countQuadruples("abb", "b")) // test 3
}
```
