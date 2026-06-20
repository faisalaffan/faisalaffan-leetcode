# 3070 — Count Submatrices With Top Left Element And Sum Less Than K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func countSubmatrices(grid [][]int, k int) (ans int)
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m*n)  
**Kompleksitas Ruang:** O(m*n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3070: Count Submatrices with Top-Left Element and Sum Less Than k
// https://leetcode.com/problems/count-submatrices-with-top-left-element-and-sum-less-than-k/
// Difficulty: Medium
// Time: O(m*n) | Space: O(m*n)

import "fmt"

func main() {
	fmt.Println(countSubmatrices([][]int{{7, 2, 9}, {1, 5, 0}, {2, 6, 6}}, 20))
	fmt.Println(countSubmatrices([][]int{{1, 2}, {3, 4}}, 5))
}

func countSubmatrices(grid [][]int, k int) (ans int) {
	m, n := len(grid), len(grid[0])
  // Membuat matriks/slice 2D untuk DP
	pref := make([][]int, m+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range pref {
		pref[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			pref[i+1][j+1] = grid[i][j] + pref[i][j+1] + pref[i+1][j] - pref[i][j]
			if pref[i+1][j+1] <= k {
				ans++
			}
		}
	}
	return
}
```
