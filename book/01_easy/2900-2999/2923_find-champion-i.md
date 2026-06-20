# 2923 — Find Champion I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindChampionI(grid [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2923: Find Champion I
// https://leetcode.com/problems/find-champion-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: findChampion
	fmt.Println(FindChampionI([][]int{{0, 1}, {0, 0}}))          // 0
	fmt.Println(FindChampionI([][]int{{0, 0, 1}, {1, 0, 1}, {0, 0, 0}})) // 1
}

// Time: O(n^2) | Space: O(1)
// LeetCode submission name: findChampion
func FindChampionI(grid [][]int) int {
	n := len(grid)
	for i := 0; i < n; i++ {
		isChampion := true
		for j := 0; j < n; j++ {
			if i != j && grid[i][j] != 1 {
				isChampion = false
				break
			}
		}
		if isChampion {
			return i
		}
	}
	return -1
}
```
