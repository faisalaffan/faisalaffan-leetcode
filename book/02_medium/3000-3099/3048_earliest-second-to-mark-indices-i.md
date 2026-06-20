# 3048 — Earliest Second To Mark Indices I

## Deskripsi

**Soal:** [3048. Earliest Second To Mark Indices I](https://leetcode.com/problems/earliest-second-to-mark-indices-i/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m log m)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3048: Earliest Second to Mark Indices I
// https://leetcode.com/problems/earliest-second-to-mark-indices-i/
// Difficulty: Medium
// Time: O(m log m) | Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(earliestSecondToMarkIndices([]int{2, 2, 0}, []int{2, 2, 2, 2, 3, 2, 2, 1}))
	fmt.Println(earliestSecondToMarkIndices([]int{1, 3}, []int{1, 1, 1, 2, 1, 1, 1}))
	fmt.Println(earliestSecondToMarkIndices([]int{0, 1}, []int{2, 2, 2}))
}

func earliestSecondToMarkIndices(nums []int, changeIndices []int) int {
	n, m := len(nums), len(changeIndices)
	ans := sort.Search(m+1, func(t int) bool {
		if t == 0 {
			return false
		}
  // Membuat slice untuk menyimpan hasil
		last := make([]int, n+1)
		for s, idx := range changeIndices[:t] {
			last[idx] = s
		}
		for i := 1; i <= n; i++ {
			if last[i] == 0 && i != changeIndices[0] {
				return false
			}
		}
		decrement := 0
		marked := 0
		for s, idx := range changeIndices[:t] {
			if last[idx] == s {
				if decrement < nums[idx-1] {
					return false
				}
				decrement -= nums[idx-1]
				marked++
			} else {
				decrement++
			}
		}
		return marked == n
	})
	if ans > m {
		return -1
	}
	return ans
}
```
