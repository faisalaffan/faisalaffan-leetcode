# 3422 — Minimum Operations To Make Subarray Elements Equal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minOperations(nums []int, k int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Sliding Window

**Kompleksitas Waktu:** O(n log k) Space: O(k)  
**Kompleksitas Ruang:** O(k)

> **Untuk fresh graduate:** Kuasai dulu teknik **Sliding Window** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3422: Minimum Operations to Make Subarray Elements Equal
// https://leetcode.com/problems/minimum-operations-to-make-subarray-elements-equal/
// Difficulty: Medium [Paid]
// Time: O(n log k) Space: O(k)

import (
	"fmt"
	"math"
	"sort"
)

func minOperations(nums []int, k int) int64 {
	n := len(nums)
	if k > n {
		return 0
	}

	var ans int64 = math.MaxInt64
  // Alokasi slice integer
	window := make([]int, k)

	for i := 0; i <= n-k; i++ {
		copy(window, nums[i:i+k])
  // Urutkan secara ascending — O(n log n)
		sort.Ints(window)
		median := window[k/2]

		var ops int64
		for _, v := range window {
			diff := v - median
			if diff < 0 {
				diff = -diff
			}
			ops += int64(diff)
		}
		if ops < ans {
			ans = ops
		}
	}
	return ans
}

func main() {
	fmt.Println(minOperations([]int{1, 4, 2, 6}, 3)) // 3
	fmt.Println(minOperations([]int{1, 2, 3, 4}, 2)) // 0
}
```
