# 3028 — Ant On The Boundary

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func AntOnTheBoundary(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3028: Ant on the Boundary
// https://leetcode.com/problems/ant-on-the-boundary/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: returnToBoundaryCount
	fmt.Println(AntOnTheBoundary([]int{2, 3, -5})) // 1
	fmt.Println(AntOnTheBoundary([]int{3, 2, -3, -2})) // 1
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: returnToBoundaryCount
func AntOnTheBoundary(nums []int) int {
	pos := 0
	count := 0
	for _, v := range nums {
		pos += v
		if pos == 0 {
			count++
		}
	}
	return count
}
```
