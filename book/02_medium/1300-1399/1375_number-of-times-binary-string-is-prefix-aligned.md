# 1375 — Number Of Times Binary String Is Prefix Aligned

## Deskripsi

**Soal:** [1375. Number Of Times Binary String Is Prefix Aligned](https://leetcode.com/problems/number-of-times-binary-string-is-prefix-aligned/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n) where n = length of flips  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1375: Number of Times Binary String Is Prefix-Aligned
// https://leetcode.com/problems/number-of-times-binary-string-is-prefix-aligned/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(numTimesAllBlue([]int{3, 2, 4, 1, 5})) // 2

	// Test case 2
	fmt.Println(numTimesAllBlue([]int{4, 1, 2, 3})) // 1

	// Test case 3
	fmt.Println(numTimesAllBlue([]int{2, 1, 3})) // 1
}

// Time: O(n) where n = length of flips
// Space: O(1)
func numTimesAllBlue(flips []int) int {
	count := 0
	maxFlip := 0

	for i, f := range flips {
		if f > maxFlip {
			maxFlip = f
		}
		if maxFlip == i+1 {
			count++
		}
	}

	return count
}
```
