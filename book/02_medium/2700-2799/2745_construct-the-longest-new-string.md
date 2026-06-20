# 2745 — Construct The Longest New String

## Deskripsi

**Soal:** [2745. Construct The Longest New String](https://leetcode.com/problems/construct-the-longest-new-string/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func ConstructTheLongestNewString(x int, y int, z int) int`

## Solusi Go

```go
package main

// LeetCode #2745: Construct the Longest New String
// https://leetcode.com/problems/construct-the-longest-new-string/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func ConstructTheLongestNewString(x int, y int, z int) int {
	// AA can pair with BB, BB can pair with AA, AB can go anywhere
	// Max pairs of AA and BB: min(x, y) + extra if any left
	used := 0

	// Use pairs of AA and BB
	pairs := x
	if y < pairs {
		pairs = y
	}
	used += pairs * 2

	// If both AA and BB have remaining, can add one more
	if x > pairs {
		used++
	}
	if y > pairs {
		used++
	}

	// AB can be inserted anywhere, but consumes z
	used += z

	return used * 2
}

func main() {
	fmt.Println(ConstructTheLongestNewString(1, 1, 1))
	fmt.Println(ConstructTheLongestNewString(2, 0, 2))
	fmt.Println(ConstructTheLongestNewString(0, 0, 5))
}
```
