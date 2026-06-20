# 3912 — Valid Elements In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ValidElementsInAnArray(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3912: Valid Elements in an Array
// https://leetcode.com/problems/valid-elements-in-an-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ValidElementsInAnArray([]int{1, 2, 4, 2, 3, 2}))
	fmt.Println(ValidElementsInAnArray([]int{5, 5, 5, 5}))
	fmt.Println(ValidElementsInAnArray([]int{1}))
}

// Time: O(n)
// Space: O(n)
func ValidElementsInAnArray(nums []int) []int {
	n := len(nums)
	isGreaterRight := make([]bool, n)

	maxRight := -1
	for i := n - 1; i >= 0; i-- {
		if nums[i] > maxRight {
			isGreaterRight[i] = true
			maxRight = nums[i]
		}
	}

	var result []int
	maxLeft := -1
	for i := 0; i < n; i++ {
		if nums[i] > maxLeft || isGreaterRight[i] {
			result = append(result, nums[i])
		}
		if nums[i] > maxLeft {
			maxLeft = nums[i]
		}
	}
	return result
}
```
