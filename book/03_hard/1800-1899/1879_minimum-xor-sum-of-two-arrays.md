# 1879 — Minimum Xor Sum Of Two Arrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumXorSum(nums1 []int, nums2 []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Bitmask

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1879: Minimum XOR Sum of Two Arrays
// https://leetcode.com/problems/minimum-xor-sum-of-two-arrays/
// Difficulty: Hard

import (
	"fmt"
	"math"
)

func minimumXorSum(nums1 []int, nums2 []int) int {
	n := len(nums1)
	size := 1 << n
  // Alokasi slice integer
	dp := make([]int, size)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	// dp[mask] = min XOR sum assigning first k elements of nums1
	// to elements of nums2 indicated by mask (where k = popcount(mask))
	for mask := 0; mask < size; mask++ {
		i := bitsCount(mask) // number of assigned elements in nums1
		if i >= n {
			continue
		}
		for j := 0; j < n; j++ {
			if mask&(1<<j) == 0 {
				newMask := mask | (1 << j)
				val := dp[mask] + (nums1[i] ^ nums2[j])
				if val < dp[newMask] {
					dp[newMask] = val
				}
			}
		}
	}
	return dp[size-1]
}

func bitsCount(x int) int {
	c := 0
	for x > 0 {
		c += x & 1
		x >>= 1
	}
	return c
}

func main() {
	// Example: [1,2], [2,3] -> 2
	// (1 XOR 2) + (2 XOR 3) = 3 + 1 = 4
	// (1 XOR 3) + (2 XOR 2) = 2 + 0 = 2 (minimum)
	fmt.Println(minimumXorSum([]int{1, 2}, []int{2, 3}))

	// Additional test
	fmt.Println(minimumXorSum([]int{1, 0, 3}, []int{5, 3, 4}))
}
```
