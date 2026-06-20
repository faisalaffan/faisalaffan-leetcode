# 3806 — Maximum Bitwise And After Increment Operations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumAND(nums []int, k int, m int) int
```

> **💡 Hint:** For each bit position from high to low, check if we can

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Bitmask

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Bitmask** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3806: Maximum Bitwise AND After Increment Operations
// https://leetcode.com/problems/maximum-bitwise-and-after-increment-operations/
// Difficulty: Hard
//
// Given array nums, choose k elements and apply m total increments
// (distribute arbitrarily among chosen elements). Maximize the AND
// of all chosen elements after operations.
//
// Approach: For each bit position from high to low, check if we can
// set this bit in all chosen elements. If not, try to unset it.

import "fmt"

func main() {
	// Example 1
	fmt.Println(maximumAND([]int{3, 5, 1, 9}, 3, 4))
	// Example 2
	fmt.Println(maximumAND([]int{1, 2, 3}, 2, 1))
	// Edge: k = 1
	fmt.Println(maximumAND([]int{7, 15, 3}, 1, 0))
	// Edge: all same
	fmt.Println(maximumAND([]int{4, 4, 4}, 3, 5))
}

func maximumAND(nums []int, k int, m int) int {
	n := len(nums)
	if k > n {
		k = n
	}

	// Check if we can achieve AND >= target
	can := func(target int) bool {
		count := 0
		for _, v := range nums {
			if v&target == target {
				// This element already has all bits of target set
				count++
				if count >= k {
					return true
				}
				continue
			}
			// Need to increment
			need := 0
			for b := 0; b < 31; b++ {
				if target&(1<<b) != 0 && v&(1<<b) == 0 {
					// Need to set bit b: increment until this bit turns on
					// To set bit b, we need to add 2^b - (v & (2^b - 1))
					mask := (1 << uint(b)) - 1
					add := (1 << uint(b)) - (v & mask)
					if add > need {
						need = add
					}
				}
			}
			if need > 0 && need <= m {
				// We can increment this element to achieve target AND
				count++
				if count >= k {
					return true
				}
			}
		}
		return false
	}

	ans := 0
	for b := 30; b >= 0; b-- {
		candidate := ans | (1 << uint(b))
		if can(candidate) {
			ans = candidate
		}
	}

	return ans
}
```
