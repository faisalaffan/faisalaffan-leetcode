# 0553 — Optimal Division

## Deskripsi

**Soal:** [0553. Optimal Division](https://leetcode.com/problems/optimal-division/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #553: Optimal Division
// https://leetcode.com/problems/optimal-division/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(OptimalDivision([]int{1000, 100, 10, 2}))
	fmt.Println(OptimalDivision([]int{2, 3, 4}))
	fmt.Println(OptimalDivision([]int{2}))
}

func OptimalDivision(nums []int) string {
	n := len(nums)
  // Edge case: input kosong
	if n == 0 {
		return ""
	}
	if n == 1 {
		return strconv.Itoa(nums[0])
	}
	if n == 2 {
		return strconv.Itoa(nums[0]) + "/" + strconv.Itoa(nums[1])
	}

	result := strconv.Itoa(nums[0]) + "/(" + strconv.Itoa(nums[1])
	for i := 2; i < n; i++ {
		result += "/" + strconv.Itoa(nums[i])
	}
	result += ")"

	return result
}
```
