# 0243 — Shortest Word Distance

## Deskripsi

**Soal:** [0243. Shortest Word Distance](https://leetcode.com/problems/shortest-word-distance/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func ShortestDistance(wordsDict []string, word1 string, word2 string) int`

## Solusi Go

```go
package main

// LeetCode #243: Shortest Word Distance
// https://leetcode.com/problems/shortest-word-distance/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"math"
)

// Time: O(n) | Space: O(1)
func ShortestDistance(wordsDict []string, word1 string, word2 string) int {
	i1, i2 := -1, -1
	minDist := math.MaxInt32
	for i, w := range wordsDict {
		if w == word1 {
			i1 = i
		}
		if w == word2 {
			i2 = i
		}
		if i1 != -1 && i2 != -1 {
			dist := i1 - i2
			if dist < 0 {
				dist = -dist
			}
			if dist < minDist {
				minDist = dist
			}
		}
	}
	return minDist
}

func main() {
	fmt.Println(ShortestDistance([]string{"practice", "makes", "perfect", "coding", "makes"}, "coding", "practice"))
	fmt.Println(ShortestDistance([]string{"practice", "makes", "perfect", "coding", "makes"}, "makes", "coding"))
}
```
