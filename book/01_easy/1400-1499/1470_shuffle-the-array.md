# 1470 — Shuffle The Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func shuffle(nums []int, n int) []int

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1470: Shuffle the Array
// https://leetcode.com/problems/shuffle-the-array/
// Difficulty: Easy
//
// LeetCode submission: func shuffle(nums []int, n int) []int

import "fmt"

func main() {
	fmt.Println(ShuffleTheArray([]int{2, 5, 1, 3, 4, 7}, 3)) // [2 3 5 4 1 7]
	fmt.Println(ShuffleTheArray([]int{1, 2, 3, 4, 4, 3, 2, 1}, 4)) // [1 4 2 3 3 2 4 1]
}

// Time: O(n), Space: O(n)
func ShuffleTheArray(nums []int, n int) []int {
  // Alokasi slice integer
	res := make([]int, 2*n)
	for i := 0; i < n; i++ {
		res[2*i] = nums[i]
		res[2*i+1] = nums[i+n]
	}
	return res
}
```
