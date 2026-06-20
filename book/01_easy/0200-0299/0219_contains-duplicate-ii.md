# 0219 — Contains Duplicate Ii

## Deskripsi

**Soal:** [0219. Contains Duplicate Ii](https://leetcode.com/problems/contains-duplicate-ii/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func ContainsNearbyDuplicate(nums []int, k int) bool`

## Solusi Go

```go
package main

// LeetCode #219: Contains Duplicate II
// https://leetcode.com/problems/contains-duplicate-ii/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(n)
func ContainsNearbyDuplicate(nums []int, k int) bool {
  // Membuat map untuk pencarian O(1): key → value
	seen := make(map[int]int, len(nums))
	for i, n := range nums {
		if j, ok := seen[n]; ok && i-j <= k {
			return true
		}
		seen[n] = i
	}
	return false
}

func main() {
	fmt.Println(ContainsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println(ContainsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println(ContainsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}
```
