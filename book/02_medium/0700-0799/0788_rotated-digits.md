# 0788 — Rotated Digits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func rotatedDigits(n int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #788: Rotated Digits
// https://leetcode.com/problems/rotated-digits/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(rotatedDigits(10))
	fmt.Println(rotatedDigits(1))
}

func rotatedDigits(n int) int {
	count := 0

	for i := 1; i <= n; i++ {
		if isGood(i) {
			count++
		}
	}

	return count
}

func isGood(n int) bool {
	valid := false
	for n > 0 {
		d := n % 10
		if d == 3 || d == 4 || d == 7 {
			return false
		}
		if d == 2 || d == 5 || d == 6 || d == 9 {
			valid = true
		}
		n /= 10
	}
	return valid
}
```
