# 2934 — Minimum Operations To Maximize Last Elements In Arrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minOperationsArrays(nums1 []int, nums2 []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2934: Minimum Operations to Maximize Last Elements in Arrays
// https://leetcode.com/problems/minimum-operations-to-maximize-last-elements-in-arrays/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(minOperationsArrays([]int{1, 2, 7}, []int{4, 5, 3}))
	fmt.Println(minOperationsArrays([]int{2, 3, 4, 5, 9}, []int{8, 7, 6, 1, 2}))
	fmt.Println(minOperationsArrays([]int{1, 5, 4}, []int{2, 5, 3}))
}

func minOperationsArrays(nums1 []int, nums2 []int) int {
	n := len(nums1)
	f := func(x, y int) (cnt int) {
		for i, a := range nums1[:n-1] {
			b := nums2[i]
			if a <= x && b <= y {
				continue
			}
			if !(a <= y && b <= x) {
				return -1
			}
			cnt++
		}
		return
	}
	a, b := f(nums1[n-1], nums2[n-1]), f(nums2[n-1], nums1[n-1])
	if a+b == -2 {
		return -1
	}
	if a < 0 {
		a = 1 << 30
	}
	if b < 0 {
		b = 1 << 30
	}
	if a < b+1 {
		return a
	}
	return b + 1
}
```
