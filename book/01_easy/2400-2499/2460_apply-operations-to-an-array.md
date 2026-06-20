# 2460 — Apply Operations To An Array

## Deskripsi

**Soal:** [2460. Apply Operations To An Array](https://leetcode.com/problems/apply-operations-to-an-array/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2460: Apply Operations to an Array
// https://leetcode.com/problems/apply-operations-to-an-array/
// Difficulty: Easy
// Time O(n) | Space O(n)

import "fmt"

func main() {
	fmt.Println(ApplyOperationsToAnArray([]int{1, 2, 2, 1, 1, 0})) // [1,4,2,0,0,0]
	fmt.Println(ApplyOperationsToAnArray([]int{0, 1}))              // [1,0]
}

func ApplyOperationsToAnArray(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}
	}

  // Membuat slice untuk menyimpan hasil
	res := make([]int, n)
	idx := 0
	for _, v := range nums {
		if v != 0 {
			res[idx] = v
			idx++
		}
	}
	return res
}
```
