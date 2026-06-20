# 0605 — Can Place Flowers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CanPlaceFlowers(flowerbed []int, n int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #605: Can Place Flowers
// https://leetcode.com/problems/can-place-flowers/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func CanPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(flowerbed) && count < n; i++ {
		if flowerbed[i] == 0 &&
			(i == 0 || flowerbed[i-1] == 0) &&
			(i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			count++
		}
	}
	return count >= n
}

func main() {
	fmt.Println(CanPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	fmt.Println(CanPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
	fmt.Println(CanPlaceFlowers([]int{0, 0, 1, 0, 0}, 1))
}
```
