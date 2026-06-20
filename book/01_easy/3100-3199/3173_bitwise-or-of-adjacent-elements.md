# 3173 — Bitwise Or Of Adjacent Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func BitwiseOrOfAdjacentElements(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3173: Bitwise OR of Adjacent Elements
// https://leetcode.com/problems/bitwise-or-of-adjacent-elements/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	// LeetCode name: orArray
	fmt.Println(BitwiseOrOfAdjacentElements([]int{1, 2, 3, 4})) // [3, 3, 7]
	fmt.Println(BitwiseOrOfAdjacentElements([]int{5, 1, 6}))     // [5, 7]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: orArray
func BitwiseOrOfAdjacentElements(nums []int) []int {
  // Alokasi slice integer
	result := make([]int, len(nums)-1)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums)-1; i++ {
		result[i] = nums[i] | nums[i+1]
	}
	return result
}
```
