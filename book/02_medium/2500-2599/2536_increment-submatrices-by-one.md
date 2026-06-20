# 2536 — Increment Submatrices By One

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func rangeAddQueries(n int, queries [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2 + q)  
**Kompleksitas Ruang:** O(n^2)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2536: Increment Submatrices by One
// https://leetcode.com/problems/increment-submatrices-by-one/
// Difficulty: Medium
// Time: O(n^2 + q) | Space: O(n^2)

import "fmt"

func rangeAddQueries(n int, queries [][]int) [][]int {
  // Membuat matriks/slice 2D untuk DP
	diff := make([][]int, n+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range diff {
		diff[i] = make([]int, n+1)
	}

	for _, q := range queries {
		r1, c1, r2, c2 := q[0], q[1], q[2], q[3]
		diff[r1][c1]++
		diff[r1][c2+1]--
		diff[r2+1][c1]--
		diff[r2+1][c2+1]++
	}

  // Membuat matriks/slice 2D untuk DP
	mat := make([][]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range mat {
		mat[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i > 0 {
				diff[i][j] += diff[i-1][j]
			}
			if j > 0 {
				diff[i][j] += diff[i][j-1]
			}
			if i > 0 && j > 0 {
				diff[i][j] -= diff[i-1][j-1]
			}
			mat[i][j] = diff[i][j]
		}
	}
	return mat
}

func main() {
	// Test case 1: n=2, queries=[[0,0,0,0]]
	res1 := rangeAddQueries(2, [][]int{{0, 0, 0, 0}})
	fmt.Println("Test 1:", res1)
	// Expected: [[1,0],[0,0]]

	// Test case 2: n=2, queries=[[0,0,1,1]]
	res2 := rangeAddQueries(2, [][]int{{0, 0, 1, 1}})
	fmt.Println("Test 2:", res2)
	// Expected: [[1,1],[1,1]]

	// Test case 3: n=1
	res3 := rangeAddQueries(1, [][]int{{0, 0, 0, 0}})
	fmt.Println("Test 3:", res3)
	// Expected: [[1]]
}
```
