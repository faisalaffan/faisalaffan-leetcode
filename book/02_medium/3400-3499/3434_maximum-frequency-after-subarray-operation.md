# 3434 — Maximum Frequency After Subarray Operation

## Deskripsi

**Soal:** [3434. Maximum Frequency After Subarray Operation](https://leetcode.com/problems/maximum-frequency-after-subarray-operation/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(50*n) Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func maxFrequency(nums []int, k int) int`

## Solusi Go

```go
package main

// LeetCode #3434: Maximum Frequency After Subarray Operation
// https://leetcode.com/problems/maximum-frequency-after-subarray-operation/
// Difficulty: Medium
// Time: O(50*n) Space: O(1)

import "fmt"

func maxFrequency(nums []int, k int) int {
	cntK := 0
	for _, v := range nums {
		if v == k {
			cntK++
		}
	}
	ans := cntK
	for target := 1; target <= 50; target++ {
		if target == k {
			continue
		}
		cur := 0
		for _, x := range nums {
			if x == target {
				cur++
			} else if x == k {
				cur--
			}
			if cur < 0 {
				cur = 0
			}
			if cntK+cur > ans {
				ans = cntK + cur
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(maxFrequency([]int{10, 2, 3, 4, 5, 5, 4, 3, 2, 2}, 10)) // 4
	fmt.Println(maxFrequency([]int{1, 2, 3, 4, 5}, 1)) // 2
}
```
