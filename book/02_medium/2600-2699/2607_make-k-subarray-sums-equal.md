# 2607 — Make K Subarray Sums Equal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func makeSubKSumEqual(arr []int, k int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** GCD / Matematika

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **GCD / Matematika** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2607: Make K-Subarray Sums Equal
// https://leetcode.com/problems/make-k-subarray-sums-equal/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func makeSubKSumEqual(arr []int, k int) int64 {
	n := len(arr)
	g := gcd(n, k)
	var ans int64

	for i := 0; i < g; i++ {
		group := []int{}
		for j := i; j < n; j += g {
			group = append(group, arr[j])
		}
  // Urutkan secara ascending — O(n log n)
		sort.Ints(group)
		median := group[len(group)/2]
		for _, v := range group {
			diff := v - median
			if diff < 0 {
				diff = -diff
			}
			ans += int64(diff)
		}
	}
	return ans
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", makeSubKSumEqual([]int{1, 4, 1, 3}, 2))
	// Expected: 1

	// Test case 2
	fmt.Println("Test 2:", makeSubKSumEqual([]int{2, 5, 5, 7}, 3))
	// Expected: 5

	// Test case 3
	fmt.Println("Test 3:", makeSubKSumEqual([]int{1, 2, 3, 4, 5, 6}, 3))
	// Expected: 4
}
```
