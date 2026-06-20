# 3131 — Find The Integer Added To Array I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheIntegerAddedToArrayI(nums1 []int, nums2 []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3131: Find the Integer Added to Array I
// https://leetcode.com/problems/find-the-integer-added-to-array-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: addedInteger
	fmt.Println(FindTheIntegerAddedToArrayI([]int{2, 6, 4}, []int{9, 7, 5})) // 3
	fmt.Println(FindTheIntegerAddedToArrayI([]int{10}, []int{5}))             // -5
}

// Time: O(n log n) | Space: O(1)
// LeetCode submission name: addedInteger
func FindTheIntegerAddedToArrayI(nums1 []int, nums2 []int) int {
	min1, min2 := nums1[0], nums2[0]
	for _, v := range nums1 {
		if v < min1 {
			min1 = v
		}
	}
	for _, v := range nums2 {
		if v < min2 {
			min2 = v
		}
	}
	return min2 - min1
}
```
