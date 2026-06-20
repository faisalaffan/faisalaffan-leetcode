# 0813 — Largest Sum Of Averages

## Deskripsi

**Soal:** [0813. Largest Sum Of Averages](https://leetcode.com/problems/largest-sum-of-averages/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(k * n^2)  
**Kompleksitas Ruang:** O(k * n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #813: Largest Sum of Averages
// https://leetcode.com/problems/largest-sum-of-averages/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(LargestSumOfAverages([]int{9, 1, 2, 3, 9}, 3))
	fmt.Println(LargestSumOfAverages([]int{1, 2, 3, 4, 5, 6, 7}, 4))
	fmt.Println(LargestSumOfAverages([]int{4, 1, 7, 5, 6, 2, 3}, 4))
}

// Time: O(k * n^2) | Space: O(k * n)
func LargestSumOfAverages(nums []int, k int) float64 {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	prefix := make([]float64, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + float64(nums[i])
	}

  // Membuat slice 2D untuk DP/tabel
	dp := make([][]float64, n+1)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = make([]float64, k+1)
	}

	for i := 1; i <= n; i++ {
		dp[i][1] = prefix[i] / float64(i)
	}

	for j := 2; j <= k; j++ {
		for i := j; i <= n; i++ {
			var best float64
			for x := j - 1; x < i; x++ {
				avg := (prefix[i] - prefix[x]) / float64(i-x)
				val := dp[x][j-1] + avg
				if val > best {
					best = val
				}
			}
			dp[i][j] = best
		}
	}

	return dp[n][k]
}
```
