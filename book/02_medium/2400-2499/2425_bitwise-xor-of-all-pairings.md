# 2425 — Bitwise Xor Of All Pairings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func xorAllNums(nums1 []int, nums2 []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m + n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2425: Bitwise XOR of All Pairings
// https://leetcode.com/problems/bitwise-xor-of-all-pairings/
// Difficulty: Medium
// Time: O(m + n) | Space: O(1)
// XOR of all pairings = if len(nums2) odd, XOR all nums1; if len(nums1) odd, XOR all nums2.

import "fmt"

func main() {
	fmt.Println(xorAllNums([]int{2, 1, 3}, []int{10, 2, 5, 0})) // 13
	fmt.Println(xorAllNums([]int{1, 2}, []int{3, 4}))            // 0
}

func xorAllNums(nums1 []int, nums2 []int) int {
	ans := 0
	if len(nums2)%2 == 1 {
		for _, v := range nums1 {
			ans ^= v
		}
	}
	if len(nums1)%2 == 1 {
		for _, v := range nums2 {
			ans ^= v
		}
	}
	return ans
}
```
