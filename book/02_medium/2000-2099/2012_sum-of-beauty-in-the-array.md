# 2012 — Sum Of Beauty In The Array

## Deskripsi

**Soal:** [2012. Sum Of Beauty In The Array](https://leetcode.com/problems/sum-of-beauty-in-the-array/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2012: Sum of Beauty in the Array
// https://leetcode.com/problems/sum-of-beauty-in-the-array/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(SumOfBeautyInTheArray([]int{1, 2, 3}))
	fmt.Println(SumOfBeautyInTheArray([]int{2, 4, 6, 4}))
	fmt.Println(SumOfBeautyInTheArray([]int{3, 2, 1}))
}

// Time: O(n), Space: O(n)
func SumOfBeautyInTheArray(nums []int) int {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	prefixMax := make([]int, n)
  // Membuat slice untuk menyimpan hasil
	suffixMin := make([]int, n)

	prefixMax[0] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > prefixMax[i-1] {
			prefixMax[i] = nums[i]
		} else {
			prefixMax[i] = prefixMax[i-1]
		}
	}

	suffixMin[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] < suffixMin[i+1] {
			suffixMin[i] = nums[i]
		} else {
			suffixMin[i] = suffixMin[i+1]
		}
	}

	ans := 0
	for i := 1; i < n-1; i++ {
		if nums[i] > prefixMax[i-1] && nums[i] < suffixMin[i+1] {
			ans += 2
		} else if nums[i] > nums[i-1] && nums[i] < nums[i+1] {
			ans++
		}
	}

	return ans
}
```
