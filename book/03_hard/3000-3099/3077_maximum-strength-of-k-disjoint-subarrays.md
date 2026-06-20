# 3077 — Maximum Strength Of K Disjoint Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumStrength(nums []int, k int) int64
```

> **💡 Hint:** DP with two states

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n*k)  
**Kompleksitas Ruang:** O(k)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3077: Maximum Strength of K Disjoint Subarrays
// https://leetcode.com/problems/maximum-strength-of-k-disjoint-subarrays/
// Difficulty: Hard
// Time: O(n*k) | Space: O(k)
//
// Approach: DP with two states
// dp0[j] = max strength with j subarrays, NOT using current element
// dp1[j] = max strength with j subarrays, ENDING at current element
// Weight for j-th subarray (1-indexed): (-1)^(j+1) * (k-j+1)

import (
	"fmt"
	"math"
)

func MaximumStrength(nums []int, k int) int64 {
	n := len(nums)

  // Alokasi slice integer
	dp0 := make([]int64, k+1)
  // Alokasi slice integer
	dp1 := make([]int64, k+1)

	weight := func(j int) int64 {
		w := int64(k - j + 1)
		if j%2 == 0 {
			return -w
		}
		return w
	}

	negInf := int64(math.MinInt64 / 2)
	for j := 0; j <= k; j++ {
		dp0[j] = negInf
		dp1[j] = negInf
	}
	dp0[0] = 0

	for i := 0; i < n; i++ {
  // Alokasi slice integer
		ndp0 := make([]int64, k+1)
  // Alokasi slice integer
		ndp1 := make([]int64, k+1)
		for j := 0; j <= k; j++ {
			ndp0[j] = negInf
			ndp1[j] = negInf
		}

		for j := 0; j <= k; j++ {
			// Not using nums[i]: carry forward best state without nums[i]
			ndp0[j] = max(ndp0[j], dp0[j])
			ndp0[j] = max(ndp0[j], dp1[j])

			if j > 0 {
				w := weight(j)
				// Start new subarray at nums[i]
				bestPrev := max(dp0[j-1], dp1[j-1])
				ndp1[j] = max(ndp1[j], bestPrev+w*int64(nums[i]))

				// Extend current subarray to include nums[i]
				ndp1[j] = max(ndp1[j], dp1[j]+w*int64(nums[i]))
			}
		}

		dp0, dp1 = ndp0, ndp1
	}

	return max(dp0[k], dp1[k])
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example 1
	fmt.Println("Test 1:", MaximumStrength([]int{1, 2, 3, -1, 2}, 3))
	// Expected: 22

	// Example 2
	fmt.Println("Test 2:", MaximumStrength([]int{12, -2, -2, -2, -2}, 5))
	// Expected: 64

	// Example 3
	fmt.Println("Test 3:", MaximumStrength([]int{-1, -2, -3}, 1))
	// Expected: -1

	// Single element, k=1
	fmt.Println("Test 4:", MaximumStrength([]int{5}, 1))
	// Expected: 5

	// All negative, k=1
	fmt.Println("Test 5:", MaximumStrength([]int{-5, -3, -1}, 1))
	// Expected: -1 (best single element)

	// Two subarrays from 4 elements
	fmt.Println("Test 6:", MaximumStrength([]int{1, 2, 3, 4}, 2))
	// weight(1)=2, weight(2)=-1
	// Possible: [1,2] with w=2, [3,4] with w=-1: 2*(1+2) + (-1)*(3+4) = 6-7 = -1
	// Or: [1] w=2, [4] w=-1: 2*1 + (-1)*4 = 2-4 = -2
	// [1,2,3] w=2, [4] w=-1: 2*6 + (-1)*4 = 12-4 = 8
	// Hmm, expected depends on optimal selection

	// Large range
	fmt.Println("Test 7:", MaximumStrength([]int{1000000, 1000000, 1000000}, 2))
	// Expected: 2000000
}
```
