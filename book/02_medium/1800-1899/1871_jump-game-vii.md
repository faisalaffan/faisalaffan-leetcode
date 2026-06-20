# 1871 — Jump Game Vii

## Deskripsi

**Soal:** [1871. Jump Game Vii](https://leetcode.com/problems/jump-game-vii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Dynamic Programming (DP)

## Solusi Go

```go
package main

// LeetCode #1871: Jump Game VII
// https://leetcode.com/problems/jump-game-vii/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(CanReach("011010", 2, 3))
	fmt.Println(CanReach("01101110", 2, 3))
	fmt.Println(CanReach("00", 1, 1))
}

// Time: O(n), Space: O(n)
func CanReach(s string, minJump int, maxJump int) bool {
	n := len(s)
	if s[n-1] != '0' {
		return false
	}

  // Membuat slice untuk menyimpan hasil
	dp := make([]bool, n)
  // Membuat slice untuk menyimpan hasil
	prefix := make([]int, n+1)
	dp[0] = true
	prefix[1] = 1

	for i := 1; i < n; i++ {
		if s[i] == '1' {
			prefix[i+1] = prefix[i]
			continue
		}
		left := max(0, i-maxJump)
		right := i - minJump
		if right >= left {
			reachable := prefix[right+1] - prefix[left] > 0
			if reachable {
				dp[i] = true
			}
		}
		prefix[i+1] = prefix[i]
		if dp[i] {
			prefix[i+1]++
		}
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
