# 1992 — Find All Groups Of Farmland

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindAllGroupsOfFarmland(land [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m*n), Space: O(1) (excluding output)  
**Kompleksitas Ruang:** O(1) (excluding output)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1992: Find All Groups of Farmland
// https://leetcode.com/problems/find-all-groups-of-farmland/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FindAllGroupsOfFarmland([][]int{{1, 0, 0}, {0, 1, 1}, {0, 1, 1}}))
	fmt.Println(FindAllGroupsOfFarmland([][]int{{1, 1}, {1, 1}}))
	fmt.Println(FindAllGroupsOfFarmland([][]int{{0}}))
}

// Time: O(m*n), Space: O(1) (excluding output)
func FindAllGroupsOfFarmland(land [][]int) [][]int {
	m, n := len(land), len(land[0])
  // Membuat matriks/slice 2D untuk DP
	result := make([][]int, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if land[i][j] == 1 {
				r, c := i, j
				for r+1 < m && land[r+1][j] == 1 {
					r++
				}
				for c+1 < n && land[i][c+1] == 1 {
					c++
				}
				result = append(result, []int{i, j, r, c})
				for x := i; x <= r; x++ {
					for y := j; y <= c; y++ {
						land[x][y] = 0
					}
				}
			}
		}
	}

	return result
}
```
