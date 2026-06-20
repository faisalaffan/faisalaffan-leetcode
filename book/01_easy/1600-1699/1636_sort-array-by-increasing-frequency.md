# 1636 — Sort Array By Increasing Frequency

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FrequencySort(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1636: Sort Array by Increasing Frequency
// https://leetcode.com/problems/sort-array-by-increasing-frequency/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

// Time: O(n log n), Space: O(n)
func FrequencySort(nums []int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
  // Custom sort dengan comparator
	sort.Slice(nums, func(i, j int) bool {
		if freq[nums[i]] == freq[nums[j]] {
			return nums[i] > nums[j]
		}
		return freq[nums[i]] < freq[nums[j]]
	})
	return nums
}

func main() {
	fmt.Println(FrequencySort([]int{1, 1, 2, 2, 2, 3}))
	fmt.Println(FrequencySort([]int{2, 3, 1, 3, 2}))
	fmt.Println(FrequencySort([]int{-1, 1, -6, 4, 5, -6, 1, 4, 1}))
}
```
