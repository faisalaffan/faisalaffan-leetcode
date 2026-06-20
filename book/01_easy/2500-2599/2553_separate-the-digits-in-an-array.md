# 2553 — Separate The Digits In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func SeparateTheDigitsInAnArray(nums []int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2553: Separate the Digits in an Array
// https://leetcode.com/problems/separate-the-digits-in-an-array/
// Difficulty: Easy
// Time O(n log m) | Space O(n log m)

import "fmt"

func main() {
	fmt.Println(SeparateTheDigitsInAnArray([]int{13, 25, 83, 77})) // [1,3,2,5,8,3,7,7]
	fmt.Println(SeparateTheDigitsInAnArray([]int{7, 1, 3, 9}))     // [7,1,3,9]
}

func SeparateTheDigitsInAnArray(nums []int) []int {
	res := []int{}
	for _, n := range nums {
		digits := []int{}
		for n > 0 {
			digits = append(digits, n%10)
			n /= 10
		}
		for i := len(digits) - 1; i >= 0; i-- {
			res = append(res, digits[i])
		}
	}
	return res
}
```
