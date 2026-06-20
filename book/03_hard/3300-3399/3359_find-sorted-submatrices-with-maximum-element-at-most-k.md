# 3359 — Find Sorted Submatrices With Maximum Element At Most K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindSortedSubmatricesWithMaximumElementAtMostK(grid [][]int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack, Monotonic Stack/Queue

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3359: Find Sorted Submatrices With Maximum Element at Most K
// https://leetcode.com/problems/find-sorted-submatrices-with-maximum-element-at-most-k/
// Difficulty: Hard [Paid]
//
// Count submatrices where max element <= K and each row is non-increasing.
// Histogram + monotonic stack.

import "fmt"

func main() {
	grid := [][]int{{3, 2, 1}, {2, 1, 1}, {1, 1, 1}}
	fmt.Println(FindSortedSubmatricesWithMaximumElementAtMostK(grid, 3))
}

func FindSortedSubmatricesWithMaximumElementAtMostK(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
  // Membuat matriks/slice 2D untuk DP
	rows := make([][]int, m)
	for i := 0; i < m; i++ {
		rows[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if grid[i][j] > k {
				rows[i][j] = 0
			} else if j > 0 && grid[i][j] <= grid[i][j-1] {
				rows[i][j] = rows[i][j-1] + 1
			} else {
				rows[i][j] = 1
			}
		}
	}

	result := 0
	for j := 0; j < n; j++ {
		type pair struct{ val, cnt int }
		stack := make([]pair, 0, m)
		total := 0
		for i := 0; i < m; i++ {
			cnt := 1
			for len(stack) > 0 && stack[len(stack)-1].val >= rows[i][j] {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				cnt += top.cnt
				total -= (top.val - rows[i][j]) * top.cnt
			}
			total += rows[i][j]
			stack = append(stack, pair{rows[i][j], cnt})
			result += total
		}
	}
	return result
}
```
