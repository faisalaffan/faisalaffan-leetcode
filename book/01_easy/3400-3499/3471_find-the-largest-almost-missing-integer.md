# 3471 — Find The Largest Almost Missing Integer

## Deskripsi

**Soal:** [3471. Find The Largest Almost Missing Integer](https://leetcode.com/problems/find-the-largest-almost-missing-integer/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3471: Find the Largest Almost Missing Integer
// https://leetcode.com/problems/find-the-largest-almost-missing-integer/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindTheLargestAlmostMissingInteger([]int{3, 9, 2, 3, 1, 6, 7, 8, 9}, 2))
	fmt.Println(FindTheLargestAlmostMissingInteger([]int{0, 0}, 1))
}

// FindTheLargestAlmostMissingInteger returns the largest integer that appears fewer than k times in nums.
// Time: O(n). Space: O(n).
func FindTheLargestAlmostMissingInteger(nums []int, k int) int {
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	largest := -1
	for val, count := range freq {
		if count < k && val > largest {
			largest = val
		}
	}
	return largest
}
```
