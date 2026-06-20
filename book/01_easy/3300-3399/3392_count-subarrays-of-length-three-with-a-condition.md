# 3392 — Count Subarrays Of Length Three With A Condition

## Deskripsi

**Soal:** [3392. Count Subarrays Of Length Three With A Condition](https://leetcode.com/problems/count-subarrays-of-length-three-with-a-condition/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3392: Count Subarrays of Length Three With a Condition
// https://leetcode.com/problems/count-subarrays-of-length-three-with-a-condition/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountSubarraysOfLengthThreeWithACondition([]int{1, 2, 1, 2, 1}))
	fmt.Println(CountSubarraysOfLengthThreeWithACondition([]int{1, 3, 5, 7, 9}))
}

// CountSubarraysOfLengthThreeWithACondition counts subarrays of length 3 where the sum of first and last equals the middle.
// Time: O(n). Space: O(1).
func CountSubarraysOfLengthThreeWithACondition(nums []int) int {
	count := 0
	for i := 0; i <= len(nums)-3; i++ {
		if nums[i]+nums[i+2] == nums[i+1] {
			count++
		}
	}
	return count
}
```
