# 1267 — Count Servers That Communicate

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countServers(grid [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m*n)  
**Kompleksitas Ruang:** O(m+n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1267: Count Servers that Communicate
// https://leetcode.com/problems/count-servers-that-communicate/
// Difficulty: Medium

// Count servers that can communicate with at least one other server
// in the same row or column.

// Time: O(m*n)
// Space: O(m+n)

func countServers(grid [][]int) int {
	m, n := len(grid), len(grid[0])
  // Alokasi slice integer
	rowCount := make([]int, m)
  // Alokasi slice integer
	colCount := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				rowCount[i]++
				colCount[j]++
			}
		}
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && (rowCount[i] > 1 || colCount[j] > 1) {
				count++
			}
		}
	}

	return count
}

func main() {
	fmt.Printf("%d (expected: 3)\n",
		countServers([][]int{{1, 0}, {0, 1}}))

	fmt.Printf("%d (expected: 4)\n",
		countServers([][]int{{1, 0}, {1, 1}}))

	fmt.Printf("%d (expected: 0)\n",
		countServers([][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}}))
}
```
