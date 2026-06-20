# 3945 — Digit Frequency Score

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func DigitFrequencyScore(n int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3945: Digit Frequency Score
// https://leetcode.com/problems/digit-frequency-score/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(DigitFrequencyScore(122))
	fmt.Println(DigitFrequencyScore(101))
}

// Time: O(log n)
// Space: O(1)
func DigitFrequencyScore(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
```
