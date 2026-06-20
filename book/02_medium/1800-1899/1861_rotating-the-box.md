# 1861 — Rotating The Box

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func RotateTheBox(box [][]byte) [][]byte
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m*n), Space: O(m*n) for result  
**Kompleksitas Ruang:** O(m*n) for result

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1861: Rotating the Box
// https://leetcode.com/problems/rotating-the-box/
// Difficulty: Medium

import "fmt"

func main() {
	box1 := [][]byte{{'#', '.', '#'}}
	fmt.Println(RotateTheBox(box1))

	box2 := [][]byte{{'#', '.', '*', '.'}, {'#', '#', '*', '.'}}
	fmt.Println(RotateTheBox(box2))

	box3 := [][]byte{{'#', '#', '*', '.', '*', '.'},
		{'#', '#', '#', '*', '.', '.'},
		{'#', '#', '#', '.', '#', '.'}}
	fmt.Println(RotateTheBox(box3))
}

// Time: O(m*n), Space: O(m*n) for result
func RotateTheBox(box [][]byte) [][]byte {
	m, n := len(box), len(box[0])

	// Apply gravity to each row (stones fall to the right)
	for i := 0; i < m; i++ {
		empty := n - 1
		for j := n - 1; j >= 0; j-- {
			if box[i][j] == '*' {
				empty = j - 1
			} else if box[i][j] == '#' {
				box[i][j] = '.'
				box[i][empty] = '#'
				empty--
			}
		}
	}

	// Rotate 90 degrees clockwise
  // Membuat matriks/slice 2D untuk DP
	result := make([][]byte, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range result {
		result[i] = make([]byte, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result[j][m-1-i] = box[i][j]
		}
	}
	return result
}
```
