# 0080 — Remove Duplicates From Sorted Array Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func removeDuplicates(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #80: Remove Duplicates from Sorted Array II
// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
// Difficulty: Medium

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	write := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[write-2] {
			nums[write] = nums[i]
			write++
		}
	}

	return write
}

func main() {
	// Test case 1
	nums1 := []int{1, 1, 1, 2, 2, 3}
	k1 := removeDuplicates(nums1)
	fmt.Println(nums1[:k1]) // [1 1 2 2 3]

	// Test case 2
	nums2 := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	k2 := removeDuplicates(nums2)
	fmt.Println(nums2[:k2]) // [0 0 1 1 2 3 3]
}

// Time: O(n) | Space: O(1)
```
