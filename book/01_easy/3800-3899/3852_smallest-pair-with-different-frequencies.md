# 3852 — Smallest Pair With Different Frequencies

## Deskripsi

**Soal:** [3852. Smallest Pair With Different Frequencies](https://leetcode.com/problems/smallest-pair-with-different-frequencies/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3852: Smallest Pair With Different Frequencies
// https://leetcode.com/problems/smallest-pair-with-different-frequencies/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SmallestPairWithDifferentFrequencies([]int{1, 1, 2, 2, 3, 4}))
	fmt.Println(SmallestPairWithDifferentFrequencies([]int{1, 5}))
	fmt.Println(SmallestPairWithDifferentFrequencies([]int{7}))
}

// Time: O(n)
// Space: O(n)
func SmallestPairWithDifferentFrequencies(nums []int) []int {
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	minVal := nums[0]
	for _, v := range nums {
		if v < minVal {
			minVal = v
		}
	}
	minFreq := freq[minVal]
  // Membuat map untuk pencarian O(1): key → value
	seen := make(map[int]bool)
	for _, v := range nums {
		if seen[v] {
			continue
		}
		seen[v] = true
		if v > minVal && freq[v] != minFreq {
			return []int{minVal, v}
		}
	}
	return []int{-1, -1}
}
```
