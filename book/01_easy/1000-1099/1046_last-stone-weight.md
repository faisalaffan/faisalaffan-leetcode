# 1046 — Last Stone Weight

## Deskripsi

**Soal:** [1046. Last Stone Weight](https://leetcode.com/problems/last-stone-weight/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1046: Last Stone Weight
// https://leetcode.com/problems/last-stone-weight/
// Difficulty: Easy
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1})) // 1
	fmt.Println(lastStoneWeight([]int{1}))                 // 1
	fmt.Println(lastStoneWeight([]int{2, 2}))              // 0
}

// LeetCode submission: lastStoneWeight
func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		sort.Ints(stones)
		n := len(stones)
		if stones[n-1] == stones[n-2] {
			stones = stones[:n-2]
		} else {
			stones[n-2] = stones[n-1] - stones[n-2]
			stones = stones[:n-1]
		}
	}
	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}
```
