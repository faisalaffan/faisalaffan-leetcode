# 3159 — Find Occurrences Of An Element In An Array

## Deskripsi

**Soal:** [3159. Find Occurrences Of An Element In An Array](https://leetcode.com/problems/find-occurrences-of-an-element-in-an-array/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n + q)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func occurrencesOfElement(nums []int, queries []int, x int) []int`

## Solusi Go

```go
package main

// LeetCode #3159: Find Occurrences of an Element in an Array
// https://leetcode.com/problems/find-occurrences-of-an-element-in-an-array/
// Difficulty: Medium
// Time: O(n + q) | Space: O(n)

import "fmt"

func occurrencesOfElement(nums []int, queries []int, x int) []int {
  // Membuat slice untuk menyimpan hasil
	pos := make([]int, 0)
	for i, v := range nums {
		if v == x {
			pos = append(pos, i)
		}
	}

  // Membuat slice untuk menyimpan hasil
	ans := make([]int, len(queries))
	for i, q := range queries {
		if q-1 < len(pos) {
			ans[i] = pos[q-1]
		} else {
			ans[i] = -1
		}
	}
	return ans
}

func main() {
	fmt.Println(occurrencesOfElement([]int{1, 3, 1, 7}, []int{1, 3, 2, 4}, 1)) // Expected: [0, -1, 2, -1]
	fmt.Println(occurrencesOfElement([]int{1, 2, 3}, []int{10}, 5))             // Expected: [-1]
}
```
