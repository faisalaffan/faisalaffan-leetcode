# 3736 — Minimum Moves To Equal Array Elements Iii

## Deskripsi

**Soal:** [3736. Minimum Moves To Equal Array Elements Iii](https://leetcode.com/problems/minimum-moves-to-equal-array-elements-iii/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3736: Minimum Moves to Equal Array Elements III
// https://leetcode.com/problems/minimum-moves-to-equal-array-elements-iii/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumMovesToEqualArrayElementsIii([]int{2, 1, 3}))
	fmt.Println(MinimumMovesToEqualArrayElementsIii([]int{4, 4, 5}))
}

// Time: O(n)
// Space: O(1)
func MinimumMovesToEqualArrayElementsIii(nums []int) int {
	maxVal := nums[0]
	sum := 0
	for _, v := range nums {
		sum += v
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal*len(nums) - sum
}
```
