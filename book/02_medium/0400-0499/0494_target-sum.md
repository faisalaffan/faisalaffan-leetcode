# 0494 — Target Sum

## Deskripsi

**Soal:** [0494. Target Sum](https://leetcode.com/problems/target-sum/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * sum)  
**Kompleksitas Ruang:** O(sum)

**Algoritma:** Dynamic Programming (DP)

## Solusi Go

```go
package main

// LeetCode #494: Target Sum
// https://leetcode.com/problems/target-sum/
// Difficulty: Medium
// Time: O(n * sum)
// Space: O(sum)

import "fmt"

func main() {
	fmt.Println(TargetSum([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(TargetSum([]int{1}, 1))
}

func TargetSum(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum < target || (sum-target)%2 != 0 {
		return 0
	}
	negSum := (sum - target) / 2

  // Membuat slice untuk menyimpan hasil
	dp := make([]int, negSum+1)
	dp[0] = 1
	for _, num := range nums {
		for s := negSum; s >= num; s-- {
			dp[s] += dp[s-num]
		}
	}
	return dp[negSum]
}
```
