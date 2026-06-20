# 1085 — Sum Of Digits In The Minimum Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func sumOfDigits(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1085: Sum of Digits in the Minimum Number
// https://leetcode.com/problems/sum-of-digits-in-the-minimum-number/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(sumOfDigits([]int{34, 23, 1, 24, 75, 33, 54, 8})) // 0 (min=1, sum=1, odd)
	fmt.Println(sumOfDigits([]int{99, 77, 33, 66, 55}))           // 1 (min=33, sum=6, even)
}

// LeetCode submission: sumOfDigits
func sumOfDigits(nums []int) int {
	x := nums[0]
	for _, v := range nums {
		if v < x {
			x = v
		}
	}
	s := 0
	for x > 0 {
		s += x % 10
		x /= 10
	}
	if s%2 == 0 {
		return 1
	}
	return 0
}
```
