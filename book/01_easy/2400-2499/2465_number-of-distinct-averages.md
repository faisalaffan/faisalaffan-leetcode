# 2465 — Number Of Distinct Averages

## Deskripsi

**Soal:** [2465. Number Of Distinct Averages](https://leetcode.com/problems/number-of-distinct-averages/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2465: Number of Distinct Averages
// https://leetcode.com/problems/number-of-distinct-averages/
// Difficulty: Easy
// Time O(n log n) | Space O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(NumberOfDistinctAverages([]int{4, 1, 4, 0, 3, 5})) // 2
	fmt.Println(NumberOfDistinctAverages([]int{1, 100}))            // 1
}

func NumberOfDistinctAverages(nums []int) int {
	sort.Ints(nums)
	seen := map[int]bool{}
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		seen[nums[i]+nums[j]] = true
	}
	return len(seen)
}
```
