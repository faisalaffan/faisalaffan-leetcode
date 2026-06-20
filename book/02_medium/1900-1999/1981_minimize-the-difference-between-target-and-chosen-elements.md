# 1981 — Minimize The Difference Between Target And Chosen Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimizeTheDifference(mat [][]int, target int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m * n * maxSum), Space: O(maxSum) where maxSum = 70*70 = 4900  
**Kompleksitas Ruang:** O(maxSum) where maxSum = 70*70 = 4900

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1981: Minimize the Difference Between Target and Chosen Elements
// https://leetcode.com/problems/minimize-the-difference-between-target-and-chosen-elements/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinimizeTheDifference([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 13))
	fmt.Println(MinimizeTheDifference([][]int{{1}, {2}, {3}}, 100))
	fmt.Println(MinimizeTheDifference([][]int{{1, 2, 9, 8, 7}}, 6))
}

// Time: O(m * n * maxSum), Space: O(maxSum) where maxSum = 70*70 = 4900
func MinimizeTheDifference(mat [][]int, target int) int {
	m, n := len(mat), len(mat[0])
	possible := make([]bool, 4901)
	possible[0] = true

	for i := 0; i < m; i++ {
		next := make([]bool, 4901)
		for s := 0; s < len(possible); s++ {
			if possible[s] {
				for j := 0; j < n; j++ {
					if s+mat[i][j] < len(next) {
						next[s+mat[i][j]] = true
					}
				}
			}
		}
		possible = next
	}

	ans := 1 << 30
	for s := 0; s < len(possible); s++ {
		if possible[s] {
			diff := s - target
			if diff < 0 {
				diff = -diff
			}
			if diff < ans {
				ans = diff
			}
		}
	}
	return ans
}
```
