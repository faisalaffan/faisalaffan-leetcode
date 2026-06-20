# 2779 — Maximum Beauty Of An Array After Applying Operation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumBeautyOfAnArrayAfterApplyingOperation(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2779: Maximum Beauty of an Array After Applying Operation
// https://leetcode.com/problems/maximum-beauty-of-an-array-after-applying-operation/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func MaximumBeautyOfAnArrayAfterApplyingOperation(nums []int, k int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	left := 0
	best := 0
	for right := 0; right < len(nums); right++ {
		for nums[right]-nums[left] > 2*k {
			left++
		}
		if right-left+1 > best {
			best = right - left + 1
		}
	}
	return best
}

func main() {
	fmt.Println(MaximumBeautyOfAnArrayAfterApplyingOperation([]int{4, 6, 1, 2}, 2))
	fmt.Println(MaximumBeautyOfAnArrayAfterApplyingOperation([]int{1, 1, 1, 1}, 10))
}
```
