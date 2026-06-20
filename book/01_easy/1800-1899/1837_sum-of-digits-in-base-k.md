# 1837 — Sum Of Digits In Base K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func SumBase(n int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log_k(n)), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1837: Sum of Digits in Base K
// https://leetcode.com/problems/sum-of-digits-in-base-k/
// Difficulty: Easy

import "fmt"

// Time: O(log_k(n)), Space: O(1)
func SumBase(n int, k int) int {
	sum := 0
	for n > 0 {
		sum += n % k
		n /= k
	}
	return sum
}

func main() {
	fmt.Println(SumBase(34, 6))
	fmt.Println(SumBase(10, 10))
	fmt.Println(SumBase(42, 2))
}
```
