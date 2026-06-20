# 0661 — Image Smoother

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func imageSmoother(img [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m*n). Space: O(m*n).  
**Kompleksitas Ruang:** O(m*n).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #661: Image Smoother
// https://leetcode.com/problems/image-smoother/
// Difficulty: Easy

import "fmt"

func main() {
	img := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	fmt.Println(imageSmoother(img))
	// [[0,0,0],[0,0,0],[0,0,0]]

	img2 := [][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}}
	fmt.Println(imageSmoother(img2))
}

// imageSmoother applies a 3x3 smoother to each cell of the image.
// Time: O(m*n). Space: O(m*n).
func imageSmoother(img [][]int) [][]int {
	m, n := len(img), len(img[0])
  // Membuat matriks/slice 2D untuk DP
	result := make([][]int, m)
  // Range loop: iterasi dengan indeks + nilai
	for i := range result {
		result[i] = make([]int, n)
	}

	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 0}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum, count := 0, 0
			for _, d := range directions {
				ni, nj := i+d[0], j+d[1]
				if ni >= 0 && ni < m && nj >= 0 && nj < n {
					sum += img[ni][nj]
					count++
				}
			}
			result[i][j] = sum / count
		}
	}
	return result
}
```
