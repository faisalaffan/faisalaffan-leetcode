# 1991 — Find The Middle Index In Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheMiddleIndexInArray(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1991: Find the Middle Index in Array
// https://leetcode.com/problems/find-the-middle-index-in-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindTheMiddleIndexInArray([]int{2, 3, -1, 8, 4}))   // 3
	fmt.Println(FindTheMiddleIndexInArray([]int{1, -1, 4}))          // 2
	fmt.Println(FindTheMiddleIndexInArray([]int{2, 5}))              // -1
}

// Time: O(n), Space: O(1)
func FindTheMiddleIndexInArray(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	leftSum := 0
	for i, v := range nums {
		if leftSum == total-leftSum-v {
			return i
		}
		leftSum += v
	}
	return -1
}
```
