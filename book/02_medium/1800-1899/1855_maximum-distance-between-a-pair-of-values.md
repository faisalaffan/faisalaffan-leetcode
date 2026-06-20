# 1855 — Maximum Distance Between A Pair Of Values

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaxDistance(nums1 []int, nums2 []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m+n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1855: Maximum Distance Between a Pair of Values
// https://leetcode.com/problems/maximum-distance-between-a-pair-of-values/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaxDistance([]int{55, 30, 5, 4, 2}, []int{100, 20, 10, 10, 5}))
	fmt.Println(MaxDistance([]int{2, 2, 2}, []int{10, 10, 1}))
	fmt.Println(MaxDistance([]int{30, 29, 19, 5}, []int{25, 25, 25, 25, 25}))
}

// Time: O(m+n), Space: O(1)
func MaxDistance(nums1 []int, nums2 []int) int {
	i, j := 0, 0
	maxDist := 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			if j-i > maxDist {
				maxDist = j - i
			}
			j++
		} else {
			i++
		}
	}
	return maxDist
}
```
