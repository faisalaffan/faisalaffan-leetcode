# 0789 — Escape The Ghosts

## Deskripsi

**Soal:** [0789. Escape The Ghosts](https://leetcode.com/problems/escape-the-ghosts/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #789: Escape The Ghosts
// https://leetcode.com/problems/escape-the-ghosts/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(escapeGhosts([][]int{{1, 0}, {0, 3}}, []int{0, 1}))
	fmt.Println(escapeGhosts([][]int{{1, 0}}, []int{2, 0}))
}

func escapeGhosts(ghosts [][]int, target []int) bool {
	myDist := abs(target[0]) + abs(target[1])

	for _, g := range ghosts {
		ghostDist := abs(g[0]-target[0]) + abs(g[1]-target[1])
		if ghostDist <= myDist {
			return false
		}
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
