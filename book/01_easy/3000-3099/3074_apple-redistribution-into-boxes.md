# 3074 — Apple Redistribution Into Boxes

## Deskripsi

**Soal:** [3074. Apple Redistribution Into Boxes](https://leetcode.com/problems/apple-redistribution-into-boxes/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3074: Apple Redistribution into Boxes
// https://leetcode.com/problems/apple-redistribution-into-boxes/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	// LeetCode name: minimumBoxes
	fmt.Println(AppleRedistributionIntoBoxes([]int{1, 3, 2}, []int{4, 3, 1, 5, 2})) // 2
	fmt.Println(AppleRedistributionIntoBoxes([]int{5, 5, 5}, []int{2, 4, 2, 7}))    // 4
}

// Time: O(n log n) | Space: O(1)
// LeetCode submission name: minimumBoxes
func AppleRedistributionIntoBoxes(apples []int, capacity []int) int {
	totalApples := 0
	for _, a := range apples {
		totalApples += a
	}
	sort.Sort(sort.Reverse(sort.IntSlice(capacity)))
	boxes := 0
	for _, c := range capacity {
		boxes++
		totalApples -= c
		if totalApples <= 0 {
			return boxes
		}
	}
	return boxes
}
```
