# 0532 — K Diff Pairs In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindPairs(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1) (ignoring sorting overhead)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #532: K-diff Pairs in an Array
// https://leetcode.com/problems/k-diff-pairs-in-an-array/
// Difficulty: Medium
// Time: O(n log n)
// Space: O(1) (ignoring sorting overhead)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(FindPairs([]int{3, 1, 4, 1, 5}, 2))
	fmt.Println(FindPairs([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(FindPairs([]int{1, 3, 1, 5, 4}, 0))
}

func FindPairs(nums []int, k int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	n := len(nums)
	count := 0
	i, j := 0, 1

	for i < n && j < n {
		if i == j || nums[j]-nums[i] < k {
			j++
		} else if nums[j]-nums[i] > k {
			i++
		} else {
			count++
			i++
			j++
			// Skip duplicates
			for j < n && nums[j] == nums[j-1] {
				j++
			}
		}
	}

	return count
}
```
