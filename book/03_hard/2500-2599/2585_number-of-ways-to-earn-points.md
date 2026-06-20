# 2585 — Number Of Ways To Earn Points

## Deskripsi

**Soal:** [2585. Number Of Ways To Earn Points](https://leetcode.com/problems/number-of-ways-to-earn-points/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func waysToReachTarget(target int, types [][]int) int`

## Solusi Go

```go
package main

// LeetCode #2585: Number of Ways to Earn Points
// https://leetcode.com/problems/number-of-ways-to-earn-points/
// Difficulty: Hard

import "fmt"

// waysToReachTarget counts ways to earn exactly `target` points.
// Each type[i] = [count_i, marks_i]: up to count_i problems each worth marks_i.
//
// Knapsack DP: dp[j] = ways to earn j points.
// For each type, process in reverse order adding 1..count of that problem
// to avoid reusing the same problem.
//
// Complexity: O(target * sum(count)) time, O(target) space
func waysToReachTarget(target int, types [][]int) int {
	const mod = 1_000_000_007
  // Membuat slice untuk menyimpan hasil
	dp := make([]int, target+1)
	dp[0] = 1

	for _, t := range types {
		count, marks := t[0], t[1]
		for j := target; j >= 0; j-- {
			if dp[j] == 0 {
				continue
			}
			for used := 1; used <= count; used++ {
				points := used * marks
				if j+points > target {
					break
				}
				dp[j+points] = (dp[j+points] + dp[j]) % mod
			}
		}
	}

	return dp[target]
}

func main() {
	// Example from LeetCode
	fmt.Println("Test 1: target=6, types=[[6,1],[3,2],[2,3]] ->",
		waysToReachTarget(6, [][]int{{6, 1}, {3, 2}, {2, 3}})) // 7

	// Additional test cases
	fmt.Println("Test 2: target=5, types=[[5,1]] ->",
		waysToReachTarget(5, [][]int{{5, 1}})) // 1

	fmt.Println("Test 3: target=10, types=[[10,1],[5,2]] ->",
		waysToReachTarget(10, [][]int{{10, 1}, {5, 2}})) // 2

	fmt.Println("Test 4: target=0, types=[[1,1]] ->",
		waysToReachTarget(0, [][]int{{1, 1}})) // 1

	fmt.Println("Test 5: target=3, types=[[1,1],[1,1],[1,1]] ->",
		waysToReachTarget(3, [][]int{{1, 1}, {1, 1}, {1, 1}})) // 1
}
```
