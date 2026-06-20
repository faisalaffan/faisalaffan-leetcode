# 2515 — Shortest Distance To Target String In A Circular Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func ShortestDistanceToTargetStringInACircularArray(words []string, target string, startIndex int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2515: Shortest Distance to Target String in a Circular Array
// https://leetcode.com/problems/shortest-distance-to-target-string-in-a-circular-array/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(ShortestDistanceToTargetStringInACircularArray([]string{"hello", "i", "am", "leetcode", "hello"}, "hello", 1)) // 1
	fmt.Println(ShortestDistanceToTargetStringInACircularArray([]string{"a", "b", "leetcode"}, "leetcode", 0))                // 1
}

func ShortestDistanceToTargetStringInACircularArray(words []string, target string, startIndex int) int {
	n := len(words)
	minDist := n

	for i, w := range words {
		if w == target {
			dist := abs(i - startIndex)
			if dist > n-dist {
				dist = n - dist
			}
			if dist < minDist {
				minDist = dist
			}
		}
	}
	return minDist
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
