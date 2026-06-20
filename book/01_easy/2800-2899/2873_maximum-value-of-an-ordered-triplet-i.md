# 2873 — Maximum Value Of An Ordered Triplet I

## Deskripsi

**Soal:** [2873. Maximum Value Of An Ordered Triplet I](https://leetcode.com/problems/maximum-value-of-an-ordered-triplet-i/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2873: Maximum Value of an Ordered Triplet I
// https://leetcode.com/problems/maximum-value-of-an-ordered-triplet-i/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(MaximumValueOfAnOrderedTripletI([]int{12, 6, 1, 2, 7}))
	fmt.Println(MaximumValueOfAnOrderedTripletI([]int{1, 10, 3, 2, 5}))
}

func MaximumValueOfAnOrderedTripletI(nums []int) int64 {
	n := len(nums)
	if n < 3 {
		return 0
	}

	maxNum := int64(nums[0])
	maxDiff := int64(nums[0] - nums[1])
	ans := int64(0)

	for i := 2; i < n; i++ {
		val := maxDiff * int64(nums[i])
		if val > ans {
			ans = val
		}
		if int64(nums[i-1]) > maxNum {
			maxNum = int64(nums[i-1])
		}
		if maxNum-int64(nums[i]) > maxDiff {
			maxDiff = maxNum - int64(nums[i])
		}
	}

	return ans
}
```
