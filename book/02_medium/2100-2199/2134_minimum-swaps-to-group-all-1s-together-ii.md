# 2134 — Minimum Swaps To Group All 1S Together Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minSwaps(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Sliding Window

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Sliding Window** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2134: Minimum Swaps to Group All 1's Together II
// https://leetcode.com/problems/minimum-swaps-to-group-all-1s-together-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minSwaps(nums []int) int {
	n := len(nums)
	totalOnes := 0
	for _, v := range nums {
		if v == 1 {
			totalOnes++
		}
	}
	if totalOnes <= 1 {
		return 0
	}

	// Count zeros in first window of size totalOnes
	zeros := 0
	for i := 0; i < totalOnes; i++ {
		if nums[i] == 0 {
			zeros++
		}
	}
	minZeros := zeros

	// Slide window
	for i := totalOnes; i < n+totalOnes; i++ {
		// Remove element leaving window
		if nums[(i-totalOnes)%n] == 0 {
			zeros--
		}
		// Add element entering window
		if nums[i%n] == 0 {
			zeros++
		}
		if zeros < minZeros {
			minZeros = zeros
		}
	}

	return minZeros
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minSwaps([]int{0, 1, 0, 1, 1, 0, 0}))
	// Expected: 1

	// Test case 2
	fmt.Println("Test 2:", minSwaps([]int{0, 1, 1, 1, 0, 0, 1, 1, 0}))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", minSwaps([]int{1, 1, 0, 0, 1}))
	// Expected: 0
}
```
