# 2498 — Frog Jump Ii

## Deskripsi

**Soal:** [2498. Frog Jump Ii](https://leetcode.com/problems/frog-jump-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2498: Frog Jump II
// https://leetcode.com/problems/frog-jump-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Frog can jump forward, max distance = max of (stones[i+2] - stones[i]) for i in 0..n-3
// and also stones[1] - stones[0] and stones[n-1] - stones[n-2].

import "fmt"

func main() {
	fmt.Println(maxJump([]int{0, 2, 5, 6, 7})) // 5
	fmt.Println(maxJump([]int{0, 3, 9}))        // 9
}

func maxJump(stones []int) int {
	n := len(stones)
	ans := stones[1] - stones[0]
	if n > 2 {
		ans = stones[n-1] - stones[n-2]
	}
	for i := 2; i < n; i++ {
		diff := stones[i] - stones[i-2]
		if diff > ans {
			ans = diff
		}
	}
	return ans
}
```
