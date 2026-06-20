# 1291 — Sequential Digits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func sequentialDigits(low int, high int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1) — bounded by number of possible sequential digits  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1291: Sequential Digits
// https://leetcode.com/problems/sequential-digits/
// Difficulty: Medium

// Generate all sequential digits numbers in range [low, high].
// Sequential = each digit is 1 more than previous.

// Time: O(1) — bounded by number of possible sequential digits
// Space: O(1)

func sequentialDigits(low int, high int) []int {
  // Alokasi slice integer
	result := make([]int, 0)
	digits := "123456789"

	for length := len(fmt.Sprintf("%d", low)); length <= len(fmt.Sprintf("%d", high)); length++ {
		for start := 0; start+length <= 9; start++ {
			num := 0
			for i := 0; i < length; i++ {
				num = num*10 + int(digits[start+i]-'0')
			}
			if num >= low && num <= high {
				result = append(result, num)
			}
		}
	}

  // Urutkan secara ascending — O(n log n)
	sort.Ints(result)
	return result
}

func main() {
	fmt.Printf("%v (expected: [123 234])\n", sequentialDigits(100, 300))
	fmt.Printf("%v (expected: [1234 2345 3456 4567 5678 6789 12345])\n", sequentialDigits(1000, 13000))
}
```
