# 1121 — Divide Array Into Increasing Sequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func canDivideIntoIncreasingSequences(nums []int, k int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1121: Divide Array Into Increasing Sequences
// https://leetcode.com/problems/divide-array-into-increasing-sequences/
// Difficulty: Hard [Paid]
//
// Given a sorted integer array nums and an integer k, determine if nums can
// be divided into exactly k increasing sequences. Since nums is sorted and
// each sequence must be strictly increasing, equal elements must go to
// different sequences. Therefore, we need at least as many sequences as the
// maximum frequency of any element.

import "fmt"

func main() {
	// Example test case
	fmt.Println(canDivideIntoIncreasingSequences([]int{1, 2, 2, 3, 3, 4, 4}, 3)) // true

	// Single element
	fmt.Println(canDivideIntoIncreasingSequences([]int{5}, 1)) // true

	// Insufficient sequences
	fmt.Println(canDivideIntoIncreasingSequences([]int{1, 1, 1, 1}, 1)) // false (maxFreq=4 > k=1)

	// All distinct
	fmt.Println(canDivideIntoIncreasingSequences([]int{1, 2, 3, 4, 5}, 1)) // true

	// Empty array (len(nums) >= 1 per constraints, but handle gracefully)
	fmt.Println(canDivideIntoIncreasingSequences([]int{}, 2)) // true (vacuously)

	// Exact match
	fmt.Println(canDivideIntoIncreasingSequences([]int{1, 1, 2, 2, 3, 3}, 2)) // true
}

// canDivideIntoIncreasingSequences returns true if nums (sorted ascending) can
// be divided into exactly k strictly increasing sequences.
//
// Key insight: Because the array is sorted, equal values cannot share the same
// increasing sequence. The maximum frequency of any value determines the
// minimum number of sequences required. If maxFreq <= k, it is always possible.
//
// Proof: Greedily assign each element in order to sequences 0..k-1 in a
// round-robin fashion. Since we never assign equal elements to the same
// sequence (they are spaced at least k apart) and elements are processed in
// sorted order, each sequence is strictly increasing.
func canDivideIntoIncreasingSequences(nums []int, k int) bool {
  // Edge case: input kosong — langsung return
	if len(nums) == 0 {
		return true
	}
	if k == 0 {
		return false
	}

	maxFreq := 1
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
			if count > maxFreq {
				maxFreq = count
			}
		} else {
			count = 1
		}
	}

	return maxFreq <= k
}
```
