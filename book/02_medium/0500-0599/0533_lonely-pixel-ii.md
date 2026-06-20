# 0533 — Lonely Pixel Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindBlackPixel(picture [][]byte, target int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(m * n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #533: Lonely Pixel II
// https://leetcode.com/problems/lonely-pixel-ii/
// Difficulty: Medium [Paid]
// Time: O(m * n)
// Space: O(m * n)

import "fmt"

func main() {
	picture := [][]byte{
		{'W', 'B', 'W', 'B', 'B', 'W'},
		{'W', 'B', 'W', 'B', 'B', 'W'},
		{'W', 'B', 'W', 'B', 'B', 'W'},
		{'W', 'W', 'B', 'W', 'B', 'W'},
	}
	fmt.Println(FindBlackPixel(picture, 3))
}

func FindBlackPixel(picture [][]byte, target int) int {
	m, n := len(picture), len(picture[0])
  // Alokasi slice integer
	rows := make([]int, m)
  // Alokasi slice integer
	cols := make([]int, n)
  // Membuat map (HashMap) — pencarian O(1)
	rowPattern := make(map[string]int)

	for i := 0; i < m; i++ {
		rowStr := ""
		for j := 0; j < n; j++ {
			rowStr += string(picture[i][j])
			if picture[i][j] == 'B' {
				rows[i]++
				cols[j]++
			}
		}
		rowPattern[rowStr]++
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if picture[i][j] == 'B' && rows[i] == target && cols[j] == target {
				// All rows with 'B' in column j must be identical
				allSame := true
				for r := 0; r < m; r++ {
					if picture[r][j] == 'B' && picture[r][j] == picture[i][j] {
						// Check if rows i and r are identical
						for c := 0; c < n; c++ {
							if picture[i][c] != picture[r][c] {
								allSame = false
								break
							}
						}
						if !allSame {
							break
						}
					}
				}
				if allSame {
					count++
				}
			}
		}
	}

	return count
}
```
