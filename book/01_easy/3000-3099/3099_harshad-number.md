# 3099 — Harshad Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func HarshadNumber(x int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3099: Harshad Number
// https://leetcode.com/problems/harshad-number/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: sumOfTheDigitsOfHarshadNumber
	fmt.Println(HarshadNumber(18)) // 9
	fmt.Println(HarshadNumber(23)) // -1
}

// Time: O(log n) | Space: O(1)
// LeetCode submission name: sumOfTheDigitsOfHarshadNumber
func HarshadNumber(x int) int {
	sum := 0
	n := x
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	if x%sum == 0 {
		return sum
	}
	return -1
}
```
