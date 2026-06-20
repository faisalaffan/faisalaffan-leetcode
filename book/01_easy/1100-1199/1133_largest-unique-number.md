# 1133 — Largest Unique Number

## Deskripsi

**Soal:** [1133. Largest Unique Number](https://leetcode.com/problems/largest-unique-number/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1133: Largest Unique Number
// https://leetcode.com/problems/largest-unique-number/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(largestUniqueNumber([]int{5, 7, 3, 9, 4, 9, 8, 3, 1})) // 8
	fmt.Println(largestUniqueNumber([]int{9, 9, 8, 8}))                // -1
}

// LeetCode submission: largestUniqueNumber
func largestUniqueNumber(nums []int) int {
  // Membuat map untuk pencarian O(1): key → value
	count := make(map[int]int)
	for _, v := range nums {
		count[v]++
	}
	ans := -1
	for k, v := range count {
		if v == 1 && k > ans {
			ans = k
		}
	}
	return ans
}
```
