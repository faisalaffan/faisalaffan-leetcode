# 1605 — Find Valid Matrix Given Row And Column Sums

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func minInt(a, b int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(R*C), Space: O(R*C)  |  **Ruang:** O(R*C)


## 💻 Solusi Go

```go
package main

// LeetCode #1605: Find Valid Matrix Given Row and Column Sums
// https://leetcode.com/problems/find-valid-matrix-given-row-and-column-sums/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(RestoreMatrix([]int{3, 8}, []int{4, 7}))
	fmt.Println(RestoreMatrix([]int{5, 7, 10}, []int{8, 6, 8}))
	fmt.Println(RestoreMatrix([]int{14, 9}, []int{6, 9, 8}))
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func RestoreMatrix(rowSum []int, colSum []int) [][]int {
	// Time: O(R*C), Space: O(R*C)
	rows, cols := len(rowSum), len(colSum)
  // Matriks 2D
	result := make([][]int, rows)
	for i := 0; i < rows; i++ {
		result[i] = make([]int, cols)
	}

	i, j := 0, 0
	for i < rows && j < cols {
		val := minInt(rowSum[i], colSum[j])
		result[i][j] = val
		rowSum[i] -= val
		colSum[j] -= val

		if rowSum[i] == 0 {
			i++
		}
		if colSum[j] == 0 {
			j++
		}
	}

	return result
}
```
