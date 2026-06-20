# 0977 — Squares Of A Sorted Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func sortedSquares(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #977: Squares of a Sorted Array
// https://leetcode.com/problems/squares-of-a-sorted-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10})) // [0,1,9,16,100]
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11})) // [4,9,9,49,121]
}

// sortedSquares returns squares of each number sorted in non-decreasing order.
// Time: O(n). Space: O(n).
func sortedSquares(nums []int) []int {
	n := len(nums)
  // Alokasi slice integer
	result := make([]int, n)
	l, r := 0, n-1
	pos := n - 1
	for l <= r {
		leftSq := nums[l] * nums[l]
		rightSq := nums[r] * nums[r]
		if leftSq > rightSq {
			result[pos] = leftSq
			l++
		} else {
			result[pos] = rightSq
			r--
		}
		pos--
	}
	return result
}
```
