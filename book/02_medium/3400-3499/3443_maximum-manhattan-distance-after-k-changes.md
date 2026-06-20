# 3443 — Maximum Manhattan Distance After K Changes

## Deskripsi

**Soal:** [3443. Maximum Manhattan Distance After K Changes](https://leetcode.com/problems/maximum-manhattan-distance-after-k-changes/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n) Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func maxDistance(s string, k int) int`

## Solusi Go

```go
package main

// LeetCode #3443: Maximum Manhattan Distance After K Changes
// https://leetcode.com/problems/maximum-manhattan-distance-after-k-changes/
// Difficulty: Medium
// Time: O(n) Space: O(1)

import "fmt"

func maxDistance(s string, k int) int {
	x, y := 0, 0
	ans := 0
	for i, ch := range s {
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
		dist := abs(x) + abs(y) + 2*k
		if i+1 < dist {
			dist = i + 1
		}
		if dist > ans {
			ans = dist
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(maxDistance("NWSE", 1))   // 3
	fmt.Println(maxDistance("NSWWEW", 3)) // 6
}
```
