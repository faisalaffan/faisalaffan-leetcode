# 2401 — Longest Nice Subarray

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func longestNiceSubarray(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Sliding Window

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2401: Longest Nice Subarray
// https://leetcode.com/problems/longest-nice-subarray/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Sliding window: maintain OR of current window. If new element conflicts, shrink left.

import "fmt"

func main() {
	fmt.Println(longestNiceSubarray([]int{1, 3, 8, 48, 10})) // 3
	fmt.Println(longestNiceSubarray([]int{3, 1, 5, 11, 13}))  // 1
}

func longestNiceSubarray(nums []int) int {
	left, orMask, ans := 0, 0, 0
	for right, v := range nums {
		for orMask&v != 0 {
			orMask ^= nums[left]
			left++
		}
		orMask |= v
		if right-left+1 > ans {
			ans = right - left + 1
		}
	}
	return ans
}
```
