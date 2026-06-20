# 1295 — Find Numbers With Even Number Of Digits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func findNumbers(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1295: Find Numbers with Even Number of Digits
// https://leetcode.com/problems/find-numbers-with-even-number-of-digits/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(findNumbers([]int{12, 345, 2, 6, 7896})) // 2
	fmt.Println(findNumbers([]int{555, 901, 482, 1771})) // 1
}

// LeetCode submission: findNumbers
func findNumbers(nums []int) int {
	count := 0
	for _, v := range nums {
		digits := 0
		for x := v; x > 0; x /= 10 {
			digits++
		}
		if digits%2 == 0 {
			count++
		}
	}
	return count
}
```
