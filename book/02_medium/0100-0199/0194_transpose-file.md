# 0194 — Transpose File

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func transposeFile(content string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m*n), Space: O(m*n)  
**Kompleksitas Ruang:** O(m*n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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

  // Membuat matriks/slice 2D untuk DP
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
