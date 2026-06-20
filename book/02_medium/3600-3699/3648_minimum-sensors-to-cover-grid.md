# 3648 — Minimum Sensors To Cover Grid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumSensorsToCoverGrid(n int, m int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3648: Minimum Sensors to Cover Grid
// https://leetcode.com/problems/minimum-sensors-to-cover-grid/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func minimumSensorsToCoverGrid(n int, m int, k int) int {
	s := 2*k + 1
	rows := (n + s - 1) / s // ceil division
	cols := (m + s - 1) / s
	return rows * cols
}

func main() {
	fmt.Println(minimumSensorsToCoverGrid(5, 5, 1))
	fmt.Println(minimumSensorsToCoverGrid(2, 2, 2))
	fmt.Println(minimumSensorsToCoverGrid(3, 4, 0))
}
```
