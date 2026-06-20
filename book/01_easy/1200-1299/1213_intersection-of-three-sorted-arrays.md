# 1213 — Intersection Of Three Sorted Arrays

## Deskripsi

**Soal:** [1213. Intersection Of Three Sorted Arrays](https://leetcode.com/problems/intersection-of-three-sorted-arrays/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1) excluding output

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1213: Intersection of Three Sorted Arrays
// https://leetcode.com/problems/intersection-of-three-sorted-arrays/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(1) excluding output

import "fmt"

func main() {
	fmt.Println(arraysIntersection([]int{1, 2, 3, 4, 5}, []int{1, 2, 5, 7, 9}, []int{1, 3, 4, 5, 8}))
	// [1,5]
	fmt.Println(arraysIntersection([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}))
	// []
}

// LeetCode submission: arraysIntersection
func arraysIntersection(arr1, arr2, arr3 []int) []int {
	var ans []int
	i, j, k := 0, 0, 0
	for i < len(arr1) && j < len(arr2) && k < len(arr3) {
		if arr1[i] == arr2[j] && arr2[j] == arr3[k] {
			ans = append(ans, arr1[i])
			i++; j++; k++
		} else if arr1[i] < arr2[j] {
			i++
		} else if arr2[j] < arr3[k] {
			j++
		} else {
			k++
		}
	}
	return ans
}
```
