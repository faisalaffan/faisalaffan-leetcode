# 2456 — Most Popular Video Creator

## Deskripsi

**Soal:** [2456. Most Popular Video Creator](https://leetcode.com/problems/most-popular-video-creator/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2456: Most Popular Video Creator
// https://leetcode.com/problems/most-popular-video-creator/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Group by creator: track total views, best video (max views, smallest lexicographic).

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(mostPopularCreator([]string{"alice", "bob", "alice", "chris"}, []string{"one", "two", "three", "four"}, []int{5, 10, 5, 4}))
	// [[bob, two], [alice, one]]
}

type Creator struct {
	total    int
	bestID   string
	bestView int
}

func mostPopularCreator(creators []string, ids []string, views []int) [][]string {
  // Membuat map untuk pencarian O(1): key → value
	creatorsMap := make(map[string]*Creator)
	maxTotal := 0

	for i, name := range creators {
		c, ok := creatorsMap[name]
		if !ok {
			c = &Creator{bestView: math.MinInt32}
			creatorsMap[name] = c
		}
		c.total += views[i]
		if views[i] > c.bestView || (views[i] == c.bestView && ids[i] < c.bestID) {
			c.bestView = views[i]
			c.bestID = ids[i]
		}
		if c.total > maxTotal {
			maxTotal = c.total
		}
	}

  // Membuat slice 2D untuk DP/tabel
	ans := make([][]string, 0)
	for name, c := range creatorsMap {
		if c.total == maxTotal {
			ans = append(ans, []string{name, c.bestID})
		}
	}
	return ans
}
```
