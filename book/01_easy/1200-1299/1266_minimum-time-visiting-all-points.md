# 1266 — Minimum Time Visiting All Points

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minTimeToVisitAllPoints(points [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1266: Minimum Time Visiting All Points
// https://leetcode.com/problems/minimum-time-visiting-all-points/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(minTimeToVisitAllPoints([][]int{{1, 1}, {3, 4}, {-1, 0}})) // 7
	fmt.Println(minTimeToVisitAllPoints([][]int{{3, 2}, {-2, 2}}))         // 5
}

// LeetCode submission: minTimeToVisitAllPoints
func minTimeToVisitAllPoints(points [][]int) int {
	ans := 0
	for i := 1; i < len(points); i++ {
		dx := points[i][0] - points[i-1][0]
		dy := points[i][1] - points[i-1][1]
		if dx < 0 {
			dx = -dx
		}
		if dy < 0 {
			dy = -dy
		}
		if dx > dy {
			ans += dx
		} else {
			ans += dy
		}
	}
	return ans
}
```
