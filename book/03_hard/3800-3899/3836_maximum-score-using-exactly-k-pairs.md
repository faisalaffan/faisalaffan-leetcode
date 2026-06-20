# 3836 — Maximum Score Using Exactly K Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxScore(nums1 []int, nums2 []int, K int) int64
```

> **💡 Hint:** Sort both arrays. The optimal strategy pairs the largest

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3836: Maximum Score Using Exactly K Pairs
// https://leetcode.com/problems/maximum-score-using-exactly-k-pairs/
// Difficulty: Hard
//
// Pick exactly K pairs (i, j) from nums1 and nums2. Score of each
// pair = nums1[i] + nums2[j]. Maximize total score.
//
// Approach: Sort both arrays. The optimal strategy pairs the largest
// K elements from each array.

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(maxScore([]int{1, 2, 3}, []int{4, 5, 6}, 2))
	// Example 2
	fmt.Println(maxScore([]int{10, 20, 30}, []int{1, 2, 3}, 3))
	// Edge: K=1
	fmt.Println(maxScore([]int{5, 1}, []int{3, 9}, 1))
	// Edge: empty
	fmt.Println(maxScore([]int{}, []int{}, 0))
}

func maxScore(nums1 []int, nums2 []int, K int) int64 {
	if K <= 0 || K > len(nums1) || K > len(nums2) {
		return 0
	}

  // Custom sort dengan comparator
	sort.Slice(nums1, func(i, j int) bool { return nums1[i] > nums1[j] })
  // Custom sort dengan comparator
	sort.Slice(nums2, func(i, j int) bool { return nums2[i] > nums2[j] })

	var total int64
	for i := 0; i < K; i++ {
		total += int64(nums1[i] + nums2[i])
	}

	return total
}
```
