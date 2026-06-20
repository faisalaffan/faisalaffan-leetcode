# 3028 — Ant On The Boundary

## Deskripsi

**Soal:** [3028. Ant On The Boundary](https://leetcode.com/problems/ant-on-the-boundary/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3028: Ant on the Boundary
// https://leetcode.com/problems/ant-on-the-boundary/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: returnToBoundaryCount
	fmt.Println(AntOnTheBoundary([]int{2, 3, -5})) // 1
	fmt.Println(AntOnTheBoundary([]int{3, 2, -3, -2})) // 1
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: returnToBoundaryCount
func AntOnTheBoundary(nums []int) int {
	pos := 0
	count := 0
	for _, v := range nums {
		pos += v
		if pos == 0 {
			count++
		}
	}
	return count
}
```
