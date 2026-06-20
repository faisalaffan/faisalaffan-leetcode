# 0755 — Pour Water

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func pourWater(heights []int, volume int, k int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(V * N)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #755: Pour Water
// https://leetcode.com/problems/pour-water/
// Difficulty: Medium [Paid]
// Time: O(V * N)
// Space: O(1)

import "fmt"

func main() {
	heights := []int{2, 1, 1, 2, 1, 2, 2}
	result := pourWater(heights, 4, 3)
	fmt.Println(result)
}

func pourWater(heights []int, volume int, k int) []int {
	n := len(heights)

	for v := 0; v < volume; v++ {
		pos := k

		// Try left
		left := k
		for left > 0 && heights[left] >= heights[left-1] {
			left--
		}
		for left < k && heights[left] == heights[left+1] {
			left++
		}
		if heights[left] < heights[pos] {
			pos = left
		}

		if pos == k {
			// Try right
			right := k
			for right < n-1 && heights[right] >= heights[right+1] {
				right++
			}
			for right > k && heights[right] == heights[right-1] {
				right--
			}
			if heights[right] < heights[pos] {
				pos = right
			}
		}

		heights[pos]++
	}

	return heights
}
```
