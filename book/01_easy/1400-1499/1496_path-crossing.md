# 1496 — Path Crossing

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func isPathCrossing(path string) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1496: Path Crossing
// https://leetcode.com/problems/path-crossing/
// Difficulty: Easy
//
// LeetCode submission: func isPathCrossing(path string) bool

import "fmt"

func main() {
	fmt.Println(PathCrossing("NES"))   // false
	fmt.Println(PathCrossing("NESWW")) // true
}

// Time: O(n), Space: O(n)
func PathCrossing(path string) bool {
  // HashMap: O(1) lookup
	visited := make(map[[2]int]bool)
	x, y := 0, 0
	visited[[2]int{0, 0}] = true
	for _, ch := range path {
		switch ch {
		case 'N':
			y++
		case 'S':
			y--
		case 'E':
			x++
		case 'W':
			x--
		}
		if visited[[2]int{x, y}] {
			return true
		}
		visited[[2]int{x, y}] = true
	}
	return false
}
```
