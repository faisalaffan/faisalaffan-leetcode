# 2243 — Calculate Digit Sum Of A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func CalculateDigitSumOfAString(s string, k int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2243: Calculate Digit Sum of a String
// https://leetcode.com/problems/calculate-digit-sum-of-a-string/
// Difficulty: Easy

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(CalculateDigitSumOfAString("11111222223", 3)) // "135"
	fmt.Println(CalculateDigitSumOfAString("00000000", 3))    // "000"
}

// Time: O(n), Space: O(n)
func CalculateDigitSumOfAString(s string, k int) string {
	for len(s) > k {
		var next string
  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(s); i += k {
			end := i + k
			if end > len(s) {
				end = len(s)
			}
			sum := 0
			for j := i; j < end; j++ {
				sum += int(s[j] - '0')
			}
			next += strconv.Itoa(sum)
		}
		s = next
	}
	return s
}
```
