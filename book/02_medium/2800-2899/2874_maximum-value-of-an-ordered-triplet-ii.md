# 2874 — Maximum Value Of An Ordered Triplet Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumValueOfAnOrderedTripletIi(nums []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2874: Maximum Value of an Ordered Triplet II
// https://leetcode.com/problems/maximum-value-of-an-ordered-triplet-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func MaximumValueOfAnOrderedTripletIi(nums []int) int64 {
	n := len(nums)
	if n < 3 {
		return 0
	}

	var best int64
	maxDiff := int64(0) // max(nums[i] - nums[j]) for i < j
	maxVal := int64(nums[0])

	for k := 1; k < n; k++ {
		// Try nums[k] as the third element
		val := maxDiff * int64(nums[k])
		if val > best {
			best = val
		}

		// Update maxDiff with nums[k] as nums[j]
		diff := maxVal - int64(nums[k])
		if diff > maxDiff {
			maxDiff = diff
		}

		// Update maxVal
		if int64(nums[k]) > maxVal {
			maxVal = int64(nums[k])
		}
	}

	return best
}

func main() {
	fmt.Println(MaximumValueOfAnOrderedTripletIi([]int{12, 6, 1, 2, 7}))
	fmt.Println(MaximumValueOfAnOrderedTripletIi([]int{1, 10, 3, 4, 19}))
}
```
