# 2849 — Determine If A Cell Is Reachable At A Given Time

## Deskripsi

**Soal:** [2849. Determine If A Cell Is Reachable At A Given Time](https://leetcode.com/problems/determine-if-a-cell-is-reachable-at-a-given-time/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func DetermineIfACellIsReachableAtAGivenTime(sx int, sy int, fx int, fy int, t int) bool`

## Solusi Go

```go
package main

// LeetCode #2849: Determine if a Cell Is Reachable at a Given Time
// https://leetcode.com/problems/determine-if-a-cell-is-reachable-at-a-given-time/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func DetermineIfACellIsReachableAtAGivenTime(sx int, sy int, fx int, fy int, t int) bool {
	dx := sx - fx
	if dx < 0 {
		dx = -dx
	}
	dy := sy - fy
	if dy < 0 {
		dy = -dy
	}
	minDist := dx
	if dy > minDist {
		minDist = dy
	}
	if minDist == 0 {
		return t != 1
	}
	return t >= minDist
}

func main() {
	fmt.Println(DetermineIfACellIsReachableAtAGivenTime(1, 2, 3, 4, 3))
	fmt.Println(DetermineIfACellIsReachableAtAGivenTime(1, 2, 1, 2, 1))
	fmt.Println(DetermineIfACellIsReachableAtAGivenTime(1, 1, 1, 1, 0))
}
```
