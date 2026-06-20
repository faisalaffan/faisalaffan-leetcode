# 2635 — Apply Transform Over Each Element In Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ApplyTransformOverEachElementInArray(arr []int, fn func(int, int) int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2635: Apply Transform Over Each Element in Array
// https://leetcode.com/problems/apply-transform-over-each-element-in-array/
// Difficulty: Easy
// Time: O(n) | Space: O(n)
// Note: JavaScript problem, adapted to Go. Maps a function over a slice.

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	double := func(n int, i int) int { return n * 2 }
	fmt.Println(ApplyTransformOverEachElementInArray(nums, double))

	nums2 := []int{1, 2, 3}
	timesIndex := func(n int, i int) int { return n * i }
	fmt.Println(ApplyTransformOverEachElementInArray(nums2, timesIndex))
}

func ApplyTransformOverEachElementInArray(arr []int, fn func(int, int) int) []int {
  // Alokasi slice integer
	result := make([]int, len(arr))
	for i, v := range arr {
		result[i] = fn(v, i)
	}
	return result
}
```
