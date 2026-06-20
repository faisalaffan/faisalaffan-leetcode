# 3658 — Gcd Of Odd And Even Sums

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func GcdOfOddAndEvenSums(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** GCD / Matematika

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **GCD / Matematika** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
