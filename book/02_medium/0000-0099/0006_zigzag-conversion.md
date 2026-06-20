# 0006 — Zigzag Conversion

## Deskripsi

**Soal:** [0006. Zigzag Conversion](https://leetcode.com/problems/zigzag-conversion/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func convert(s string, numRows int) string`

## Solusi Go

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

  // Membuat slice 2D untuk DP/tabel
	rows := make([][]byte, numRows)
	curRow := 0
	goingDown := false

  // Loop standar: indeks 0 sampai n-1
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

  // Membuat slice untuk menyimpan hasil
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
