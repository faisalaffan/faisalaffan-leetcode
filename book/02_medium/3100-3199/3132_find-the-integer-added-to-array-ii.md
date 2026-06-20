# 3132 — Find The Integer Added To Array Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumAddedInteger(nums1 []int, nums2 []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3132: Find the Integer Added to Array II
// https://leetcode.com/problems/find-the-integer-added-to-array-ii/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import (
	"fmt"
	"sort"
)

func minimumAddedInteger(nums1 []int, nums2 []int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums1)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums2)

	// Try all pairs from nums1 as the two removed elements
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums1); i++ {
		for j := i + 1; j < len(nums1); j++ {
			// Check if nums2 can be matched after removing nums1[i] and nums1[j]
			diff := -1001
			idx := 0
			match := true
			for k := 0; k < len(nums1) && match; k++ {
				if k == i || k == j {
					continue
				}
				curDiff := nums2[idx] - nums1[k]
				if diff == -1001 {
					diff = curDiff
				} else if curDiff != diff {
					match = false
				}
				idx++
			}
			if match && diff >= 0 {
				return diff
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(minimumAddedInteger([]int{4, 20, 16, 12, 8}, []int{14, 18, 10})) // Expected: -2
	fmt.Println(minimumAddedInteger([]int{3, 5, 5, 3}, []int{7, 7}))             // Expected: 2
}
```
