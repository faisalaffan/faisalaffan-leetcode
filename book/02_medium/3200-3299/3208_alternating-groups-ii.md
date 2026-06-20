# 3208 — Alternating Groups Ii

## Deskripsi

**Soal:** [3208. Alternating Groups Ii](https://leetcode.com/problems/alternating-groups-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func numberOfAlternatingGroups(colors []int, k int) int`

## Solusi Go

```go
package main

// LeetCode #3208: Alternating Groups II
// https://leetcode.com/problems/alternating-groups-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func numberOfAlternatingGroups(colors []int, k int) int {
	n := len(colors)
	ans := 0
	len := 1

	for i := 1; i < n+k-1; i++ {
		if colors[i%n] != colors[(i-1)%n] {
			len++
		} else {
			len = 1
		}
		if len >= k {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfAlternatingGroups([]int{0, 1, 0, 1, 0}, 3)) // Expected: 3
	fmt.Println(numberOfAlternatingGroups([]int{0, 1, 0, 0, 1}, 3)) // Expected: 2
}
```
