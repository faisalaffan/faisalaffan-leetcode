# 3741 — Minimum Distance Between Three Equal Elements Ii

## Deskripsi

**Soal:** [3741. Minimum Distance Between Three Equal Elements Ii](https://leetcode.com/problems/minimum-distance-between-three-equal-elements-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func minimumDistanceBetweenThreeEqualElementsIi(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #3741: Minimum Distance Between Three Equal Elements II
// https://leetcode.com/problems/minimum-distance-between-three-equal-elements-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func minimumDistanceBetweenThreeEqualElementsIi(nums []int) int {
	ans := -1
  // Membuat map untuk pencarian O(1): key → value
	prev1 := make(map[int]int)
  // Membuat map untuk pencarian O(1): key → value
	prev2 := make(map[int]int)

	// prev1[v] = last index where v appeared
	// prev2[v] = second-last index where v appeared

	for i, v := range nums {
		if p2, ok := prev2[v]; ok {
			dist := 2 * (i - p2)
			if ans == -1 || dist < ans {
				ans = dist
			}
		}
		// Shift: prev2 gets prev1, prev1 gets current
		prev2[v] = prev1[v]
		prev1[v] = i
	}

	return ans
}

func main() {
	fmt.Println(minimumDistanceBetweenThreeEqualElementsIi([]int{1, 3, 1, 1, 2, 1}))
	fmt.Println(minimumDistanceBetweenThreeEqualElementsIi([]int{1, 2, 3, 4}))
	fmt.Println(minimumDistanceBetweenThreeEqualElementsIi([]int{1, 1, 1}))
}
```
