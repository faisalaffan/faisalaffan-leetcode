# 1317 — Convert Integer To The Sum Of Two No Zero Integers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func getNoZeroIntegers(n int) []int

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1317: Convert Integer to the Sum of Two No-Zero Integers
// https://leetcode.com/problems/convert-integer-to-the-sum-of-two-no-zero-integers/
// Difficulty: Easy
//
// LeetCode submission: func getNoZeroIntegers(n int) []int

import "fmt"

func main() {
	fmt.Println(ConvertIntegerToTheSumOfTwoNoZeroIntegers(2))    // [1 1]
	fmt.Println(ConvertIntegerToTheSumOfTwoNoZeroIntegers(11))   // [2 9]
	fmt.Println(ConvertIntegerToTheSumOfTwoNoZeroIntegers(1010)) // [122 888]
}

// Time: O(n log n), Space: O(1)
func ConvertIntegerToTheSumOfTwoNoZeroIntegers(n int) []int {
	for a := 1; a < n; a++ {
		b := n - a
		if !hasZero(a) && !hasZero(b) {
			return []int{a, b}
		}
	}
	return []int{}
}

func hasZero(x int) bool {
	for x > 0 {
		if x%10 == 0 {
			return true
		}
		x /= 10
	}
	return false
}
```
