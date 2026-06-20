# 3676 — Count Bowl Subarrays

## Deskripsi

**Soal:** [3676. Count Bowl Subarrays](https://leetcode.com/problems/count-bowl-subarrays/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Stack (tumpukan LIFO)

**Fungsi Solusi:** `func countBowlSubarrays(nums []int) int64`

## Solusi Go

```go
package main

// LeetCode #3676: Count Bowl Subarrays
// https://leetcode.com/problems/count-bowl-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func countBowlSubarrays(nums []int) int64 {
	var ans int64 = 0
	var stack []int

	for _, num := range nums {
		for len(stack) > 0 && stack[len(stack)-1] < num {
			if len(stack) >= 2 {
				ans++
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}

	return ans
}

func main() {
	fmt.Println(countBowlSubarrays([]int{1, 3, 5, 4, 2}))
	fmt.Println(countBowlSubarrays([]int{3, 1, 2, 4}))
	fmt.Println(countBowlSubarrays([]int{1, 2, 3, 4}))
}
```
