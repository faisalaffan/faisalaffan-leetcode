# 0728 — Self Dividing Numbers

## Deskripsi

**Soal:** [0728. Self Dividing Numbers](https://leetcode.com/problems/self-dividing-numbers/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n * d) where n = count, d = digits per number. Space: O(1) excluding output.  
**Kompleksitas Ruang:** O(1) excluding output.

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #728: Self Dividing Numbers
// https://leetcode.com/problems/self-dividing-numbers/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(selfDividingNumbers(1, 22)) // [1,2,3,4,5,6,7,8,9,11,12,15,22]
	fmt.Println(selfDividingNumbers(47, 85)) // [48,55,66,77]
}

// selfDividingNumbers returns all self-dividing numbers in range [left, right].
// Time: O(n * d) where n = count, d = digits per number. Space: O(1) excluding output.
func selfDividingNumbers(left int, right int) []int {
  // Membuat slice untuk menyimpan hasil
	result := make([]int, 0)
	for n := left; n <= right; n++ {
		if isSelfDividing(n) {
			result = append(result, n)
		}
	}
	return result
}

func isSelfDividing(n int) bool {
	original := n
	for n > 0 {
		digit := n % 10
		if digit == 0 || original%digit != 0 {
			return false
		}
		n /= 10
	}
	return true
}
```
