# 1529 — Minimum Suffix Flips

## Deskripsi

**Soal:** [1529. Minimum Suffix Flips](https://leetcode.com/problems/minimum-suffix-flips/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1529: Minimum Suffix Flips
// https://leetcode.com/problems/minimum-suffix-flips/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinFlips("10111"))
	fmt.Println(MinFlips("101"))
	fmt.Println(MinFlips("00000"))
}

func MinFlips(target string) int {
	// Time: O(N), Space: O(1)
	// Count transitions from 0 to 1 or 1 to 0
	flips := 0
	curr := byte('0') // current state of flipped prefix

  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(target); i++ {
		if target[i] != curr {
			flips++
			curr = target[i]
		}
	}

	return flips
}
```
