# 1995 — Count Special Quadruplets

## Deskripsi

**Soal:** [1995. Count Special Quadruplets](https://leetcode.com/problems/count-special-quadruplets/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n^3), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1995: Count Special Quadruplets
// https://leetcode.com/problems/count-special-quadruplets/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountSpecialQuadruplets([]int{1, 2, 3, 6}))   // 1
	fmt.Println(CountSpecialQuadruplets([]int{3, 3, 6, 4, 5})) // 0
	fmt.Println(CountSpecialQuadruplets([]int{1, 1, 1, 3, 5})) // 4
}

// Time: O(n^3), Space: O(1)
func CountSpecialQuadruplets(nums []int) int {
	n := len(nums)
	count := 0
	for a := 0; a < n; a++ {
		for b := a + 1; b < n; b++ {
			for c := b + 1; c < n; c++ {
				for d := c + 1; d < n; d++ {
					if nums[a]+nums[b]+nums[c] == nums[d] {
						count++
					}
				}
			}
		}
	}
	return count
}
```
