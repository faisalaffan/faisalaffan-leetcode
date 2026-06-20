# 2975 — Maximum Square Area By Removing Fences From A Field

## Deskripsi

**Soal:** [2975. Maximum Square Area By Removing Fences From A Field](https://leetcode.com/problems/maximum-square-area-by-removing-fences-from-a-field/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(h^2 + v^2)  
**Kompleksitas Ruang:** O(h^2 + v^2)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2975: Maximum Square Area by Removing Fences From a Field
// https://leetcode.com/problems/maximum-square-area-by-removing-fences-from-a-field/
// Difficulty: Medium
// Time: O(h^2 + v^2) | Space: O(h^2 + v^2)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximizeSquareAreaField(4, 3, []int{2, 3}, []int{2}))
	fmt.Println(maximizeSquareAreaField(6, 4, []int{3, 5}, []int{2, 3}))
}

func maximizeSquareAreaField(m int, n int, hFences []int, vFences []int) int {
	const mod = 1_000_000_007

	getGaps := func(fences []int, limit int) map[int]bool {
		arr := append(fences, 1, limit)
		sort.Ints(arr)
		gaps := map[int]bool{}
  // Loop standar: indeks 0 sampai n-1
		for i := 0; i < len(arr); i++ {
			for j := i + 1; j < len(arr); j++ {
				gaps[arr[j]-arr[i]] = true
			}
		}
		return gaps
	}

	hGaps := getGaps(hFences, m)
	vGaps := getGaps(vFences, n)

	maxLen := 0
	for d := range hGaps {
		if vGaps[d] && d > maxLen {
			maxLen = d
		}
	}
	if maxLen == 0 {
		return -1
	}
	return (maxLen * maxLen) % mod
}
```
