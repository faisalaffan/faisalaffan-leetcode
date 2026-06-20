# 3105 — Longest Strictly Increasing Or Strictly Decreasing Subarray

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func LongestStrictlyIncreasingOrStrictlyDecreasingSubarray(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Monotonic Stack/Queue

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Monotonic Stack/Queue** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3105: Longest Strictly Increasing or Strictly Decreasing Subarray
// https://leetcode.com/problems/longest-strictly-increasing-or-strictly-decreasing-subarray/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: longestMonotonicSubarray
	fmt.Println(LongestStrictlyIncreasingOrStrictlyDecreasingSubarray([]int{1, 4, 3, 3, 2})) // 2
	fmt.Println(LongestStrictlyIncreasingOrStrictlyDecreasingSubarray([]int{3, 3, 3, 3}))  // 1
	fmt.Println(LongestStrictlyIncreasingOrStrictlyDecreasingSubarray([]int{3, 2, 1}))     // 3
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: longestMonotonicSubarray
func LongestStrictlyIncreasingOrStrictlyDecreasingSubarray(nums []int) int {
  // Edge case: input kosong — langsung return
	if len(nums) == 0 {
		return 0
	}
	inc := 1
	dec := 1
	maxLen := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			inc++
			dec = 1
		} else if nums[i] < nums[i-1] {
			dec++
			inc = 1
		} else {
			inc = 1
			dec = 1
		}
		if inc > maxLen {
			maxLen = inc
		}
		if dec > maxLen {
			maxLen = dec
		}
	}
	return maxLen
}
```
