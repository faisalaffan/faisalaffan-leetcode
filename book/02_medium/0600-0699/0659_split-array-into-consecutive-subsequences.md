# 0659 — Split Array Into Consecutive Subsequences

## Deskripsi

**Soal:** [0659. Split Array Into Consecutive Subsequences](https://leetcode.com/problems/split-array-into-consecutive-subsequences/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #659: Split Array into Consecutive Subsequences
// https://leetcode.com/problems/split-array-into-consecutive-subsequences/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(isPossible([]int{1, 2, 3, 3, 4, 5}))
	fmt.Println(isPossible([]int{1, 2, 3, 3, 4, 4, 5, 5}))
	fmt.Println(isPossible([]int{1, 2, 3, 4, 4, 5}))
}

func isPossible(nums []int) bool {
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[int]int)
  // Membuat map untuk pencarian O(1): key → value
	tail := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	for _, num := range nums {
		if freq[num] == 0 {
			continue
		}

		if tail[num] > 0 {
			tail[num]--
			freq[num]--
			tail[num+1]++
		} else if freq[num+1] > 0 && freq[num+2] > 0 {
			freq[num]--
			freq[num+1]--
			freq[num+2]--
			tail[num+3]++
		} else {
			return false
		}
	}

	return true
}
```
