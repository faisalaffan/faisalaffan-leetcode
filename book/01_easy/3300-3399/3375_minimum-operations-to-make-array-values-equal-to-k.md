# 3375 — Minimum Operations To Make Array Values Equal To K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumOperationsToMakeArrayValuesEqualToK(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3375: Minimum Operations to Make Array Values Equal to K
// https://leetcode.com/problems/minimum-operations-to-make-array-values-equal-to-k/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumOperationsToMakeArrayValuesEqualToK([]int{5, 2, 5, 4, 5}, 2))
	fmt.Println(MinimumOperationsToMakeArrayValuesEqualToK([]int{2, 1, 2}, 2))
	fmt.Println(MinimumOperationsToMakeArrayValuesEqualToK([]int{9, 7, 5, 3}, 1))
}

// MinimumOperationsToMakeArrayValuesEqualToK returns the minimum operations to reduce all numbers to k.
// In one operation, you can change any number > x to x for some x.
// Time: O(n). Space: O(n).
func MinimumOperationsToMakeArrayValuesEqualToK(nums []int, k int) int {
	minVal := nums[0]
	for _, v := range nums {
		if v < minVal {
			minVal = v
		}
	}
	if minVal < k {
		return -1
	}

  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[int]bool)
	for _, v := range nums {
		if v > k {
			seen[v] = true
		}
	}
	return len(seen)
}
```
