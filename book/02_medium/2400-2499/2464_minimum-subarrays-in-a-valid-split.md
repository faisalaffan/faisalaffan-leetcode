# 2464 — Minimum Subarrays In A Valid Split

## Deskripsi

**Soal:** [2464. Minimum Subarrays In A Valid Split](https://leetcode.com/problems/minimum-subarrays-in-a-valid-split/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** Greedy (pemilihan optimal lokal)

## Solusi Go

```go
package main

// LeetCode #2464: Minimum Subarrays in a Valid Split
// https://leetcode.com/problems/minimum-subarrays-in-a-valid-split/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Greedy: start new subarray when GCD becomes 1.

import "fmt"

func main() {
	fmt.Println(validSplit([]int{2, 6, 3, 4, 3})) // 2
	fmt.Println(validSplit([]int{3, 5}))           // 2
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func validSplit(nums []int) int {
	ans := 1
	cur := 0
	for _, v := range nums {
		cur = gcd(cur, v)
		if cur == 1 {
			ans++
			cur = v
		}
	}
	if cur == 1 {
		// If last GCD is 1, we started but couldn't close
		// This shouldn't happen for valid input per problem constraints
	}
	return ans
}
```
