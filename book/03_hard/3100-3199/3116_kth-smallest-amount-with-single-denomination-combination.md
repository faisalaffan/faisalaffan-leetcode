# 3116 — Kth Smallest Amount With Single Denomination Combination

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func lcmSafe(a, b, limit int64) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search, DFS, GCD / Matematika

**Kompleksitas Waktu:** O(2^m * log(k * min_coin)) where m = filtered coin count  
**Kompleksitas Ruang:** O(m)

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3116: Kth Smallest Amount With Single Denomination Combination
// https://leetcode.com/problems/kth-smallest-amount-with-single-denomination-combination/
// Difficulty: Hard
// Time: O(2^m * log(k * min_coin)) where m = filtered coin count
// Space: O(m)
//
// Find the k-th smallest amount that can be represented as a positive multiple of
// at least one coin denomination. Use inclusion-exclusion with LCM and binary search.

import (
	"fmt"
	"sort"
)

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcmSafe(a, b, limit int64) int64 {
	g := gcd(a, b)
	aDivG := a / g
	if aDivG > limit/b {
		return limit + 1
	}
	return aDivG * b
}

func kthSmallestAmount(coins []int, k int) int64 {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(coins)
  // Alokasi slice integer
	filtered := make([]int, 0)
	for _, c := range coins {
		redundant := false
		for _, f := range filtered {
			if c%f == 0 {
				redundant = true
				break
			}
		}
		if !redundant {
			filtered = append(filtered, c)
		}
	}

	m := len(filtered)
  // Alokasi slice integer
	coinI64 := make([]int64, m)
	for i, c := range filtered {
		coinI64[i] = int64(c)
	}

	count := func(X int64) int64 {
		var dfs func(idx int, curLCM int64, cnt int) int64
		dfs = func(idx int, curLCM int64, cnt int) int64 {
			if idx == m {
				if cnt == 0 {
					return 0
				}
				if cnt%2 == 1 {
					return X / curLCM
				}
				return -(X / curLCM)
			}
			total := dfs(idx+1, curLCM, cnt)
			newLCM := lcmSafe(curLCM, coinI64[idx], X)
			if newLCM <= X {
				total += dfs(idx+1, newLCM, cnt+1)
			}
			return total
		}
		return dfs(0, 1, 0)
	}

	minCoin := int64(filtered[0])
	low := int64(1)
	high := minCoin * int64(k)

	for low < high {
		mid := low + (high-low)/2
		if count(mid) >= int64(k) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", kthSmallestAmount([]int{3, 6, 9}, 3))
	// Expected: 9

	// Test case 2
	fmt.Println("Test 2:", kthSmallestAmount([]int{5, 2}, 7))
	// Expected: 12

	// Test case 3
	fmt.Println("Test 3:", kthSmallestAmount([]int{2, 3, 4}, 5))
	// Expected: 8

	// Test case 4: single coin
	fmt.Println("Test 4:", kthSmallestAmount([]int{5}, 4))
	// Expected: 20
}
```
