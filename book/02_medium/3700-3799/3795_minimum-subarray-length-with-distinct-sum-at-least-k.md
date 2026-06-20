# 3795 — Minimum Subarray Length With Distinct Sum At Least K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumSubarrayLengthWithDistinctSumAtLeastK(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3795: Minimum Subarray Length With Distinct Sum At Least K
// https://leetcode.com/problems/minimum-subarray-length-with-distinct-sum-at-least-k/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"math"
)

func minimumSubarrayLengthWithDistinctSumAtLeastK(nums []int, k int) int {
	n := len(nums)
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	left := 0
	distinctSum := 0
	ans := math.MaxInt32

	for right := 0; right < n; right++ {
		freq[nums[right]]++
		if freq[nums[right]] == 1 {
			distinctSum += nums[right]
		}

		for distinctSum >= k {
			length := right - left + 1
			if length < ans {
				ans = length
			}
			freq[nums[left]]--
			if freq[nums[left]] == 0 {
				distinctSum -= nums[left]
			}
			left++
		}
	}

	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func main() {
	fmt.Println(minimumSubarrayLengthWithDistinctSumAtLeastK([]int{2, 2, 3, 1}, 4))
	fmt.Println(minimumSubarrayLengthWithDistinctSumAtLeastK([]int{3, 2, 3, 4}, 5))
	fmt.Println(minimumSubarrayLengthWithDistinctSumAtLeastK([]int{5, 5, 4}, 5))
}
```
