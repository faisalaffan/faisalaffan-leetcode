# 3184 — Count Pairs That Form A Complete Day I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountPairsThatFormACompleteDayI(hours []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(24).  
**Kompleksitas Ruang:** O(24).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3184: Count Pairs That Form a Complete Day I
// https://leetcode.com/problems/count-pairs-that-form-a-complete-day-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountPairsThatFormACompleteDayI([]int{12, 12, 30, 24, 24}))
	fmt.Println(CountPairsThatFormACompleteDayI([]int{72, 48, 24, 3}))
}

// CountPairsThatFormACompleteDayI counts pairs (i, j) where i < j and hours[i] + hours[j] is divisible by 24.
// Time: O(n). Space: O(24).
func CountPairsThatFormACompleteDayI(hours []int) int {
	count := 0
  // Alokasi slice integer
	rem := make([]int, 24)
	for _, h := range hours {
		r := h % 24
		need := (24 - r) % 24
		count += rem[need]
		rem[r]++
	}
	return count
}
```
