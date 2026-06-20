# 0905 — Sort Array By Parity

## Deskripsi

**Soal:** [0905. Sort Array By Parity](https://leetcode.com/problems/sort-array-by-parity/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #905: Sort Array By Parity
// https://leetcode.com/problems/sort-array-by-parity/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4})) // [2,4,3,1] or [4,2,1,3] etc.
	fmt.Println(sortArrayByParity([]int{0}))           // [0]
	fmt.Println(sortArrayByParity([]int{1, 3, 5}))     // [1,3,5]
}

// sortArrayByParity moves all even numbers to the front, odd to the back.
// Time: O(n). Space: O(1).
func sortArrayByParity(nums []int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		if nums[l]%2 == 0 {
			l++
		} else {
			nums[l], nums[r] = nums[r], nums[l]
			r--
		}
	}
	return nums
}
```
