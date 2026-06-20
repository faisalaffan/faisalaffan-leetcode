# 3818 — Minimum Prefix Removal To Make Array Strictly Increasing

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumPrefixRemovalToMakeArrayStrictlyIncreasing(nums []int) int
```

> **💡 Hint:** Scan from right to left to find the longest strictly increasing suffix.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Prefix Sum

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3818: Minimum Prefix Removal to Make Array Strictly Increasing
// https://leetcode.com/problems/minimum-prefix-removal-to-make-array-strictly-increasing/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: Scan from right to left to find the longest strictly increasing suffix.

import "fmt"

func MinimumPrefixRemovalToMakeArrayStrictlyIncreasing(nums []int) int {
	n := len(nums)
	for i := n - 1; i > 0; i-- {
		if nums[i-1] >= nums[i] {
			return i
		}
	}
	return 0
}

func main() {
	// Example 1
	fmt.Println(MinimumPrefixRemovalToMakeArrayStrictlyIncreasing([]int{1, -1, 2, 3, 3, 4, 5})) // Expected: 4

	// Example 2
	fmt.Println(MinimumPrefixRemovalToMakeArrayStrictlyIncreasing([]int{4, 3, -2, -5})) // Expected: 3

	// Example 3
	fmt.Println(MinimumPrefixRemovalToMakeArrayStrictlyIncreasing([]int{1, 2, 3, 4})) // Expected: 0
}
```
