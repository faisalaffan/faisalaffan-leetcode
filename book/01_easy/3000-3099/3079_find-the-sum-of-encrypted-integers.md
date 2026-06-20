# 3079 — Find The Sum Of Encrypted Integers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheSumOfEncryptedIntegers(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * d) where d is number of digits  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3079: Find the Sum of Encrypted Integers
// https://leetcode.com/problems/find-the-sum-of-encrypted-integers/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: sumOfEncryptedInt
	fmt.Println(FindTheSumOfEncryptedIntegers([]int{10, 21, 31})) // 66
	fmt.Println(FindTheSumOfEncryptedIntegers([]int{1, 2, 3}))    // 6
}

// Time: O(n * d) where d is number of digits | Space: O(1)
// LeetCode submission name: sumOfEncryptedInt
func FindTheSumOfEncryptedIntegers(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += encrypt(v)
	}
	return sum
}

func encrypt(n int) int {
	maxDigit := 0
	digits := 0
	for n > 0 {
		d := n % 10
		if d > maxDigit {
			maxDigit = d
		}
		digits++
		n /= 10
	}
	result := 0
	for i := 0; i < digits; i++ {
		result = result*10 + maxDigit
	}
	return result
}
```
