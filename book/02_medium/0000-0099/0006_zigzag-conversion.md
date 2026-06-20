# 0006 — Zigzag Conversion

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func convert(s string, numRows int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #6: Zigzag Conversion
// https://leetcode.com/problems/zigzag-conversion/
// Difficulty: Medium

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

  // Matriks 2D
	rows := make([][]byte, numRows)
	curRow := 0
	goingDown := false

  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		rows[curRow] = append(rows[curRow], s[i])
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			curRow++
		} else {
			curRow--
		}
	}

	result := make([]byte, 0, len(s))
	for _, row := range rows {
		result = append(result, row...)
	}

	return string(result)
}

func main() {
	// Test case 1
	fmt.Println(convert("PAYPALISHIRING", 3)) // "PAHNAPLSIIGYIR"

	// Test case 2
	fmt.Println(convert("PAYPALISHIRING", 4)) // "PINALSIGYAHRPI"

	// Test case 3
	fmt.Println(convert("A", 1)) // "A"
}

// Time: O(n) | Space: O(n)
```
