# 0166 — Fraction To Recurring Decimal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func fractionToDecimal(numerator int, denominator int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n) where n is the length of the repeating cycle, Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #166: Fraction to Recurring Decimal
// https://leetcode.com/problems/fraction-to-recurring-decimal/
// Difficulty: Medium
// Time: O(n) where n is the length of the repeating cycle, Space: O(n)

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	result := ""

	if (numerator < 0) != (denominator < 0) {
		result += "-"
	}

	num := abs(numerator)
	den := abs(denominator)

	result += strconv.Itoa(num / den)

	remainder := num % den
	if remainder == 0 {
		return result
	}

	result += "."

  // Membuat map (HashMap) — pencarian O(1)
	remainderMap := make(map[int]int)
	remainderMap[remainder] = len(result)

	for remainder != 0 {
		remainder *= 10
		result += strconv.Itoa(remainder / den)
		remainder = remainder % den

		if pos, ok := remainderMap[remainder]; ok {
			result = result[:pos] + "(" + result[pos:] + ")"
			break
		}
		remainderMap[remainder] = len(result)
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(fractionToDecimal(1, 2))
	fmt.Println(fractionToDecimal(2, 1))
	fmt.Println(fractionToDecimal(4, 333))
}
```
