# 1464 — Maximum Product Of Two Elements In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxProduct(nums []int) int

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1464: Maximum Product of Two Elements in an Array
// https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/
// Difficulty: Easy
//
// LeetCode submission: func maxProduct(nums []int) int

import "fmt"

func main() {
	fmt.Println(MaximumProductOfTwoElementsInAnArray([]int{3, 4, 5, 2})) // 12
	fmt.Println(MaximumProductOfTwoElementsInAnArray([]int{1, 5, 4, 5})) // 16
	fmt.Println(MaximumProductOfTwoElementsInAnArray([]int{3, 7}))       // 12
}

// Time: O(n), Space: O(1)
func MaximumProductOfTwoElementsInAnArray(nums []int) int {
	first, second := 0, 0
	for _, v := range nums {
		if v > first {
			second = first
			first = v
		} else if v > second {
			second = v
		}
	}
	return (first - 1) * (second - 1)
}
```
