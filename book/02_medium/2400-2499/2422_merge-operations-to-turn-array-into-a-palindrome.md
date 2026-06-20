# 2422 — Merge Operations To Turn Array Into A Palindrome

## Deskripsi

**Soal:** [2422. Merge Operations To Turn Array Into A Palindrome](https://leetcode.com/problems/merge-operations-to-turn-array-into-a-palindrome/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** Two Pointer (penunjuk kiri & kanan), Two Pointer (penunjuk kiri & kanan)

## Solusi Go

```go
package main

// LeetCode #2422: Merge Operations to Turn Array Into a Palindrome
// https://leetcode.com/problems/merge-operations-to-turn-array-into-a-palindrome/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Two pointers, merge smaller side towards larger.

import "fmt"

func main() {
	fmt.Println(minMerges([]int{1, 4, 1, 3}))        // 1 (merge 4+1 to 5, then [1,5,3])
	fmt.Println(minMerges([]int{1, 2, 3, 4, 5, 1})) // 2
}

func minMerges(nums []int) int {
	ops := 0
	i, j := 0, len(nums)-1
	left, right := nums[i], nums[j]

	for i < j {
		if left < right {
			i++
			left += nums[i]
			ops++
		} else if left > right {
			j--
			right += nums[j]
			ops++
		} else {
			i++
			j--
			if i < j {
				left, right = nums[i], nums[j]
			}
		}
	}
	return ops
}
```
