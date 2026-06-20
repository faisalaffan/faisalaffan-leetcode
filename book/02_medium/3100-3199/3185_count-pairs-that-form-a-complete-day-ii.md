# 3185 — Count Pairs That Form A Complete Day Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countCompleteDayPairs(hours []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(24)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3185: Count Pairs That Form a Complete Day II
// https://leetcode.com/problems/count-pairs-that-form-a-complete-day-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(24)

import "fmt"

func countCompleteDayPairs(hours []int) int64 {
  // Alokasi slice integer
	count := make([]int, 24)
	var ans int64

	for _, h := range hours {
		r := h % 24
		need := (24 - r) % 24
		ans += int64(count[need])
		count[r]++
	}
	return ans
}

func main() {
	fmt.Println(countCompleteDayPairs([]int{12, 12, 30, 24, 24})) // Expected: 2
	fmt.Println(countCompleteDayPairs([]int{72, 48, 24, 3}))       // Expected: 3
}
```
