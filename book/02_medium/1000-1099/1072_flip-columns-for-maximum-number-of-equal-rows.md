# 1072 — Flip Columns For Maximum Number Of Equal Rows

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxEqualRowsAfterFlips(matrix [][]int) int
```

> **💡 Hint:** Normalize each row to a pattern (starting with 0).

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(m * n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1072: Flip Columns For Maximum Number of Equal Rows
// https://leetcode.com/problems/flip-columns-for-maximum-number-of-equal-rows/
// Difficulty: Medium
//
// Approach: Normalize each row to a pattern (starting with 0).
//           Rows with same or complementary pattern can be made equal.
// Time: O(m * n)
// Space: O(m * n)

import "fmt"

func main() {
	fmt.Println(maxEqualRowsAfterFlips([][]int{{0, 1}, {1, 1}})) // 1
	fmt.Println(maxEqualRowsAfterFlips([][]int{{0, 0, 0}, {0, 0, 1}, {1, 1, 0}})) // 2
}

func maxEqualRowsAfterFlips(matrix [][]int) int {
  // Membuat map (HashMap) — pencarian O(1)
	patternCount := make(map[string]int)

	for _, row := range matrix {
		pattern := make([]byte, len(row))
		for j := 0; j < len(row); j++ {
			if row[0] == 0 {
				pattern[j] = byte('0' + row[j])
			} else {
				pattern[j] = byte('0' + 1 - row[j])
			}
		}
		patternCount[string(pattern)]++
	}

	result := 0
	for _, count := range patternCount {
		if count > result {
			result = count
		}
	}
	return result
}
```
