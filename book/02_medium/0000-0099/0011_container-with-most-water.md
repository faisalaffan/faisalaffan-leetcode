# 0011 — Container With Most Water

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxArea(height []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #11: Container With Most Water
// https://leetcode.com/problems/container-with-most-water/
// Difficulty: Medium

import "fmt"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxWater := 0

  // Two-pointer: gerakkan kiri atau kanan
	for left < right {
		h := height[left]
		if height[right] < h {
			h = height[right]
		}
		area := h * (right - left)
		if area > maxWater {
			maxWater = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxWater
}

func main() {
	// Test case 1
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49

	// Test case 2
	fmt.Println(maxArea([]int{1, 1})) // 1

	// Test case 3
	fmt.Println(maxArea([]int{4, 3, 2, 1, 4})) // 16
}

// Time: O(n) | Space: O(1)
```
