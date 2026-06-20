# 2206 — Divide Array Into Equal Pairs

## Deskripsi

**Soal:** [2206. Divide Array Into Equal Pairs](https://leetcode.com/problems/divide-array-into-equal-pairs/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2206: Divide Array Into Equal Pairs
// https://leetcode.com/problems/divide-array-into-equal-pairs/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(DivideArrayIntoEqualPairs([]int{3, 2, 3, 2, 2, 2})) // true
	fmt.Println(DivideArrayIntoEqualPairs([]int{1, 2, 3, 4}))       // false
}

// Time: O(n), Space: O(n)
func DivideArrayIntoEqualPairs(nums []int) bool {
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	for _, c := range freq {
		if c%2 != 0 {
			return false
		}
	}
	return true
}
```
