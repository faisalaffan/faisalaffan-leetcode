# 2133 — Check If Every Row And Column Contains All Numbers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckIfEveryRowAndColumnContainsAllNumbers(matrix [][]int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2133: Check if Every Row and Column Contains All Numbers
// https://leetcode.com/problems/check-if-every-row-and-column-contains-all-numbers/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckIfEveryRowAndColumnContainsAllNumbers([][]int{{1, 2, 3}, {3, 1, 2}, {2, 3, 1}})) // true
	fmt.Println(CheckIfEveryRowAndColumnContainsAllNumbers([][]int{{1, 1, 1}, {1, 2, 3}, {1, 2, 3}})) // false
}

// Time: O(n^2), Space: O(n)
func CheckIfEveryRowAndColumnContainsAllNumbers(matrix [][]int) bool {
	n := len(matrix)

	for i := 0; i < n; i++ {
		rowSet := make([]bool, n+1)
		colSet := make([]bool, n+1)
		for j := 0; j < n; j++ {
			if matrix[i][j] < 1 || matrix[i][j] > n || rowSet[matrix[i][j]] {
				return false
			}
			rowSet[matrix[i][j]] = true

			if matrix[j][i] < 1 || matrix[j][i] > n || colSet[matrix[j][i]] {
				return false
			}
			colSet[matrix[j][i]] = true
		}
	}
	return true
}
```
