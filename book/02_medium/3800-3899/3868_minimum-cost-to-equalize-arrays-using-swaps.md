# 3868 — Minimum Cost To Equalize Arrays Using Swaps

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumCostToEqualizeArraysUsingSwaps(nums1 []int, nums2 []int) int
```

> **💡 Hint:** Count frequencies in both arrays. If any value's total count is odd,

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3868: Minimum Cost to Equalize Arrays Using Swaps
// https://leetcode.com/problems/minimum-cost-to-equalize-arrays-using-swaps/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Count frequencies in both arrays. If any value's total count is odd,
// impossible (return -1). Min cost = sum of positive differences / 2.

import "fmt"

func MinimumCostToEqualizeArraysUsingSwaps(nums1 []int, nums2 []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	cnt := make(map[int]int)
	for _, v := range nums1 {
		cnt[v]++
	}
	for _, v := range nums2 {
		cnt[v]--
	}

	posDiff := 0
	for _, c := range cnt {
		if c%2 != 0 {
			return -1
		}
		if c > 0 {
			posDiff += c
		}
	}
	return posDiff / 2
}

func main() {
	// Example 1
	fmt.Println(MinimumCostToEqualizeArraysUsingSwaps([]int{10, 20}, []int{20, 10})) // Expected: 0

	// Example 2
	fmt.Println(MinimumCostToEqualizeArraysUsingSwaps([]int{10, 10}, []int{20, 20})) // Expected: 1

	// Example 3
	fmt.Println(MinimumCostToEqualizeArraysUsingSwaps([]int{10, 20}, []int{30, 40})) // Expected: -1
}
```
