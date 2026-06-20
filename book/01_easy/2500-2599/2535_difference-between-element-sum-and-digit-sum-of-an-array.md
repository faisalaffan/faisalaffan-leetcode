# 2535 — Difference Between Element Sum And Digit Sum Of An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func DifferenceBetweenElementSumAndDigitSumOfAnArray(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2535: Difference Between Element Sum and Digit Sum of an Array
// https://leetcode.com/problems/difference-between-element-sum-and-digit-sum-of-an-array/
// Difficulty: Easy
// Time O(n log m) | Space O(1)

import "fmt"

func main() {
	fmt.Println(DifferenceBetweenElementSumAndDigitSumOfAnArray([]int{1, 15, 6, 3})) // 9
	fmt.Println(DifferenceBetweenElementSumAndDigitSumOfAnArray([]int{1, 2, 3, 4}))  // 0
}

func DifferenceBetweenElementSumAndDigitSumOfAnArray(nums []int) int {
	elementSum := 0
	digitSum := 0
	for _, n := range nums {
		elementSum += n
		for n > 0 {
			digitSum += n % 10
			n /= 10
		}
	}
	result := elementSum - digitSum
	if result < 0 {
		return -result
	}
	return result
}
```
