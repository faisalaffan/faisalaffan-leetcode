# 3326 — Minimum Division Operations To Make Array Non Decreasing

## Deskripsi

**Soal:** [3326. Minimum Division Operations To Make Array Non Decreasing](https://leetcode.com/problems/minimum-division-operations-to-make-array-non-decreasing/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(M log log M + n) Space: O(M) where M = 1e6  
**Kompleksitas Ruang:** O(M) where M = 1e6

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3326: Minimum Division Operations to Make Array Non Decreasing
// https://leetcode.com/problems/minimum-division-operations-to-make-array-non-decreasing/
// Difficulty: Medium
// Time: O(M log log M + n) Space: O(M) where M = 1e6

import "fmt"

func main() {
	fmt.Println(minOperations([]int{25, 7}))       // 1
	fmt.Println(minOperations([]int{7, 7, 6}))     // -1
	fmt.Println(minOperations([]int{1, 1, 1, 1}))  // 0
}

const mx = 1000001

var lpf [mx]int

func init() {
	for i := 2; i < mx; i++ {
		if lpf[i] == 0 {
			for j := i; j < mx; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}
}

func minOperations(nums []int) int {
	ans := 0
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] > nums[i+1] {
			if lpf[nums[i]] > nums[i+1] {
				return -1
			}
			nums[i] = lpf[nums[i]]
			ans++
		}
	}
	return ans
}
```
