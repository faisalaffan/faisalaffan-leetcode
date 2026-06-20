# 0245 — Shortest Word Distance Iii

## Deskripsi

**Soal:** [0245. Shortest Word Distance Iii](https://leetcode.com/problems/shortest-word-distance-iii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func shortestWordDistance(wordsDict []string, word1 string, word2 string) int`

## Solusi Go

```go
package main

// LeetCode #245: Shortest Word Distance III
// https://leetcode.com/problems/shortest-word-distance-iii/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1)

import (
	"fmt"
	"math"
)

func shortestWordDistance(wordsDict []string, word1 string, word2 string) int {
	idx1, idx2 := -1, -1
	minDist := math.MaxInt32

	for i, word := range wordsDict {
		if word == word1 {
			idx1 = i
		}
		if word == word2 {
			if word1 == word2 {
				idx1 = idx2
			}
			idx2 = i
		}
		if idx1 != -1 && idx2 != -1 {
			dist := idx1 - idx2
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
	fmt.Println(shortestWordDistance([]string{"practice", "makes", "perfect", "coding", "makes"}, "makes", "coding"))
	fmt.Println(shortestWordDistance([]string{"practice", "makes", "perfect", "coding", "makes"}, "makes", "makes"))
	fmt.Println(shortestWordDistance([]string{"a", "a"}, "a", "a"))
}
```
