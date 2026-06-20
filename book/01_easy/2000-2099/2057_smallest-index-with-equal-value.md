# 2057 — Smallest Index With Equal Value

## Deskripsi

**Soal:** [2057. Smallest Index With Equal Value](https://leetcode.com/problems/smallest-index-with-equal-value/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2057: Smallest Index With Equal Value
// https://leetcode.com/problems/smallest-index-with-equal-value/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SmallestIndexWithEqualValue([]int{0, 1, 2}))          // 0
	fmt.Println(SmallestIndexWithEqualValue([]int{4, 3, 2, 1}))       // 2
	fmt.Println(SmallestIndexWithEqualValue([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})) // -1
}

// Time: O(n), Space: O(1)
func SmallestIndexWithEqualValue(nums []int) int {
	for i, v := range nums {
		if i%10 == v {
			return i
		}
	}
	return -1
}
```
