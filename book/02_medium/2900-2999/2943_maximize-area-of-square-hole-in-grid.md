# 2943 — Maximize Area Of Square Hole In Grid

## Deskripsi

**Soal:** [2943. Maximize Area Of Square Hole In Grid](https://leetcode.com/problems/maximize-area-of-square-hole-in-grid/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(h log h + v log v)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2943: Maximize Area of Square Hole in Grid
// https://leetcode.com/problems/maximize-area-of-square-hole-in-grid/
// Difficulty: Medium
// Time: O(h log h + v log v) | Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximizeSquareArea(3, 4, []int{2}, []int{2}))
	fmt.Println(maximizeSquareArea(2, 2, []int{1}, []int{1}))
}

func maximizeSquareArea(m int, n int, hBars []int, vBars []int) int {
	calc := func(nums []int, limit int) int {
		nums = append(nums, 1)
		nums = append(nums, limit)
		sort.Ints(nums)
		ans, cnt := 1, 1
		for i := 1; i < len(nums); i++ {
			if nums[i] == nums[i-1]+1 {
				cnt++
				if cnt > ans {
					ans = cnt
				}
			} else {
				cnt = 1
			}
		}
		return ans
	}
	x := calc(hBars, m)
	y := calc(vBars, n)
	if x > y {
		x = y
	}
	return x * x
}
```
