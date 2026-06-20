# 2875 — Minimum Size Subarray In Infinite Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumSizeSubarrayInInfiniteArray(nums []int, target int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2875: Minimum Size Subarray in Infinite Array
// https://leetcode.com/problems/minimum-size-subarray-in-infinite-array/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"math"
)

func MinimumSizeSubarrayInInfiniteArray(nums []int, target int) int {
	n := len(nums)
	var totalSum int
	for _, v := range nums {
		totalSum += v
	}

	// If target is 0, we need empty subarray
	if target == 0 {
		return 0
	}

	repeats := target / totalSum
	remainder := target % totalSum

	if remainder == 0 {
		return repeats * n
	}

	// Find minimum subarray with sum == remainder in doubled array
  // Alokasi slice integer
	extended := make([]int, n*2)
	copy(extended, nums)
	copy(extended[n:], nums)

	best := math.MaxInt32
	left := 0
	sum := 0

	for right := 0; right < len(extended); right++ {
		sum += extended[right]
		for sum > remainder && left <= right {
			sum -= extended[left]
			left++
		}
		if sum == remainder {
			length := right - left + 1
			if length < best {
				best = length
			}
		}
	}

	if best == math.MaxInt32 {
		return -1
	}

	return repeats*n + best
}

func main() {
	fmt.Println(MinimumSizeSubarrayInInfiniteArray([]int{1, 2, 3}, 5))
	fmt.Println(MinimumSizeSubarrayInInfiniteArray([]int{1, 1, 1}, 4))
}
```
