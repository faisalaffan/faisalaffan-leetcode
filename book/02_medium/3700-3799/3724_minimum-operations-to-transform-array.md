# 3724 — Minimum Operations To Transform Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumOperationsToTransformArray(nums1 []int, nums2 []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3724: Minimum Operations to Transform Array
// https://leetcode.com/problems/minimum-operations-to-transform-array/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minimumOperationsToTransformArray(nums1 []int, nums2 []int) int64 {
	n := len(nums1)
	last := nums2[n]
	var ops int64 = 1
	var extra int64 = 1 << 60

	for i := 0; i < n; i++ {
		lo, hi := nums1[i], nums2[i]
		if lo > hi {
			lo, hi = hi, lo
		}
		ops += int64(hi - lo)

		if lo <= last && last <= hi {
			extra = 0
		} else if last < lo {
			if int64(1+lo-last) < extra {
				extra = int64(1 + lo - last)
			}
		} else {
			if int64(1+last-hi) < extra {
				extra = int64(1 + last - hi)
			}
		}
	}

	return ops + extra
}

func main() {
	fmt.Println(minimumOperationsToTransformArray([]int{2, 8}, []int{1, 7, 3}))
	fmt.Println(minimumOperationsToTransformArray([]int{1, 2}, []int{3, 4, 5}))
	fmt.Println(minimumOperationsToTransformArray([]int{5, 5}, []int{5, 5, 5}))
}
```
