# 0194 — Transpose File

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func transposeFile(content string) []string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(m*n), Space: O(m*n)  |  **Ruang:** O(m*n)


## 💻 Solusi Go

```go
package main

// LeetCode #194: Transpose File
// https://leetcode.com/problems/transpose-file/
// Difficulty: Medium
// Time: O(m*n), Space: O(m*n)

import (
	"fmt"
	"strings"
)

func transposeFile(content string) []string {
	lines := strings.Split(strings.TrimSpace(content), "\n")
	if len(lines) == 0 || lines[0] == "" {
		return nil
	}

  // Matriks 2D
	matrix := make([][]string, len(lines))
	for i, line := range lines {
		matrix[i] = strings.Fields(line)
	}

	if len(matrix) == 0 {
		return nil
	}

	cols := 0
	for _, row := range matrix {
		if len(row) > cols {
			cols = len(row)
		}
	}

	result := make([]string, cols)
	for c := 0; c < cols; c++ {
		var row []string
		for r := 0; r < len(matrix); r++ {
			if c < len(matrix[r]) {
				row = append(row, matrix[r][c])
			}
		}
		result[c] = strings.Join(row, " ")
	}

	return result
}

func main() {
	content := "name age\nalice 21\nryan 30"
	for _, line := range transposeFile(content) {
		fmt.Println(line)
	}
	fmt.Println("---")
	content2 := "a b c\nd e f"
	for _, line := range transposeFile(content2) {
		fmt.Println(line)
	}
}
```
