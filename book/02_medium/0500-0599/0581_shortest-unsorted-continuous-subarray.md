# 0581 — Shortest Unsorted Continuous Subarray

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindUnsortedSubarray(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #581: Shortest Unsorted Continuous Subarray
// https://leetcode.com/problems/shortest-unsorted-continuous-subarray/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(FindUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
	fmt.Println(FindUnsortedSubarray([]int{1, 2, 3, 4}))
	fmt.Println(FindUnsortedSubarray([]int{1}))
}

func FindUnsortedSubarray(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	left := -1
	minRight := nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] > minRight {
			left = i
		} else {
			minRight = nums[i]
		}
	}

	right := -1
	maxLeft := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] < maxLeft {
			right = i
		} else {
			maxLeft = nums[i]
		}
	}

	if right == -1 {
		return 0
	}
	return right - left + 1
}
```
