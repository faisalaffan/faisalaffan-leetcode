# 2229 — Check If An Array Is Consecutive

## Deskripsi

**Soal:** [2229. Check If An Array Is Consecutive](https://leetcode.com/problems/check-if-an-array-is-consecutive/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2229: Check if an Array Is Consecutive
// https://leetcode.com/problems/check-if-an-array-is-consecutive/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(CheckIfAnArrayIsConsecutive([]int{1, 3, 4, 2})) // true
	fmt.Println(CheckIfAnArrayIsConsecutive([]int{1, 3, 5}))    // false
	fmt.Println(CheckIfAnArrayIsConsecutive([]int{1, 4}))       // false
}

// Time: O(n), Space: O(n)
func CheckIfAnArrayIsConsecutive(nums []int) bool {
  // Edge case: input kosong
	if len(nums) == 0 {
		return false
	}

  // Membuat map untuk pencarian O(1): key → value
	set := make(map[int]bool)
	min, max := nums[0], nums[0]

	for _, v := range nums {
		set[v] = true
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	if max-min+1 != len(nums) {
		return false
	}

	for i := min; i <= max; i++ {
		if !set[i] {
			return false
		}
	}
	return true
}
```
