# 3036 — Number Of Subarrays That Match A Pattern Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func countMatchingSubarrays(nums, pattern []int) int
```

> **💡 Hint:** Z-algorithm (linear time)

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3036: Number of Subarrays That Match a Pattern II
// https://leetcode.com/problems/number-of-subarrays-that-match-a-pattern-ii/
// Difficulty: Hard
//
// Given an array nums and a pattern array where each element is -1, 0, or 1,
// count the number of subarrays of nums of length len(pattern)+1 that match
// the pattern. A subarray matches if for each adjacent pair in the subarray,
// the comparison result (nums[i+1] - nums[i] sign) equals the pattern value.
//
// Approach: Z-algorithm (linear time)
//   Build a combined array: pattern + [-2] + (nums[i+1] cmp nums[i] for i in range).
//   Use Z-algorithm to find all positions where the pattern appears.

import (
	"cmp"
	"fmt"
)

func countMatchingSubarrays(nums, pattern []int) int {
	m := len(pattern)

	// Build combined array: pattern | sentinel | diff array
  // Alokasi slice integer
	arr := make([]int, 0, m+1+len(nums)-1)
	arr = append(arr, pattern...)
	arr = append(arr, 2) // sentinel (any value not in {-1,0,1})
	for i := 1; i < len(nums); i++ {
		arr = append(arr, cmp.Compare(nums[i], nums[i-1]))
	}

	n := len(arr)

	// Z-algorithm
  // Alokasi slice integer
	z := make([]int, n)
	l, r := 0, 0
	for i := 1; i < n; i++ {
		if i <= r {
			z[i] = min2(r-i+1, z[i-l])
		}
		for i+z[i] < n && arr[z[i]] == arr[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > r {
			l, r = i, i+z[i]-1
		}
	}

	// Count matches
	ans := 0
	for i := m + 1; i < n; i++ {
		if z[i] == m {
			ans++
		}
	}
	return ans
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Example: [1,2,3,4,5,6], pattern [1,1]
	// Comparisons: 1,1,1,1,1 -> matches at [1,2,3], [2,3,4], [3,4,5], [4,5,6] -> 4
	fmt.Println("Test 1:", countMatchingSubarrays([]int{1, 2, 3, 4, 5, 6}, []int{1, 1}))

	// Example: [1,4,4,1,3,5,5,3], pattern [1,0,-1]
	// comparisons: 1,0,-1,1,1,0,-1
	// matches at [1,4,4,1] -> 1
	fmt.Println("Test 2:", countMatchingSubarrays([]int{1, 4, 4, 1, 3, 5, 5, 3}, []int{1, 0, -1}))

	// All equal
	fmt.Println("Test 3:", countMatchingSubarrays([]int{5, 5, 5, 5}, []int{0, 0}))

	// Decreasing
	fmt.Println("Test 4:", countMatchingSubarrays([]int{5, 4, 3, 2, 1}, []int{-1, -1}))

	// No match
	fmt.Println("Test 5:", countMatchingSubarrays([]int{1, 2, 3}, []int{-1}))

	// Single pattern element
	fmt.Println("Test 6:", countMatchingSubarrays([]int{1, 2, 1, 2}, []int{1}))
}
```
