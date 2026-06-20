# 2541 — Minimum Operations To Make Array Equal Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minOperations(nums1 []int, nums2 []int, k int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2541: Minimum Operations to Make Array Equal II
// https://leetcode.com/problems/minimum-operations-to-make-array-equal-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minOperations(nums1 []int, nums2 []int, k int) int64 {
	n := len(nums1)
	if k == 0 {
		for i := 0; i < n; i++ {
			if nums1[i] != nums2[i] {
				return -1
			}
		}
		return 0
	}

	var posDiff, negDiff int64
	for i := 0; i < n; i++ {
		diff := nums1[i] - nums2[i]
		if diff%k != 0 {
			return -1
		}
		if diff > 0 {
			posDiff += int64(diff)
		} else if diff < 0 {
			negDiff += int64(-diff)
		}
	}

	if posDiff != negDiff {
		return -1
	}
	return posDiff / int64(k)
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minOperations([]int{4, 3, 1, 4}, []int{1, 3, 7, 1}, 3))
	// Expected: 2

	// Test case 2: valid case
	fmt.Println("Test 2:", minOperations([]int{5, 10, 3}, []int{8, 6, 4}, 1))
	// Expected: 4

	// Test case 3: impossible (sums don't match or diff not divisible)
	fmt.Println("Test 3:", minOperations([]int{1, 2}, []int{2, 1}, 3))
	// Expected: -1
}
```
