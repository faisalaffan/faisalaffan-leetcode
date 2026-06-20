# 1064 — Fixed Point

## Deskripsi

**Soal:** [1064. Fixed Point](https://leetcode.com/problems/fixed-point/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** Dynamic Programming (DP)

## Solusi Go

```go
package main

// LeetCode #1064: Fixed Point
// https://leetcode.com/problems/fixed-point/
// Difficulty: Easy [Paid]
// Time: O(log n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(fixedPoint([]int{-10, -5, 0, 3, 7}))  // 3
	fmt.Println(fixedPoint([]int{0, 2, 5, 8, 17}))    // 0
	fmt.Println(fixedPoint([]int{-10, -5, 3, 4, 7, 9})) // -1
}

// LeetCode submission: fixedPoint
func fixedPoint(arr []int) int {
	lo, hi := 0, len(arr)-1
	for lo < hi {
		mid := (lo + hi) >> 1
		if arr[mid] >= mid {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	if arr[lo] == lo {
		return lo
	}
	return -1
}
```
