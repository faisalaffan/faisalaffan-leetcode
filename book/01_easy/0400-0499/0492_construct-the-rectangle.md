# 0492 — Construct The Rectangle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ConstructTheRectangle(area int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(sqrt(n)), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #492: Construct the Rectangle
// https://leetcode.com/problems/construct-the-rectangle/
// Difficulty: Easy

import "fmt"

// Time: O(sqrt(n)), Space: O(1)
func ConstructTheRectangle(area int) []int {
	w := 1
	for i := 1; i*i <= area; i++ {
		if area%i == 0 {
			w = i
		}
	}
	return []int{area / w, w}
}

func main() {
	fmt.Println(ConstructTheRectangle(4))
	fmt.Println(ConstructTheRectangle(37))
	fmt.Println(ConstructTheRectangle(122122))
}
```
