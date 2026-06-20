# 0896 — Monotonic Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func isMonotonic(nums []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Monotonic Stack/Queue

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Kuasai dulu teknik **Monotonic Stack/Queue** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #896: Monotonic Array
// https://leetcode.com/problems/monotonic-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(isMonotonic([]int{1, 2, 2, 3}))   // true
	fmt.Println(isMonotonic([]int{6, 5, 4, 4}))   // true
	fmt.Println(isMonotonic([]int{1, 3, 2}))      // false
	fmt.Println(isMonotonic([]int{1, 1, 1}))      // true
}

// isMonotonic checks if the array is monotonic (either non-decreasing or non-increasing).
// Time: O(n). Space: O(1).
func isMonotonic(nums []int) bool {
	inc, dec := true, true
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dec = false
		}
		if nums[i] < nums[i-1] {
			inc = false
		}
	}
	return inc || dec
}
```
