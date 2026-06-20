# 0514 — Freedom Trail

## Deskripsi

**Soal:** [0514. Freedom Trail](https://leetcode.com/problems/freedom-trail/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP), LIS (Longest Increasing Subsequence)

## Solusi Go

```go
package main

// LeetCode #514: Freedom Trail
// https://leetcode.com/problems/freedom-trail/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println(findRotateSteps("godding", "gd")) // Expected: 4
}

func findRotateSteps(ring string, key string) int {
	m, n := len(ring), len(key)
	// pos[c] = list of indices in ring where character c appears
  // Membuat slice 2D untuk DP/tabel
	pos := make([][]int, 26)
	for i := 0; i < m; i++ {
		c := ring[i] - 'a'
		pos[c] = append(pos[c], i)
	}

	// dp[j] = min steps to spell up to current key char ending at ring index j
  // Membuat slice untuk menyimpan hasil
	dp := make([]int, m)
	for j := 0; j < m; j++ {
		if ring[j] == key[0] {
			dp[j] = minDist(j, 0, m) + 1 // steps to rotate + press
		} else {
			dp[j] = 1 << 30 // large number
		}
	}

	for i := 1; i < n; i++ {
  // Membuat slice untuk menyimpan hasil
		next := make([]int, m)
		for j := 0; j < m; j++ {
			next[j] = 1 << 30
		}
		for _, j := range pos[key[i]-'a'] {
			// from any previous position where we could have been
			for _, k := range pos[key[i-1]-'a'] {
				cost := dp[k] + minDist(j, k, m) + 1
				if cost < next[j] {
					next[j] = cost
				}
			}
		}
		dp = next
	}

	ans := 1 << 30
	for j := 0; j < m; j++ {
		if dp[j] < ans {
			ans = dp[j]
		}
	}
	return ans
}

func minDist(i, j, m int) int {
	d := i - j
	if d < 0 {
		d = -d
	}
	if d > m-d {
		return m - d
	}
	return d
}
```
