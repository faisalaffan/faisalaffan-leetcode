# 0992 — Subarrays With K Different Integers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func subarraysWithKDistinct(nums []int, k int) int
```

> **💡 Hint:** atMostK trick.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer, Sliding Window

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #992: Subarrays with K Different Integers
// https://leetcode.com/problems/subarrays-with-k-different-integers/
// Difficulty: Hard
//
// Approach: atMostK trick.
//   subarraysWithKDistinct(nums, k) = atMostK(nums, k) - atMostK(nums, k-1)
//   atMostK counts subarrays with <= K distinct integers using a sliding window.

import "fmt"

func main() {
	fmt.Println(subarraysWithKDistinct([]int{1, 2, 1, 2, 3}, 2)) // 7
	fmt.Println(subarraysWithKDistinct([]int{1, 2, 1, 3, 4}, 3)) // 3
}

func subarraysWithKDistinct(nums []int, k int) int {
	return atMostK(nums, k) - atMostK(nums, k-1)
}

func atMostK(nums []int, k int) int {
	if k == 0 {
		return 0
	}
  // Membuat map (HashMap) — pencarian O(1)
	count := make(map[int]int)
	left, result := 0, 0
	for right := 0; right < len(nums); right++ {
		count[nums[right]]++
		for len(count) > k {
			count[nums[left]]--
			if count[nums[left]] == 0 {
				delete(count, nums[left])
			}
			left++
		}
		result += right - left + 1
	}
	return result
}
```
