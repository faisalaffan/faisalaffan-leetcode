# 1848 — Minimum Distance To The Target Element

## Deskripsi

**Soal:** [1848. Minimum Distance To The Target Element](https://leetcode.com/problems/minimum-distance-to-the-target-element/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func GetMinDistance(nums []int, target int, start int) int`

## Solusi Go

```go
package main

// LeetCode #1848: Minimum Distance to the Target Element
// https://leetcode.com/problems/minimum-distance-to-the-target-element/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func GetMinDistance(nums []int, target int, start int) int {
	minDist := len(nums)
	for i, num := range nums {
		if num == target {
			dist := start - i
			if dist < 0 {
				dist = -dist
			}
			if dist < minDist {
				minDist = dist
			}
		}
	}
	return minDist
}

func main() {
	fmt.Println(GetMinDistance([]int{1, 2, 3, 4, 5}, 5, 3))
	fmt.Println(GetMinDistance([]int{1}, 1, 0))
	fmt.Println(GetMinDistance([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1, 0))
}
```
