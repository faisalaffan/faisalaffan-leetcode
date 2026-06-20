# 0665 — Non Decreasing Array

## Deskripsi

**Soal:** [0665. Non Decreasing Array](https://leetcode.com/problems/non-decreasing-array/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #665: Non-decreasing Array
// https://leetcode.com/problems/non-decreasing-array/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(checkPossibility([]int{4, 2, 3}))
	fmt.Println(checkPossibility([]int{4, 2, 1}))
	fmt.Println(checkPossibility([]int{3, 4, 2, 3}))
}

func checkPossibility(nums []int) bool {
	modified := false

  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if modified {
				return false
			}

			if i == 0 || nums[i-1] <= nums[i+1] {
				nums[i] = nums[i+1]
			} else {
				nums[i+1] = nums[i]
			}

			modified = true
		}
	}

	return true
}
```
