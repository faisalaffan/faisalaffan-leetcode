# 3748 — Count Stable Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func countStableSubarrays(nums []int, queries [][]int) []int64
```

> **💡 Hint:** Precompute longest non-decreasing run ending at each position.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3748: Count Stable Subarrays
// https://leetcode.com/problems/count-stable-subarrays/
// Difficulty: Hard
//
// A subarray is stable (non-decreasing) if it contains no inversion pairs
// (nums[i] > nums[j] when i < j). Given range queries, count stable
// subarrays within each [l, r] range.
//
// Approach: Precompute longest non-decreasing run ending at each position.
// Use prefix sums for fast range query responses.

import "fmt"

func main() {
	// Example 1
	fmt.Println(countStableSubarrays([]int{1, 2, 2, 1}, [][]int{{0, 3}, {0, 1}}))
	// Example 2
	fmt.Println(countStableSubarrays([]int{1, 1, 1}, [][]int{{0, 2}}))
	// Edge: single element
	fmt.Println(countStableSubarrays([]int{5}, [][]int{{0, 0}}))
	// Edge: all decreasing
	fmt.Println(countStableSubarrays([]int{3, 2, 1}, [][]int{{0, 2}}))
}

func countStableSubarrays(nums []int, queries [][]int) []int64 {
	n := len(nums)

	// lenEnd[i] = length of longest non-decreasing subarray ending at i
  // Alokasi slice integer
	lenEnd := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 || nums[i] < nums[i-1] {
			lenEnd[i] = 1
		} else {
			lenEnd[i] = lenEnd[i-1] + 1
		}
	}

	// pref[i] = total non-decreasing subarrays in nums[0..i] (inclusive)
  // Alokasi slice integer
	pref := make([]int64, n)
	for i := 0; i < n; i++ {
		pref[i] = int64(lenEnd[i])
		if i > 0 {
			pref[i] += pref[i-1]
		}
	}

  // Alokasi slice integer
	ans := make([]int64, len(queries))
	for qi, q := range queries {
		l, r := q[0], q[1]
		if l > r {
			ans[qi] = 0
			continue
		}

		// Count all non-decreasing subarrays in [l, r]:
		// Total = pref[r] - pref[l-1], but need to subtract subarrays
		// that start before l and end at or after l.
		// Those subarrays start at positions < l and end at l, l+1, ..., r.
		// For each position i in [l, r], we know lenEnd[i] is the max length
		// ending at i. Some of those subarrays start before l.
		// The number of subarrays ending at i that start before l is:
		// max(0, lenEnd[i] - (i - l + 1))
		// = max(0, lenEnd[i] - (i-l+1))
		// = max(0, lenEnd[i] - i + l - 1)

		total := pref[r]
		if l > 0 {
			total -= pref[l-1]
		}

		// Subtract subarrays that start before l
		for i := l; i <= r; i++ {
			overflow := lenEnd[i] - (i - l + 1)
			if overflow > 0 {
				total -= int64(overflow)
			}
		}

		ans[qi] = total
	}

	return ans
}
```
