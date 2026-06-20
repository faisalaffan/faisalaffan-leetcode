# 3726 — Remove Zeros In Decimal Representation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func RemoveZerosInDecimalRepresentation(n int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n) - number of digits  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3726: Remove Zeros in Decimal Representation
// https://leetcode.com/problems/remove-zeros-in-decimal-representation/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(RemoveZerosInDecimalRepresentation(1020030))
	fmt.Println(RemoveZerosInDecimalRepresentation(1000))
}

// Time: O(log n) - number of digits
// Space: O(1)
func RemoveZerosInDecimalRepresentation(n int) int {
	ans := 0
	k := 1
	for n > 0 {
		x := n % 10
		if x > 0 {
			ans = k*x + ans
			k *= 10
		}
		n /= 10
	}
	return ans
}
```
