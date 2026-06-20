# 3895 — Count Digit Appearances

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CountDigitAppearances(nums []int, digit int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(N * log M)  |  **Ruang:** O(1) where M = max value in nums


## 💻 Solusi Go

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
