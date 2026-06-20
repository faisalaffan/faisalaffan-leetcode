# 3376 — Minimum Time To Break Locks I

## Deskripsi

**Soal:** [3376. Minimum Time To Break Locks I](https://leetcode.com/problems/minimum-time-to-break-locks-i/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n! * n) Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman)

## Solusi Go

```go
package main

// LeetCode #3376: Minimum Time to Break Locks I
// https://leetcode.com/problems/minimum-time-to-break-locks-i/
// Difficulty: Medium
// Time: O(n! * n) Space: O(n)

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findMinimumTime([]int{3, 4, 1}, 1)) // 4
	fmt.Println(findMinimumTime([]int{2, 5, 4}, 2)) // 5
}

func findMinimumTime(strength []int, K int) int {
	n := len(strength)
  // Membuat slice untuk menyimpan hasil
	used := make([]bool, n)
	ans := math.MaxInt32

	var dfs func(idx int, time int, X int)
	dfs = func(idx int, time int, X int) {
		if idx == n {
			if time < ans {
				ans = time
			}
			return
		}
		if time >= ans {
			return
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			need := (strength[i] + X - 1) / X
			dfs(idx+1, time+need, X+K)
			used[i] = false
		}
	}

	dfs(0, 0, 1)
	return ans
}
```
