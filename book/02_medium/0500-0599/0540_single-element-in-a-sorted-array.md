# 0540 — Single Element In A Sorted Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func SingleNonDuplicate(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #540: Single Element in a Sorted Array
// https://leetcode.com/problems/single-element-in-a-sorted-array/
// Difficulty: Medium
// Time: O(log n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(SingleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
	fmt.Println(SingleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11}))
}

func SingleNonDuplicate(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		// Ensure mid is even to compare with mid+1
		if mid%2 == 1 {
			mid--
		}
		if nums[mid] == nums[mid+1] {
			lo = mid + 2
		} else {
			hi = mid
		}
	}
	return nums[lo]
}
```
