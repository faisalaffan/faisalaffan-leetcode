# 1064 — Fixed Point

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func fixedPoint(arr []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1064: Fixed Point
// https://leetcode.com/problems/fixed-point/
// Difficulty: Easy [Paid]
// Time: O(log n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(fixedPoint([]int{-10, -5, 0, 3, 7}))  // 3
	fmt.Println(fixedPoint([]int{0, 2, 5, 8, 17}))    // 0
	fmt.Println(fixedPoint([]int{-10, -5, 3, 4, 7, 9})) // -1
}

// LeetCode submission: fixedPoint
func fixedPoint(arr []int) int {
	lo, hi := 0, len(arr)-1
	for lo < hi {
		mid := (lo + hi) >> 1
		if arr[mid] >= mid {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	if arr[lo] == lo {
		return lo
	}
	return -1
}
```
