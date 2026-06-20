# 3495 — Minimum Operations To Make Array Elements Zero

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minOperationsToZero(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3495: Minimum Operations to Make Array Elements Zero
// https://leetcode.com/problems/minimum-operations-to-make-array-elements-zero/
// Difficulty: Hard
//
// Each operation: pick a prefix [0..r-1] and power k, subtract 2^k from each element.
// For each bit k (0..60), consider the binary array b[i] = (nums[i]>>k)&1.
// Each operation clears bit k from a prefix, so we need 1 operation per
// contiguous group of 1s in the bit-k array. Total = sum over all bits of
// the number of groups of 1s.

import "fmt"

func minOperationsToZero(nums []int) int {
	ops := 0
	// Process each bit independently
	for k := 0; k <= 60; k++ {
		inGroup := false
		for _, x := range nums {
			hasBit := (x >> k) & 1
			if hasBit == 1 && !inGroup {
				ops++
				inGroup = true
			} else if hasBit == 0 {
				inGroup = false
			}
		}
	}
	return ops
}

func main() {
	// Test: nums=[1,2,3] -> expected 3
	// bit 0: [1,0,1] -> 2 groups
	// bit 1: [0,1,1] -> 1 group
	// total: 3
	fmt.Printf("[1,2,3] -> %d (expected 3)\n", minOperationsToZero([]int{1, 2, 3}))

	// Test: nums=[5] -> binary 101, bit 0: [1]->1, bit 2: [1]->1, total=2
	fmt.Printf("[5] -> %d (expected 2)\n", minOperationsToZero([]int{5}))

	// Test: nums=[2] -> binary 10, bit 1: [1]->1, total=1
	fmt.Printf("[2] -> %d (expected 1)\n", minOperationsToZero([]int{2}))

	// Test: nums=[0]
	fmt.Printf("[0] -> %d (expected 0)\n", minOperationsToZero([]int{0}))

	// Test: nums=[1,3,5,7] -> each bit group
	fmt.Printf("[1,3,5,7] -> %d\n", minOperationsToZero([]int{1, 3, 5, 7}))

	// Test: nums=[10,20,30]
	fmt.Printf("[10,20,30] -> %d\n", minOperationsToZero([]int{10, 20, 30}))
}
```
