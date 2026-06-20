# 1838 — Frequency Of The Most Frequent Element

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxFrequency(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Sliding Window

**Kompleksitas Waktu:** O(n log n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1838: Frequency of the Most Frequent Element
// https://leetcode.com/problems/frequency-of-the-most-frequent-element/
// Difficulty: Medium
// Time: O(n log n), Space: O(1)

import (
	"fmt"
	"sort"
)

func maxFrequency(nums []int, k int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	left := 0
	total := 0
	maxFreq := 0

	for right := 0; right < len(nums); right++ {
		total += nums[right]

		// Shrink window if we can't make all elements in window equal
		for nums[right]*(right-left+1)-total > k {
			total -= nums[left]
			left++
		}

		if right-left+1 > maxFreq {
			maxFreq = right - left + 1
		}
	}
	return maxFreq
}

func main() {
	fmt.Println(maxFrequency([]int{1, 2, 4}, 5))       // Expected: 3
	fmt.Println(maxFrequency([]int{1, 4, 8, 13}, 5))   // Expected: 2
	fmt.Println(maxFrequency([]int{3, 9, 6}, 2))       // Expected: 1
}
```
