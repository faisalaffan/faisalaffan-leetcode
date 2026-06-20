# 2799 — Count Complete Subarrays In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CountCompleteSubarraysInAnArray(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2799: Count Complete Subarrays in an Array
// https://leetcode.com/problems/count-complete-subarrays-in-an-array/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func CountCompleteSubarraysInAnArray(nums []int) int {
	// Count distinct elements
  // Membuat map (HashMap) — pencarian O(1)
	distinct := make(map[int]bool)
	for _, n := range nums {
		distinct[n] = true
	}
	target := len(distinct)

	left := 0
	count := 0
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	unique := 0

	for right := 0; right < len(nums); right++ {
		freq[nums[right]]++
		if freq[nums[right]] == 1 {
			unique++
		}

		for unique == target {
			// All subarrays from left to right-end are valid
			count += len(nums) - right
			freq[nums[left]]--
			if freq[nums[left]] == 0 {
				unique--
			}
			left++
		}
	}

	return count
}

func main() {
	fmt.Println(CountCompleteSubarraysInAnArray([]int{1, 3, 1, 2, 2}))
	fmt.Println(CountCompleteSubarraysInAnArray([]int{1, 1}))
}
```
