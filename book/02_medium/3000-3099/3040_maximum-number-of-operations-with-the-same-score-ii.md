# 3040 — Maximum Number Of Operations With The Same Score Ii

## Deskripsi

**Soal:** [3040. Maximum Number Of Operations With The Same Score Ii](https://leetcode.com/problems/maximum-number-of-operations-with-the-same-score-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n^2)

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman)

## Solusi Go

```go
package main

// LeetCode #3040: Maximum Number of Operations With the Same Score II
// https://leetcode.com/problems/maximum-number-of-operations-with-the-same-score-ii/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n^2)

import "fmt"

func main() {
	fmt.Println(maxOperations3040([]int{3, 2, 1, 2, 3, 4}))
	fmt.Println(maxOperations3040([]int{3, 2, 6, 1, 4}))
}

func maxOperations3040(nums []int) int {
	n := len(nums)
	memo := map[[2]int]int{}

	var dfs func(l, r, target int) int
	dfs = func(l, r, target int) int {
		if l >= r {
			return 0
		}
		key := [2]int{l, r}
		if v, ok := memo[key]; ok {
			return v
		}
		best := 0
		if nums[l]+nums[l+1] == target {
			if res := 1 + dfs(l+2, r, target); res > best {
				best = res
			}
		}
		if nums[r-1]+nums[r] == target {
			if res := 1 + dfs(l, r-2, target); res > best {
				best = res
			}
		}
		if nums[l]+nums[r] == target {
			if res := 1 + dfs(l+1, r-1, target); res > best {
				best = res
			}
		}
		memo[key] = best
		return best
	}

	ans := 0
	// Try all 3 possible targets from first operation
	targets := []int{nums[0] + nums[1], nums[n-2] + nums[n-1], nums[0] + nums[n-1]}
	for _, t := range targets {
		memo = map[[2]int]int{}
		if res := dfs(0, n-1, t); res > ans {
			ans = res
		}
	}
	return ans
}
```
