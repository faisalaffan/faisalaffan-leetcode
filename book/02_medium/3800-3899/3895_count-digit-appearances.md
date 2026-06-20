# 3895 — Count Digit Appearances

## Deskripsi

**Soal:** [3895. Count Digit Appearances](https://leetcode.com/problems/count-digit-appearances/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N * log M)  
**Kompleksitas Ruang:** O(1) where M = max value in nums

**Algoritma:** —

**Fungsi Solusi:** `func CountDigitAppearances(nums []int, digit int) int`

> **Ide Kunci:** For each number, extract digits and count matches to target digit.

## Solusi Go

```go
package main

// LeetCode #3895: Count Digit Appearances
// https://leetcode.com/problems/count-digit-appearances/
// Difficulty: Medium
// Time: O(N * log M) | Space: O(1) where M = max value in nums
// Approach: For each number, extract digits and count matches to target digit.

import "fmt"

func CountDigitAppearances(nums []int, digit int) int {
	ans := 0
	for _, v := range nums {
		for v > 0 {
			if v%10 == digit {
				ans++
			}
			v /= 10
		}
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(CountDigitAppearances([]int{12, 54, 32, 22}, 2)) // Expected: 4

	// Example 2
	fmt.Println(CountDigitAppearances([]int{1, 34, 7}, 9)) // Expected: 0
}
```
