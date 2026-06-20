# 0519 — Random Flip Matrix

## Deskripsi

**Soal:** [0519. Random Flip Matrix](https://leetcode.com/problems/random-flip-matrix/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1) per flip/reset amortized  
**Kompleksitas Ruang:** O(k) where k = number of flips

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #519: Random Flip Matrix
// https://leetcode.com/problems/random-flip-matrix/
// Difficulty: Medium
// Time: O(1) per flip/reset amortized
// Space: O(k) where k = number of flips

import (
	"fmt"
	"math/rand"
)

func main() {
	sol := Constructor(3, 1)
	for i := 0; i < 3; i++ {
		fmt.Println(sol.Flip())
	}
	sol.Reset()
	fmt.Println("reset done")
}

type Solution struct {
	m, n, total int
	used        map[int]int
}

func Constructor(m int, n int) Solution {
	return Solution{m: m, n: n, total: m * n, used: make(map[int]int)}
}

func (s *Solution) Flip() []int {
	randIdx := rand.Intn(s.total)
	s.total--

	// Use Fisher-Yates style swap with map for sparse tracking
	actualIdx := randIdx
	if val, ok := s.used[randIdx]; ok {
		actualIdx = val
	}
	if val, ok := s.used[s.total]; ok {
		s.used[randIdx] = val
	} else {
		s.used[randIdx] = s.total
	}

	return []int{actualIdx / s.n, actualIdx % s.n}
}

func (s *Solution) Reset() {
	s.total = s.m * s.n
	s.used = make(map[int]int)
}
```
