# 1891 — Cutting Ribbons

## Deskripsi

**Soal:** [1891. Cutting Ribbons](https://leetcode.com/problems/cutting-ribbons/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log maxLen), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1891: Cutting Ribbons
// https://leetcode.com/problems/cutting-ribbons/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	fmt.Println(MaxLength([]int{9, 7, 5}, 3))
	fmt.Println(MaxLength([]int{7, 5, 9}, 4))
	fmt.Println(MaxLength([]int{5, 7, 9}, 22))
}

// Time: O(n log maxLen), Space: O(1)
func MaxLength(ribbons []int, k int) int {
	left, right := 1, 0
	for _, r := range ribbons {
		if r > right {
			right = r
		}
	}

	ans := 0
	for left <= right {
		mid := left + (right-left)/2
		count := 0
		for _, r := range ribbons {
			count += r / mid
		}
		if count >= k {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}
```
