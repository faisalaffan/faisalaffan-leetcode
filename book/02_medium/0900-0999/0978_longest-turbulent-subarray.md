# 0978 — Longest Turbulent Subarray

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxTurbulenceSize(arr []int) int
```

> **💡 Hint:** Sliding window

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Sliding Window

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #978: Longest Turbulent Subarray
// https://leetcode.com/problems/longest-turbulent-subarray/
// Difficulty: Medium
//
// Approach: Sliding window
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(maxTurbulenceSize([]int{9, 4, 2, 10, 7, 8, 8, 1, 9})) // 5
	fmt.Println(maxTurbulenceSize([]int{4, 8, 12, 16}))                // 2
	fmt.Println(maxTurbulenceSize([]int{100}))                         // 1
}

func maxTurbulenceSize(arr []int) int {
	n := len(arr)
	if n < 2 {
		return n
	}

	left := 0
	result := 1

	for right := 1; right < n; right++ {
		cmp := sign(arr[right-1] - arr[right])
		if cmp == 0 {
			left = right
		} else if right == n-1 || cmp*sign(arr[right]-arr[right+1]) != -1 {
			if right-left+1 > result {
				result = right - left + 1
			}
			left = right
		}
	}

	return result
}

func sign(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}
```
