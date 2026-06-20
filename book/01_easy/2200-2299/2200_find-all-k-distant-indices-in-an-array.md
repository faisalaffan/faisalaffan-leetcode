# 2200 — Find All K Distant Indices In An Array

## Deskripsi

**Soal:** [2200. Find All K Distant Indices In An Array](https://leetcode.com/problems/find-all-k-distant-indices-in-an-array/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2200: Find All K-Distant Indices in an Array
// https://leetcode.com/problems/find-all-k-distant-indices-in-an-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindAllKDistantIndicesInAnArray([]int{3, 4, 9, 1, 3, 9, 5}, 9, 1)) // [1 2 3 4 5 6]
	fmt.Println(FindAllKDistantIndicesInAnArray([]int{2, 2, 2, 2, 2}, 2, 2))       // [0 1 2 3 4]
}

// Time: O(n), Space: O(n)
func FindAllKDistantIndicesInAnArray(nums []int, key int, k int) []int {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	marked := make([]bool, n)

	farthest := -1
	for i, v := range nums {
		if v == key {
			start := i - k
			if start < 0 {
				start = 0
			}
			if start <= farthest {
				start = farthest + 1
			}
			end := i + k
			if end >= n {
				end = n - 1
			}
			for j := start; j <= end; j++ {
				marked[j] = true
			}
			farthest = end
		}
	}

	var result []int
	for i, m := range marked {
		if m {
			result = append(result, i)
		}
	}
	return result
}
```
