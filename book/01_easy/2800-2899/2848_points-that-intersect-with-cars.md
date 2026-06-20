# 2848 — Points That Intersect With Cars

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func PointsThatIntersectWithCars(nums [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * range)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2848: Points That Intersect With Cars
// https://leetcode.com/problems/points-that-intersect-with-cars/
// Difficulty: Easy
// Time: O(n * range) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(PointsThatIntersectWithCars([][]int{{3, 6}, {1, 5}, {4, 7}}))
	fmt.Println(PointsThatIntersectWithCars([][]int{{1, 3}, {5, 8}}))
}

func PointsThatIntersectWithCars(nums [][]int) int {
  // Alokasi slice integer
	points := make([]bool, 101)
	for _, car := range nums {
		for p := car[0]; p <= car[1]; p++ {
			points[p] = true
		}
	}
	count := 0
	for _, v := range points {
		if v {
			count++
		}
	}
	return count
}
```
