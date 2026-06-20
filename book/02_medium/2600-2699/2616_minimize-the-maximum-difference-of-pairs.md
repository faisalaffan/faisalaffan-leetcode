# 2616 — Minimize The Maximum Difference Of Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimizeMax(nums []int, p int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n log n + n log maxDiff)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2616: Minimize the Maximum Difference of Pairs
// https://leetcode.com/problems/minimize-the-maximum-difference-of-pairs/
// Difficulty: Medium
// Time: O(n log n + n log maxDiff) | Space: O(1)

import (
	"fmt"
	"sort"
)

func minimizeMax(nums []int, p int) int {
	if p == 0 {
		return 0
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	n := len(nums)

	canForm := func(mid int) bool {
		count := 0
		i := 0
		for i < n-1 {
			if nums[i+1]-nums[i] <= mid {
				count++
				i += 2
			} else {
				i++
			}
			if count >= p {
				return true
			}
		}
		return false
	}

	left, right := 0, nums[n-1]-nums[0]
  // Two-pointer: gerakkan kiri atau kanan
	for left < right {
		mid := left + (right-left)/2
		if canForm(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minimizeMax([]int{10, 1, 2, 7, 1, 3}, 2))
	// Expected: 1

	// Test case 2
	fmt.Println("Test 2:", minimizeMax([]int{4, 2, 1, 2}, 1))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", minimizeMax([]int{3, 5, 2, 8, 1, 9}, 3))
	// Expected: 2
}
```
