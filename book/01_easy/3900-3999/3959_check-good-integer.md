# 3959 — Check Good Integer

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CheckGoodInteger(n int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3959: Check Good Integer
// https://leetcode.com/problems/check-good-integer/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckGoodInteger(1000))
	fmt.Println(CheckGoodInteger(19))
}

// Time: O(log n)
// Space: O(1)
func CheckGoodInteger(n int) bool {
	total := 0
	for n > 0 {
		d := n % 10
		total += d * (d - 1)
		n /= 10
	}
	return total >= 50
}
```
