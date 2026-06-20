# 3678 — Smallest Absent Positive Greater Than Average

## Deskripsi

**Soal:** [3678. Smallest Absent Positive Greater Than Average](https://leetcode.com/problems/smallest-absent-positive-greater-than-average/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n + m) where m is the range of candidate values  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3678: Smallest Absent Positive Greater Than Average
// https://leetcode.com/problems/smallest-absent-positive-greater-than-average/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SmallestAbsentPositiveGreaterThanAverage([]int{3, 5}))
	fmt.Println(SmallestAbsentPositiveGreaterThanAverage([]int{-1, 1, 2}))
	fmt.Println(SmallestAbsentPositiveGreaterThanAverage([]int{4, -1}))
}

// Time: O(n + m) where m is the range of candidate values
// Space: O(n)
func SmallestAbsentPositiveGreaterThanAverage(nums []int) int {
  // Membuat map untuk pencarian O(1): key → value
	has := make(map[int]bool)
	sum := 0
	for _, x := range nums {
		has[x] = true
		sum += x
	}

	avg := float64(sum) / float64(len(nums))
	ans := 1
	if int(avg)+1 > ans {
		ans = int(avg) + 1
	}

	for has[ans] {
		ans++
	}
	return ans
}
```
