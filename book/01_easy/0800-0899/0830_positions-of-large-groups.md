# 0830 — Positions Of Large Groups

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func largeGroupPositions(s string) [][]int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n). Space: O(1) excluding output.  |  **Ruang:** O(1) excluding output.


## 💻 Solusi Go

```go
package main

// LeetCode #830: Positions of Large Groups
// https://leetcode.com/problems/positions-of-large-groups/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(largeGroupPositions("abbxxxxzzy"))    // [[3,6]]
	fmt.Println(largeGroupPositions("abc"))           // []
	fmt.Println(largeGroupPositions("abcdddeeeeaabbbcd")) // [[3,5],[6,9],[12,14]]
}

// largeGroupPositions finds all large groups (consecutive identical characters of length >= 3).
// Time: O(n). Space: O(1) excluding output.
func largeGroupPositions(s string) [][]int {
  // Matriks 2D
	result := make([][]int, 0)
	start := 0
	for i := 1; i <= len(s); i++ {
		if i == len(s) || s[i] != s[start] {
			if i-start >= 3 {
				result = append(result, []int{start, i - 1})
			}
			start = i
		}
	}
	return result
}
```
