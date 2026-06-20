# 1480 — Running Sum Of 1D Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func runningSum(nums []int) []int

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1) excluding output  
**Kompleksitas Ruang:** O(1) excluding output

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1480: Running Sum of 1d Array
// https://leetcode.com/problems/running-sum-of-1d-array/
// Difficulty: Easy
//
// LeetCode submission: func runningSum(nums []int) []int

import "fmt"

func main() {
	fmt.Println(RunningSumOfOneDArray([]int{1, 2, 3, 4}))    // [1 3 6 10]
	fmt.Println(RunningSumOfOneDArray([]int{1, 1, 1, 1, 1})) // [1 2 3 4 5]
}

// Time: O(n), Space: O(1) excluding output
func RunningSumOfOneDArray(nums []int) []int {
  // Alokasi slice integer
	res := make([]int, len(nums))
	sum := 0
	for i, v := range nums {
		sum += v
		res[i] = sum
	}
	return res
}
```
