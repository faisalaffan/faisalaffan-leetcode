# 0189 — Rotate Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func rotate(nums []int, k int) 
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #189: Rotate Array
// https://leetcode.com/problems/rotate-array/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func rotate(nums []int, k int) {
	n := len(nums)
  // Edge case: input kosong — langsung return
	if n == 0 {
		return
	}
	k = k % n
	if k == 0 {
		return
	}

	reverse := func(arr []int, l, r int) {
  // Two-pointer: gerakkan kiri atau kanan
		for l < r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums1, 3)
	fmt.Println(nums1)

	nums2 := []int{-1, -100, 3, 99}
	rotate(nums2, 2)
	fmt.Println(nums2)

	nums3 := []int{1}
	rotate(nums3, 0)
	fmt.Println(nums3)
}
```
