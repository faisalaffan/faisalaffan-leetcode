# 2834 — Find The Minimum Possible Sum Of A Beautiful Array

## Deskripsi

**Soal:** [2834. Find The Minimum Possible Sum Of A Beautiful Array](https://leetcode.com/problems/find-the-minimum-possible-sum-of-a-beautiful-array/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func FindTheMinimumPossibleSumOfABeautifulArray(n int, target int) int`

## Solusi Go

```go
package main

// LeetCode #2834: Find the Minimum Possible Sum of a Beautiful Array
// https://leetcode.com/problems/find-the-minimum-possible-sum-of-a-beautiful-array/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func FindTheMinimumPossibleSumOfABeautifulArray(n int, target int) int {
  // Membuat map untuk pencarian O(1): key → value
	used := make(map[int]bool)
	sum := 0
	for i := 1; len(used) < n; i++ {
		if used[target-i] {
			continue
		}
		used[i] = true
		sum += i
	}
	return sum
}

func main() {
	fmt.Println(FindTheMinimumPossibleSumOfABeautifulArray(2, 3))
	fmt.Println(FindTheMinimumPossibleSumOfABeautifulArray(3, 3))
	fmt.Println(FindTheMinimumPossibleSumOfABeautifulArray(5, 5))
}
```
