# 2924 — Find Champion Ii

## Deskripsi

**Soal:** [2924. Find Champion Ii](https://leetcode.com/problems/find-champion-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n+m)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2924: Find Champion II
// https://leetcode.com/problems/find-champion-ii/
// Difficulty: Medium
// Time: O(n+m) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(findChampionII(3, [][]int{{0, 1}, {1, 2}}))
	fmt.Println(findChampionII(4, [][]int{{0, 2}, {1, 3}, {1, 2}}))
	fmt.Println(findChampionII(2, [][]int{{0, 1}}))
}

func findChampionII(n int, edges [][]int) int {
  // Membuat slice untuk menyimpan hasil
	indeg := make([]int, n)
	for _, e := range edges {
		indeg[e[1]]++
	}
	ans, cnt := -1, 0
	for i, x := range indeg {
		if x == 0 {
			cnt++
			ans = i
		}
	}
	if cnt == 1 {
		return ans
	}
	return -1
}
```
