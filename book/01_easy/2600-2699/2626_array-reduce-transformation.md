# 2626 — Array Reduce Transformation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func arrayReduceTransformation(nums []int, fn func(int, int) int, init int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2626: Array Reduce Transformation
// https://leetcode.com/problems/array-reduce-transformation/
// Difficulty: Easy
// Time: O(n) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Reduces slice using a function.

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	sum := func(acc, curr int) int { return acc + curr }
	fmt.Println(arrayReduceTransformation(nums, sum, 0))

	nums2 := []int{1, 2, 3, 4}
	product := func(acc, curr int) int { return acc * curr }
	fmt.Println(arrayReduceTransformation(nums2, product, 1))
}

func arrayReduceTransformation(nums []int, fn func(int, int) int, init int) int {
	result := init
	for _, num := range nums {
		result = fn(result, num)
	}
	return result
}
```
