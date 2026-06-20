# 1079 — Letter Tile Possibilities

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func numTilePossibilities(tiles string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Backtracking

**Waktu:** O(n!) where n = len(tiles) worst case  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Backtracking** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1079: Letter Tile Possibilities
// https://leetcode.com/problems/letter-tile-possibilities/
// Difficulty: Medium
//
// Approach: Backtracking with frequency count
// Time: O(n!) where n = len(tiles) worst case
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(numTilePossibilities("AAB")) // 8
	fmt.Println(numTilePossibilities("AAABBC")) // 188
}

func numTilePossibilities(tiles string) int {
  // Alokasi slice
	freq := make([]int, 26)
	for _, c := range tiles {
		freq[c-'A']++
	}

	var dfs func() int
	dfs = func() int {
		count := 0
		for i := 0; i < 26; i++ {
			if freq[i] == 0 {
				continue
			}
			count++
			freq[i]--
			count += dfs()
			freq[i]++
		}
		return count
	}

	return dfs()
}
```
