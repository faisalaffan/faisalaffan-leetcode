# 2470 — Number Of Subarrays With Lcm Equal To K

## Deskripsi

**Soal:** [2470. Number Of Subarrays With Lcm Equal To K](https://leetcode.com/problems/number-of-subarrays-with-lcm-equal-to-k/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2) worst-case  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2470: Number of Subarrays With LCM Equal to K
// https://leetcode.com/problems/number-of-subarrays-with-lcm-equal-to-k/
// Difficulty: Medium
// Time: O(n^2) worst-case | Space: O(1)
// For each start, expand and track LCM.

import "fmt"

func main() {
	fmt.Println(subarrayLCM([]int{3, 6, 2, 7, 1}, 6)) // 4
	fmt.Println(subarrayLCM([]int{3}, 2))              // 0
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func subarrayLCM(nums []int, k int) int {
	ans := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(nums); i++ {
		cur := 1
		for j := i; j < len(nums); j++ {
			cur = lcm(cur, nums[j])
			if cur == k {
				ans++
			}
			if cur > k {
				break
			}
		}
	}
	return ans
}
```
