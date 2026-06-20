# 3917 — Count Indices With Opposite Parity

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountIndicesWithOppositeParity(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3917: Count Indices With Opposite Parity
// https://leetcode.com/problems/count-indices-with-opposite-parity/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountIndicesWithOppositeParity([]int{1, 2, 3, 4}))
	fmt.Println(CountIndicesWithOppositeParity([]int{2, 4, 6}))
}

// Time: O(n)
// Space: O(1)
func CountIndicesWithOppositeParity(nums []int) []int {
	totalEven, totalOdd := 0, 0
	for _, v := range nums {
		if v%2 == 0 {
			totalEven++
		} else {
			totalOdd++
		}
	}

  // Alokasi slice integer
	ans := make([]int, len(nums))
	for i, v := range nums {
		if v%2 == 0 {
			totalEven--
			ans[i] = totalOdd
		} else {
			totalOdd--
			ans[i] = totalEven
		}
	}
	return ans
}
```
