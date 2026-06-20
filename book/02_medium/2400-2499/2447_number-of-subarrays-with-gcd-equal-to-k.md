# 2447 — Number Of Subarrays With Gcd Equal To K

## Deskripsi

**Soal:** [2447. Number Of Subarrays With Gcd Equal To K](https://leetcode.com/problems/number-of-subarrays-with-gcd-equal-to-k/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2) worst-case  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2447: Number of Subarrays With GCD Equal to K
// https://leetcode.com/problems/number-of-subarrays-with-gcd-equal-to-k/
// Difficulty: Medium
// Time: O(n^2) worst-case | Space: O(1)
// For each i, expand j and track GCD. Count when GCD == k.

import "fmt"

func main() {
	fmt.Println(subarrayGCD([]int{9, 3, 1, 2, 6, 3}, 3)) // 4
	fmt.Println(subarrayGCD([]int{4}, 7))                  // 0
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func subarrayGCD(nums []int, k int) int {
	ans := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(nums); i++ {
		cur := 0
		for j := i; j < len(nums); j++ {
			cur = gcd(cur, nums[j])
			if cur == k {
				ans++
			}
			if cur < k {
				break // GCD only decreases, can't reach k again
			}
		}
	}
	return ans
}
```
