# 1567 — Maximum Length Of Subarray With Positive Product

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func GetMaxLen(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(N), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1567: Maximum Length of Subarray With Positive Product
// https://leetcode.com/problems/maximum-length-of-subarray-with-positive-product/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(GetMaxLen([]int{1, -2, -3, 4}))
	fmt.Println(GetMaxLen([]int{0, 1, -2, -3, -4}))
	fmt.Println(GetMaxLen([]int{-1, -2, -3, 0, 1}))
}

func GetMaxLen(nums []int) int {
	// Time: O(N), Space: O(1)
	// Track first occurrence of positive and negative prefix products
	maxLen := 0
	firstPos := -1
	firstNeg := -1
	prefix := 1 // 1 = positive, -1 = negative

	for i, num := range nums {
		if num > 0 {
			prefix = prefix // sign unchanged
		} else if num < 0 {
			prefix = -prefix
		} else {
			// Reset at zero
			prefix = 1
			firstPos = -1
			firstNeg = -1
			continue
		}

		if prefix == 1 {
			if firstPos == -1 {
				firstPos = i
			}
			if i-firstPos+1 > maxLen {
				maxLen = i - firstPos + 1
			}
		} else { // prefix == -1
			if firstNeg == -1 {
				firstNeg = i
			}
			if i-firstNeg+1 > maxLen {
				maxLen = i - firstNeg + 1
			}
		}
	}

	return maxLen
}
```
