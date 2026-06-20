# 0469 — Convex Polygon

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ConvexPolygon(polygon [][]int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #469: Convex Polygon
// https://leetcode.com/problems/convex-polygon/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(ConvexPolygon([][]int{{0, 0}, {0, 1}, {1, 1}, {1, 0}}))
	fmt.Println(ConvexPolygon([][]int{{0, 0}, {0, 10}, {10, 10}, {10, 0}, {5, 5}}))
}

func ConvexPolygon(polygon [][]int) bool {
	n := len(polygon)
	if n < 3 {
		return false
	}

	var prevCross int
	first := true

	for i := 0; i < n; i++ {
		a, b, c := polygon[i], polygon[(i+1)%n], polygon[(i+2)%n]
		cross := (b[0]-a[0])*(c[1]-b[1]) - (b[1]-a[1])*(c[0]-b[0])
		if cross != 0 {
			if first {
				prevCross = cross
				first = false
			} else if (cross > 0 && prevCross < 0) || (cross < 0 && prevCross > 0) {
				return false
			}
		}
	}
	return true
}
```
