# 0795 — Number Of Subarrays With Bounded Maximum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func numSubarrayBoundedMax(nums []int, left int, right int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #795: Number of Subarrays with Bounded Maximum
// https://leetcode.com/problems/number-of-subarrays-with-bounded-maximum/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(numSubarrayBoundedMax([]int{2, 1, 4, 3}, 2, 3))
	fmt.Println(numSubarrayBoundedMax([]int{2, 9, 2, 5, 6}, 2, 8))
}

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	count := 0
	prevCount := 0
	prevLessIdx := -1

	for i, val := range nums {
		if val >= left && val <= right {
			prevCount = i - prevLessIdx
			count += prevCount
		} else if val < left {
			count += prevCount
		} else {
			prevCount = 0
			prevLessIdx = i
		}
	}

	return count
}
```
