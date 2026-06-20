# 0350 — Intersection Of Two Arrays Ii

## Deskripsi

**Soal:** [0350. Intersection Of Two Arrays Ii](https://leetcode.com/problems/intersection-of-two-arrays-ii/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n+m), Space: O(min(n,m))  
**Kompleksitas Ruang:** O(min(n,m))

**Algoritma:** —

**Fungsi Solusi:** `func IntersectionOfTwoArraysIi(nums1, nums2 []int) []int`

## Solusi Go

```go
package main

// LeetCode #350: Intersection of Two Arrays II
// https://leetcode.com/problems/intersection-of-two-arrays-ii/
// Difficulty: Easy

import "fmt"

// Time: O(n+m), Space: O(min(n,m))
func IntersectionOfTwoArraysIi(nums1, nums2 []int) []int {
  // Membuat map untuk pencarian O(1): key → value
	count := make(map[int]int)
	for _, v := range nums1 {
		count[v]++
	}
	var result []int
	for _, v := range nums2 {
		if count[v] > 0 {
			result = append(result, v)
			count[v]--
		}
	}
	return result
}

func main() {
	fmt.Println(IntersectionOfTwoArraysIi([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(IntersectionOfTwoArraysIi([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
```
