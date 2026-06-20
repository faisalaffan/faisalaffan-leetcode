# 0915 — Partition Array Into Disjoint Intervals

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func PartitionArrayIntoDisjointIntervals(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #915: Partition Array into Disjoint Intervals
// https://leetcode.com/problems/partition-array-into-disjoint-intervals/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(PartitionArrayIntoDisjointIntervals([]int{5, 0, 3, 8, 6}))
	fmt.Println(PartitionArrayIntoDisjointIntervals([]int{1, 1, 1, 0, 6, 12}))
	fmt.Println(PartitionArrayIntoDisjointIntervals([]int{1, 1}))
}

// Time: O(n) | Space: O(1)
func PartitionArrayIntoDisjointIntervals(nums []int) int {
	leftMax, curMax, idx := nums[0], nums[0], 0

	for i, v := range nums {
		if v > curMax {
			curMax = v
		}
		if v < leftMax {
			leftMax = curMax
			idx = i
		}
	}

	return idx + 1
}
```
