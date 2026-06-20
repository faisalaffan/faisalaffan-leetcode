# 3285 — Find Indices Of Stable Mountains

## Deskripsi

**Soal:** [3285. Find Indices Of Stable Mountains](https://leetcode.com/problems/find-indices-of-stable-mountains/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3285: Find Indices of Stable Mountains
// https://leetcode.com/problems/find-indices-of-stable-mountains/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindIndicesOfStableMountains([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(FindIndicesOfStableMountains([]int{10, 1, 10, 1, 10}, 3))
}

// FindIndicesOfStableMountains returns indices of stable mountains (where the previous mountain's height > threshold).
// Time: O(n). Space: O(n).
func FindIndicesOfStableMountains(height []int, threshold int) []int {
	result := []int{}
	for i := 1; i < len(height); i++ {
		if height[i-1] > threshold {
			result = append(result, i)
		}
	}
	return result
}
```
