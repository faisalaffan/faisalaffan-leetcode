# 3381 — Maximum Subarray Sum With Length Divisible By K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxSubarraySum(nums []int, k int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n) Space: O(k)  
**Kompleksitas Ruang:** O(k)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3381: Maximum Subarray Sum With Length Divisible by K
// https://leetcode.com/problems/maximum-subarray-sum-with-length-divisible-by-k/
// Difficulty: Medium
// Time: O(n) Space: O(k)

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubarraySum([]int{1, 2, 3, 4, 5, 6}, 2)) // 21
	fmt.Println(maxSubarraySum([]int{-1, -2, -3, -4, -5}, 3)) // -6
}

func maxSubarraySum(nums []int, k int) int64 {
	n := len(nums)
  // Alokasi slice integer
	prefix := make([]int64, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + int64(nums[i])
	}

  // Alokasi slice integer
	minPref := make([]int64, k)
  // Range loop: iterasi dengan indeks + nilai
	for i := range minPref {
		minPref[i] = math.MaxInt64
	}
	minPref[0] = 0 // prefix[0] = 0

	var ans int64 = math.MinInt64

	for i := 1; i <= n; i++ {
		r := i % k
		if minPref[r] != math.MaxInt64 {
			val := prefix[i] - minPref[r]
			if val > ans {
				ans = val
			}
		}
		if prefix[i] < minPref[r] {
			minPref[r] = prefix[i]
		}
	}

	if ans == math.MinInt64 {
		return 0
	}
	return ans
}
```
