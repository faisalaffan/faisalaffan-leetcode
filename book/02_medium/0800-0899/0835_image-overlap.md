# 0835 — Image Overlap

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ImageOverlap(img1 [][]int, img2 [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^4)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #835: Image Overlap
// https://leetcode.com/problems/image-overlap/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(ImageOverlap([][]int{{1, 1, 0}, {0, 1, 0}, {0, 1, 0}}, [][]int{{0, 0, 0}, {0, 1, 1}, {0, 0, 1}}))
	fmt.Println(ImageOverlap([][]int{{1}}, [][]int{{1}}))
	fmt.Println(ImageOverlap([][]int{{0}}, [][]int{{0}}))
}

// Time: O(n^4) | Space: O(1)
func ImageOverlap(img1 [][]int, img2 [][]int) int {
	n := len(img1)
	ans := 0

	for a := 1 - n; a < n; a++ {
		for b := 1 - n; b < n; b++ {
			count := 0
			for i := max(a, 0); i < min(n, n+a); i++ {
				for j := max(b, 0); j < min(n, n+b); j++ {
					if img2[i][j] == 1 && img1[i-a][j-b] == 1 {
						count++
					}
				}
			}
			if count > ans {
				ans = count
			}
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
