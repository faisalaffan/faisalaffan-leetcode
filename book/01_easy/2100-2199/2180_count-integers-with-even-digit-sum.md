# 2180 — Count Integers With Even Digit Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CountIntegersWithEvenDigitSum(num int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n log n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2180: Count Integers With Even Digit Sum
// https://leetcode.com/problems/count-integers-with-even-digit-sum/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountIntegersWithEvenDigitSum(4))  // 2
	fmt.Println(CountIntegersWithEvenDigitSum(30))  // 14
}

// Time: O(n log n), Space: O(1)
func CountIntegersWithEvenDigitSum(num int) int {
	count := 0
	for i := 1; i <= num; i++ {
		if digitSumEven(i) {
			count++
		}
	}
	return count
}

func digitSumEven(n int) bool {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum%2 == 0
}
```
