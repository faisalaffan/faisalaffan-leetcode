# 2951 — Find The Peaks

## Deskripsi

**Soal:** [2951. Find The Peaks](https://leetcode.com/problems/find-the-peaks/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1) excluding output

**Algoritma:** Dynamic Programming (DP)

## Solusi Go

```go
package main

// LeetCode #2951: Find the Peaks
// https://leetcode.com/problems/find-the-peaks/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: findPeaks
	fmt.Println(FindThePeaks([]int{2, 4, 4}))    // []
	fmt.Println(FindThePeaks([]int{1, 4, 3, 8, 5})) // [1, 3]
}

// Time: O(n) | Space: O(1) excluding output
// LeetCode submission name: findPeaks
func FindThePeaks(mountain []int) []int {
	result := []int{}
	for i := 1; i < len(mountain)-1; i++ {
		if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
			result = append(result, i)
		}
	}
	return result
}
```
