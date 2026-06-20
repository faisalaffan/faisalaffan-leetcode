# 0470 — Implement Rand10 Using Rand7

## Deskripsi

**Soal:** [0470. Implement Rand10 Using Rand7](https://leetcode.com/problems/implement-rand10-using-rand7/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1) expected  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #470: Implement Rand10() Using Rand7()
// https://leetcode.com/problems/implement-rand10-using-rand7/
// Difficulty: Medium
// Time: O(1) expected
// Space: O(1)

import (
	"fmt"
	"math/rand"
)

func main() {
	// Test by generating a few random values
	for i := 0; i < 5; i++ {
		fmt.Println(ImplementRandOneZeroUsingRandSeven())
	}
}

func ImplementRandOneZeroUsingRandSeven() int {
	// Rejection sampling: (rand7()-1)*7 + rand7() gives 1-49 uniformly
	// Accept values 1-40 and map to 1-10
	for {
		val := (rand7()-1)*7 + rand7()
		if val <= 40 {
			return (val-1)%10 + 1
		}
	}
}

func rand7() int {
	return rand.Intn(7) + 1
}
```
