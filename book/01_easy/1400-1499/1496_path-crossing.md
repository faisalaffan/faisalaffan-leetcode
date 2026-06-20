# 1496 — Path Crossing

## Deskripsi

**Soal:** [1496. Path Crossing](https://leetcode.com/problems/path-crossing/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func isPathCrossing(path string) bool`

## Solusi Go

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
  // Membuat map untuk pencarian O(1): key → value
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
