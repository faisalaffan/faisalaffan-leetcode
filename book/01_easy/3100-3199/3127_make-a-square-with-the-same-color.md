# 3127 — Make A Square With The Same Color

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MakeASquareWithTheSameColor(grid [][]byte) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3127: Make a Square with the Same Color
// https://leetcode.com/problems/make-a-square-with-the-same-color/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: canMakeSquare
	grid1 := [][]byte{{'B', 'W', 'B'}, {'B', 'W', 'W'}, {'B', 'W', 'B'}}
	fmt.Println(MakeASquareWithTheSameColor(grid1)) // true

	grid2 := [][]byte{{'B', 'W', 'B'}, {'W', 'B', 'W'}, {'B', 'W', 'B'}}
	fmt.Println(MakeASquareWithTheSameColor(grid2)) // false

	grid3 := [][]byte{{'B', 'W', 'B'}, {'B', 'W', 'W'}, {'B', 'W', 'W'}}
	fmt.Println(MakeASquareWithTheSameColor(grid3)) // true
}

// Time: O(1) | Space: O(1)
// LeetCode submission name: canMakeSquare
func MakeASquareWithTheSameColor(grid [][]byte) bool {
	// Check all 2x2 subgrids
	dirs := [][2]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			bCount := 0
			for _, d := range dirs {
				if grid[i+d[0]][j+d[1]] == 'B' {
					bCount++
				}
			}
			// If at least 3 cells have the same color, changing 1 makes all 4
			if bCount >= 3 || bCount <= 1 {
				return true
			}
		}
	}
	return false
}
```
