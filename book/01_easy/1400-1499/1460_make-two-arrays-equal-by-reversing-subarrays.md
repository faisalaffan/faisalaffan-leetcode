# 1460 — Make Two Arrays Equal By Reversing Subarrays

## Deskripsi

**Soal:** [1460. Make Two Arrays Equal By Reversing Subarrays](https://leetcode.com/problems/make-two-arrays-equal-by-reversing-subarrays/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func canBeEqual(target []int, arr []int) bool`

## Solusi Go

```go
package main

// LeetCode #1460: Make Two Arrays Equal by Reversing Subarrays
// https://leetcode.com/problems/make-two-arrays-equal-by-reversing-subarrays/
// Difficulty: Easy
//
// LeetCode submission: func canBeEqual(target []int, arr []int) bool

import "fmt"

func main() {
	fmt.Println(MakeTwoArraysEqualByReversingSubarrays([]int{1, 2, 3, 4}, []int{2, 4, 1, 3})) // true
	fmt.Println(MakeTwoArraysEqualByReversingSubarrays([]int{7}, []int{7}))                    // true
	fmt.Println(MakeTwoArraysEqualByReversingSubarrays([]int{3, 7, 9}, []int{3, 7, 11}))       // false
}

// Time: O(n), Space: O(n)
func MakeTwoArraysEqualByReversingSubarrays(target []int, arr []int) bool {
	if len(target) != len(arr) {
		return false
	}
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[int]int, len(target))
	for _, v := range target {
		freq[v]++
	}
	for _, v := range arr {
		freq[v]--
		if freq[v] < 0 {
			return false
		}
	}
	return true
}
```
