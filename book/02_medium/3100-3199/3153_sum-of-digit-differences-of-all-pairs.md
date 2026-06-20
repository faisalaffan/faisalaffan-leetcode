# 3153 — Sum Of Digit Differences Of All Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func sumDigitDifferences(nums []int) int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * d)  |  **Ruang:** O(d * 10)


## 💻 Solusi Go

```go
package main

// LeetCode #3153: Sum of Digit Differences of All Pairs
// https://leetcode.com/problems/sum-of-digit-differences-of-all-pairs/
// Difficulty: Medium
// Time: O(n * d) | Space: O(d * 10)

import "fmt"

func sumDigitDifferences(nums []int) int64 {
	n := len(nums)
	if n < 2 {
		return 0
	}

	// Find number of digits
	x := nums[0]
	digits := 0
	for x > 0 {
		digits++
		x /= 10
	}

	var ans int64
	pow := 1
	for d := 0; d < digits; d++ {
  // Alokasi slice
		count := make([]int, 10)
		for _, num := range nums {
			count[(num/pow)%10]++
		}
		totalPairs := int64(n) * int64(n-1) / 2
		for _, c := range count {
			if c > 1 {
				totalPairs -= int64(c) * int64(c-1) / 2
			}
		}
		ans += totalPairs
		pow *= 10
	}
	return ans
}

func main() {
	fmt.Println(sumDigitDifferences([]int{13, 23, 12})) // Expected: 4
	fmt.Println(sumDigitDifferences([]int{10, 10, 10})) // Expected: 0
}
```
