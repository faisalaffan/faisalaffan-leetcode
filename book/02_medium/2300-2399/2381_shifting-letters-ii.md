# 2381 — Shifting Letters Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func shiftingLetters(s string, shifts [][]int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n + m)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #2381: Shifting Letters II
// https://leetcode.com/problems/shifting-letters-ii/
// Difficulty: Medium
// Time: O(n + m) | Space: O(n)
// Difference array to apply range shifts efficiently.

import "fmt"

func main() {
	fmt.Println(shiftingLetters("abc", [][]int{{0, 1, 0}, {1, 2, 1}, {0, 2, 1}})) // "ace"
	fmt.Println(shiftingLetters("dztz", [][]int{{0, 0, 0}, {1, 1, 1}}))       // "catz"
}

func shiftingLetters(s string, shifts [][]int) string {
	n := len(s)
  // Alokasi slice
	diff := make([]int, n+1)
	for _, sh := range shifts {
		start, end, dir := sh[0], sh[1], sh[2]
		if dir == 1 {
			diff[start]++
			diff[end+1]--
		} else {
			diff[start]--
			diff[end+1]++
		}
	}

	cur := 0
	res := make([]byte, n)
	for i, ch := range s {
		cur += diff[i]
		shift := ((int(ch-'a')+cur)%26 + 26) % 26
		res[i] = byte('a' + shift)
	}
	return string(res)
}
```
