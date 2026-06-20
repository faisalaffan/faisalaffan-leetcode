# 0052 — N Queens Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func totalNQueens(n int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Backtracking

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Backtracking** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #52: N-Queens II
// https://leetcode.com/problems/n-queens-ii/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println("52. N-Queens II")
	fmt.Println("n=4:", totalNQueens(4), "(expected 2)")
	fmt.Println("n=1:", totalNQueens(1), "(expected 1)")
	fmt.Println("n=8:", totalNQueens(8), "(expected 92)")
}

func totalNQueens(n int) int {
	cols := make([]bool, n)
	diag1 := make([]bool, 2*n-1) // r+c
	diag2 := make([]bool, 2*n-1) // r-c+n-1
	count := 0
	backtrack(n, 0, cols, diag1, diag2, &count)
	return count
}

func backtrack(n, row int, cols, diag1, diag2 []bool, count *int) {
	if row == n {
		*count++
		return
	}
	for col := 0; col < n; col++ {
		d1 := row + col
		d2 := row - col + n - 1
		if cols[col] || diag1[d1] || diag2[d2] {
			continue
		}
		cols[col], diag1[d1], diag2[d2] = true, true, true
		backtrack(n, row+1, cols, diag1, diag2, count)
		cols[col], diag1[d1], diag2[d2] = false, false, false
	}
}
```
