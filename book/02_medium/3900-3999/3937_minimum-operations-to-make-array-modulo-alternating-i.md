# 3937 — Minimum Operations To Make Array Modulo Alternating I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumOperationsToMakeArrayModuloAlternatingI(nums []int, k int) int
```

> **💡 Hint:** Enumerate all (x,y) pairs with x != y, 0 <= x,y < k.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N * K^2)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3937: Minimum Operations to Make Array Modulo Alternating I
// https://leetcode.com/problems/minimum-operations-to-make-array-modulo-alternating-i/
// Difficulty: Medium
// Time: O(N * K^2) | Space: O(1)
// Approach: Enumerate all (x,y) pairs with x != y, 0 <= x,y < k.
// Even indices need modulo x, odd need modulo y.
// Cost per element = min(|curr - target|, k - |curr - target|).

import (
	"fmt"
)

func MinimumOperationsToMakeArrayModuloAlternatingI(nums []int, k int) int {
	n := len(nums)

	// Precompute remainders
  // Alokasi slice integer
	rem := make([]int, n)
	for i := 0; i < n; i++ {
		rem[i] = nums[i] % k
	}

	ans := -1
	for x := 0; x < k; x++ {
		for y := 0; y < k; y++ {
			if x == y {
				continue
			}
			cost := 0
			for i := 0; i < n; i++ {
				target := x
				if i%2 == 1 {
					target = y
				}
				diff := rem[i] - target
				if diff < 0 {
					diff = -diff
				}
				if diff > k-diff {
					diff = k - diff
				}
				cost += diff
			}
			if ans == -1 || cost < ans {
				ans = cost
			}
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(MinimumOperationsToMakeArrayModuloAlternatingI([]int{1, 4, 2, 8}, 3)) // Expected: 2

	// Example 2
	fmt.Println(MinimumOperationsToMakeArrayModuloAlternatingI([]int{1, 1, 1}, 3)) // Expected: 1
}
```
