# 3371 — Identify The Largest Outlier In An Array

## Deskripsi

**Soal:** [3371. Identify The Largest Outlier In An Array](https://leetcode.com/problems/identify-the-largest-outlier-in-an-array/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n) Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3371: Identify the Largest Outlier in an Array
// https://leetcode.com/problems/identify-the-largest-outlier-in-an-array/
// Difficulty: Medium
// Time: O(n) Space: O(n)

import "fmt"

func main() {
	fmt.Println(getLargestOutlier([]int{2, 3, 5, 10}))   // 10
	fmt.Println(getLargestOutlier([]int{-2, -1, -3, -6, 4})) // 4
	fmt.Println(getLargestOutlier([]int{1, 1, 1, 1, 1, 5, 5}))  // 5
}

func getLargestOutlier(nums []int) int {
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[int]int)
	total := 0
	for _, v := range nums {
		freq[v]++
		total += v
	}

	ans := -(1 << 60)
	for _, v := range nums {
		outlier := total - 2*v
		if outlier == v && freq[v] < 2 {
			continue
		}
		if _, ok := freq[outlier]; ok && outlier > ans {
			ans = outlier
		}
	}
	return ans
}
```
