# 2732 — Find A Good Subset Of The Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func findAGoodSubsetOfTheMatrix(grid [][]int) []int
```

> **💡 Hint:** We never need more than 2 rows.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2732: Find a Good Subset of the Matrix
// https://leetcode.com/problems/find-a-good-subset-of-the-matrix/
// Difficulty: Hard
//
// Approach: We never need more than 2 rows.
// - Single row: valid if all zeros (each column sum = 0 <= floor(1/2) = 0).
// - Two rows: valid if bitwise AND = 0 (no column has 1s in both rows).
// If neither exists, no valid subset exists.

import "fmt"

func findAGoodSubsetOfTheMatrix(grid [][]int) []int {
	m := len(grid)
	if m == 0 {
		return []int{}
	}
	n := len(grid[0])

  // Alokasi slice integer
	masks := make([]int, m)
	allZero := -1
	for i := 0; i < m; i++ {
		mask := 0
		zero := true
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				mask |= (1 << j)
				zero = false
			}
		}
		masks[i] = mask
		if zero && allZero == -1 {
			allZero = i
		}
	}

	// Single row: must be all zeros
	if allZero >= 0 {
		return []int{allZero}
	}

	// Two rows: bitwise AND must be 0
	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {
			if masks[i]&masks[j] == 0 {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func main() {
	// Example 1
	fmt.Println(findAGoodSubsetOfTheMatrix([][]int{{0, 1, 1, 0}, {0, 0, 0, 1}, {1, 1, 1, 1}}))
	// Example 2
	fmt.Println(findAGoodSubsetOfTheMatrix([][]int{{0}}))
	// Example 3: no good subset
	fmt.Println(findAGoodSubsetOfTheMatrix([][]int{{1, 1, 1}, {1, 1, 1}}))
	// All-zero row present
	fmt.Println(findAGoodSubsetOfTheMatrix([][]int{{1, 1}, {0, 0}, {1, 0}}))
	// Empty matrix
	fmt.Println(findAGoodSubsetOfTheMatrix([][]int{}))
}
```
