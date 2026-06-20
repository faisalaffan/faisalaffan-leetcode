# 0646 — Maximum Length Of Pair Chain

## Deskripsi

**Soal:** [0646. Maximum Length Of Pair Chain](https://leetcode.com/problems/maximum-length-of-pair-chain/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n) for sorting  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #646: Maximum Length of Pair Chain
// https://leetcode.com/problems/maximum-length-of-pair-chain/
// Difficulty: Medium
// Time: O(n log n) for sorting
// Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findLongestChain([][]int{{1, 2}, {2, 3}, {3, 4}}))
	fmt.Println(findLongestChain([][]int{{1, 2}, {7, 8}, {4, 5}}))
}

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})

	count := 0
	curEnd := -1 << 31

	for _, pair := range pairs {
		if pair[0] > curEnd {
			curEnd = pair[1]
			count++
		}
	}

	return count
}
```
