# 1300 — Sum Of Mutated Array Closest To Target

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findBestValue(arr []int, target int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n log n + n log max(arr))  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1300: Sum of Mutated Array Closest to Target
// https://leetcode.com/problems/sum-of-mutated-array-closest-to-target/
// Difficulty: Medium

// Find integer value such that sum(arr[i] if arr[i] < value else value)
// is as close to target as possible. If tie, return smaller value.

// Time: O(n log n + n log max(arr))
// Space: O(1)

func findBestValue(arr []int, target int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(arr)
	n := len(arr)

  // Alokasi slice integer
	prefix := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] + arr[i-1]
	}

	lo, hi := 0, arr[n-1]
	bestVal, minDiff := 0, target

	for lo <= hi {
		mid := lo + (hi-lo)/2

		// Find first index > mid
		idx := sort.Search(n, func(i int) bool {
			return arr[i] > mid
		})

		sum := prefix[idx] + mid*(n-idx)
		diff := abs(sum - target)

		if diff < minDiff || (diff == minDiff && mid < bestVal) {
			bestVal = mid
			minDiff = diff
		}

		if sum < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return bestVal
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Printf("%d (expected: 3)\n", findBestValue([]int{4, 9, 3}, 10))
	fmt.Printf("%d (expected: 5)\n", findBestValue([]int{2, 3, 5}, 10))
	fmt.Printf("%d (expected: 11361)\n", findBestValue([]int{60864, 25176, 27249, 21296, 20204}, 56803))
}
```
