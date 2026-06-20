# 1287 — Element Appearing More Than 25 In Sorted Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findSpecialInteger(arr []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1287: Element Appearing More Than 25% In Sorted Array
// https://leetcode.com/problems/element-appearing-more-than-25-in-sorted-array/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(findSpecialInteger([]int{1, 2, 2, 6, 6, 6, 6, 7, 10})) // 6
	fmt.Println(findSpecialInteger([]int{1, 1}))                       // 1
}

// LeetCode submission: findSpecialInteger
func findSpecialInteger(arr []int) int {
	target := len(arr) / 4
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(arr)-target; i++ {
		if arr[i] == arr[i+target] {
			return arr[i]
		}
	}
	return arr[0]
}
```
