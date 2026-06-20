# 1423 — Maximum Points You Can Obtain From Cards

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxScore(cardPoints []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(k) where k = number of cards to take  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1423: Maximum Points You Can Obtain from Cards
// https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(maxScore([]int{1, 2, 3, 4, 5, 6, 1}, 3)) // 12

	// Test case 2
	fmt.Println(maxScore([]int{2, 2, 2}, 2)) // 4

	// Test case 3
	fmt.Println(maxScore([]int{9, 7, 7, 9, 7, 7, 9}, 7)) // 55

	// Test case 4
	fmt.Println(maxScore([]int{1, 1000, 1}, 1)) // 1
}

// Time: O(k) where k = number of cards to take
// Space: O(1)
func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)

	// Sum first k elements (taking from left)
	sum := 0
	for i := 0; i < k; i++ {
		sum += cardPoints[i]
	}

	maxSum := sum
	// Try taking from right instead
	for i := 0; i < k; i++ {
		// Remove one from left, add one from right
		sum = sum - cardPoints[k-1-i] + cardPoints[n-1-i]
		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}
```
