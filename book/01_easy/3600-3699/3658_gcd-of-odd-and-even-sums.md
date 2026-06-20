# 3658 — Gcd Of Odd And Even Sums

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func GcdOfOddAndEvenSums(n int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3658: GCD of Odd and Even Sums
// https://leetcode.com/problems/gcd-of-odd-and-even-sums/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(GcdOfOddAndEvenSums(4))
	fmt.Println(GcdOfOddAndEvenSums(1))
	fmt.Println(GcdOfOddAndEvenSums(10))
}

// Time: O(1)
// Space: O(1)
// The answer is always n because:
// sum of first n odd numbers = n^2
// sum of first n even numbers = n(n+1)
// gcd(n^2, n(n+1)) = n (since n and n+1 are coprime)
func GcdOfOddAndEvenSums(n int) int {
	return n
}
```
