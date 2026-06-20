# 2341 — Maximum Number Of Pairs In Array

## Deskripsi

**Soal:** [2341. Maximum Number Of Pairs In Array](https://leetcode.com/problems/maximum-number-of-pairs-in-array/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2341: Maximum Number of Pairs in Array
// https://leetcode.com/problems/maximum-number-of-pairs-in-array/
// Difficulty: Easy
// Time O(n) | Space O(n)

import "fmt"

func main() {
	fmt.Println(MaximumNumberOfPairsInArray([]int{1, 3, 2, 1, 3, 2, 2})) // [3,1]
	fmt.Println(MaximumNumberOfPairsInArray([]int{1, 1}))                  // [1,0]
}

func MaximumNumberOfPairsInArray(nums []int) []int {
	freq := map[int]int{}
	for _, n := range nums {
		freq[n]++
	}
	pairs, leftovers := 0, 0
	for _, c := range freq {
		pairs += c / 2
		leftovers += c % 2
	}
	return []int{pairs, leftovers}
}
```
