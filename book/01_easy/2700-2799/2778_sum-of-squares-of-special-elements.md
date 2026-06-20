# 2778 — Sum Of Squares Of Special Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func SumOfSquaresOfSpecialElements(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2778: Sum of Squares of Special Elements
// https://leetcode.com/problems/sum-of-squares-of-special-elements/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(SumOfSquaresOfSpecialElements([]int{1, 2, 3, 4}))
	fmt.Println(SumOfSquaresOfSpecialElements([]int{2, 7, 1, 19, 18, 3}))
}

func SumOfSquaresOfSpecialElements(nums []int) int {
	n := len(nums)
	sum := 0
	for i, v := range nums {
		if n%(i+1) == 0 {
			sum += v * v
		}
	}
	return sum
}
```
