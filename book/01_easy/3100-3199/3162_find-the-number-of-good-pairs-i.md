# 3162 — Find The Number Of Good Pairs I

## Deskripsi

**Soal:** [3162. Find The Number Of Good Pairs I](https://leetcode.com/problems/find-the-number-of-good-pairs-i/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n * m)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** Dynamic Programming (DP)

## Solusi Go

```go
package main

// LeetCode #3162: Find the Number of Good Pairs I
// https://leetcode.com/problems/find-the-number-of-good-pairs-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: numberOfPairs
	fmt.Println(FindTheNumberOfGoodPairsI([]int{1, 2, 4, 12}, []int{2, 4}, 3)) // 2
	fmt.Println(FindTheNumberOfGoodPairsI([]int{1, 2, 3}, []int{1, 2, 3}, 1))  // 5
}

// Time: O(n * m) | Space: O(1)
// LeetCode submission name: numberOfPairs
func FindTheNumberOfGoodPairsI(nums1 []int, nums2 []int, k int) int {
	count := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i]%(nums2[j]*k) == 0 {
				count++
			}
		}
	}
	return count
}
```
