# 0398 — Random Pick Index

## Deskripsi

**Soal:** [0398. Random Pick Index](https://leetcode.com/problems/random-pick-index/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n) for init, O(1) for pick  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func Constructor(nums []int) Solution`

## Solusi Go

```go
package main

// LeetCode #398: Random Pick Index
// https://leetcode.com/problems/random-pick-index/
// Difficulty: Medium
// Time: O(n) for init, O(1) for pick | Space: O(n)

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}

func (s *Solution) Pick(target int) int {
	// Reservoir sampling
	count := 0
	result := 0
	for i, num := range s.nums {
		if num == target {
			count++
			if rand.Intn(count) == 0 {
				result = i
			}
		}
	}
	return result
}

func main() {
	sol := Constructor([]int{1, 2, 3, 3, 3})
	counts := map[int]int{}
	for i := 0; i < 30000; i++ {
		counts[sol.Pick(3)]++
	}
	fmt.Println("Counts for target=3:", counts)
	// Expected: roughly 10000 each for indices 2,3,4
}
```
