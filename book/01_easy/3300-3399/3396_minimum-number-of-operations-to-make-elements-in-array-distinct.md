# 3396 — Minimum Number Of Operations To Make Elements In Array Distinct

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumNumberOfOperationsToMakeElementsInArrayDistinct(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3396: Minimum Number of Operations to Make Elements in Array Distinct
// https://leetcode.com/problems/minimum-number-of-operations-to-make-elements-in-array-distinct/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumNumberOfOperationsToMakeElementsInArrayDistinct([]int{1, 2, 3, 4, 2, 3, 3, 5, 7}))
	fmt.Println(MinimumNumberOfOperationsToMakeElementsInArrayDistinct([]int{4, 5, 6, 4, 4}))
	fmt.Println(MinimumNumberOfOperationsToMakeElementsInArrayDistinct([]int{6, 7, 8, 9}))
}

// MinimumNumberOfOperationsToMakeElementsInArrayDistinct returns minimum operations to make all elements distinct.
// Each operation removes the first 3 elements (or all remaining if < 3).
// Time: O(n). Space: O(n).
func MinimumNumberOfOperationsToMakeElementsInArrayDistinct(nums []int) int {
	ops := 0
	for {
  // Membuat map (HashMap) — pencarian O(1)
		seen := make(map[int]bool)
		distinct := true
		for _, v := range nums {
			if seen[v] {
				distinct = false
				break
			}
			seen[v] = true
		}
		if distinct || len(nums) == 0 {
			return ops
		}
		// Remove first 3 elements
		if len(nums) <= 3 {
			nums = []int{}
		} else {
			nums = nums[3:]
		}
		ops++
	}
}
```
